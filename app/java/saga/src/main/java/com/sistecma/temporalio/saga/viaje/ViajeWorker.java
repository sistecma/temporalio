package com.sistecma.temporalio.saga.viaje;

import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.support.AbstractApplicationContext;

import io.temporal.worker.WorkerFactory;

public class ViajeWorker {
	

	public static void main(String[] args) {
		@SuppressWarnings("resource")
		AbstractApplicationContext contexto = new AnnotationConfigApplicationContext(ViajeConfig.class);	    
		
		WorkerFactory fabrica= (WorkerFactory) contexto.getBean("factoryViaje");
		fabrica.start();
	
	}
}
