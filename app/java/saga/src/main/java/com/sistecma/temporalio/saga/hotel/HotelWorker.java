package com.sistecma.temporalio.saga.hotel;

import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.support.AbstractApplicationContext;

import io.temporal.worker.WorkerFactory;

public class HotelWorker {
	

	public static void main(String[] args) {
		@SuppressWarnings("resource")
		AbstractApplicationContext contexto = new AnnotationConfigApplicationContext(HotelConfig.class);	    
		
		WorkerFactory factory= (WorkerFactory) contexto.getBean("factory");
		factory.start();
	}
}
