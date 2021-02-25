package com.sistecma.temporalio.saga.vuelo;

import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.support.AbstractApplicationContext;

import io.temporal.worker.WorkerFactory;

public class VueloWorker {
	

	public static void main(String[] args) {
		@SuppressWarnings("resource")
		AbstractApplicationContext contexto = new AnnotationConfigApplicationContext(VueloConfig.class);	    
		
		WorkerFactory fabrica= (WorkerFactory) contexto.getBean("factoryVuelo");
		fabrica.start();
	
	}
}
