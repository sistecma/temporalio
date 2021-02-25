package com.sistecma.temporalio.saga.starter;

import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.support.AbstractApplicationContext;

import com.sistecma.temporalio.saga.viaje.ViajeConfig;
import com.sistecma.temporalio.saga.viaje.ViajeWorkflowI;

import io.temporal.client.WorkflowException;



public class SagaStarter {
	

	public static void main(String[] args) {
		@SuppressWarnings("resource")
		AbstractApplicationContext contexto = new AnnotationConfigApplicationContext(ViajeConfig.class);	    
		
		ViajeWorkflowI viaje= (ViajeWorkflowI) contexto.getBean("workflow");
		boolean status= false;
		try {
	        status= viaje.agendar("mi paseo"); // ejecutar el workflow con SAGA
	        
	    } catch (WorkflowException e) {
	    	throw e;
	    } 
		System.out.println("status= " + status);
		System.exit(0);
	
	}
}
