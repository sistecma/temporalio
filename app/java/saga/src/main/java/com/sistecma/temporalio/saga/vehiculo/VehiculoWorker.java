package com.sistecma.temporalio.saga.vehiculo;

import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.support.AbstractApplicationContext;

import io.temporal.worker.WorkerFactory;

public class VehiculoWorker {
	

	public static void main(String[] args) {
		@SuppressWarnings("resource")
		AbstractApplicationContext contexto = new AnnotationConfigApplicationContext(VehiculoConfig.class);	    
		
		WorkerFactory fabrica= (WorkerFactory) contexto.getBean("factoryVehiculo");
		fabrica.start();
	
	}
}
