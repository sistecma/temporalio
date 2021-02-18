package com.sistecma.temporalio.saga.hotel;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.PropertySource;

import io.temporal.client.WorkflowClient;
import io.temporal.serviceclient.WorkflowServiceStubs;
import io.temporal.worker.Worker;
import io.temporal.worker.WorkerFactory;

@Configuration
@PropertySource("classpath:temporalio.properties")
public class HotelConfig {

	@Bean
	public HotelI hotel() {
	    // Activities are stateless and thread safe. So a shared instance is used.
	    HotelI hotel = new HotelImpl();
        return hotel; 	
	}
	
	@Bean
	public WorkerFactory factory(@Value("${temporalio.taskqueue}") String colaTareas, HotelI hotel) {
		// gRPC stubs wrapper that talks to the local docker instance of temporal service.
	    WorkflowServiceStubs service = WorkflowServiceStubs.newInstance();
	    
	    // client that can be used to start and signal workflows
	    WorkflowClient client = WorkflowClient.newInstance(service);

	    // worker factory that can be used to create workers for specific task queues
	    WorkerFactory factory = WorkerFactory.newInstance(client);

	    // Worker that listens on a task queue and hosts both workflow and activity implementations.
	    Worker worker = factory.newWorker(colaTareas);

	    worker.registerActivitiesImplementations(hotel);
	    
	    // Start all workers created by this factory.
	    //factory.start();
	    
	    return factory;
	}
	
}
