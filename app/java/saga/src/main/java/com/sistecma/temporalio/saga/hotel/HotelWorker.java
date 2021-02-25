package com.sistecma.temporalio.saga.hotel;

import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.support.AbstractApplicationContext;

import io.temporal.worker.WorkerFactory;

// Worker actuando como microservicio para la lógica de negocio de hotel
public class HotelWorker {
	

	public static void main(String[] args) {
		@SuppressWarnings("resource")
		AbstractApplicationContext contexto = new AnnotationConfigApplicationContext(HotelConfig.class);	    
		
		WorkerFactory fabrica= (WorkerFactory) contexto.getBean("fabricaHotel");
		fabrica.start(); // arrancamos el servicio de hotel
	
	}
}
