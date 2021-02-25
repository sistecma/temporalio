package com.sistecma.temporalio.saga.comun;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import io.temporal.client.WorkflowClient;
import io.temporal.serviceclient.WorkflowServiceStubs;

/*
 * Configuración común
 * 
 * Simplemente encapsula el bean que obtiene el Workflow Client, que es usado por todas las otras configuraciones
 */

@Configuration
public class Config {

	// bean cliente de workflow
	@Bean
	public WorkflowClient workflowClient() {
		// nos permite interactuar con el servicio de temporal
	    WorkflowServiceStubs service = WorkflowServiceStubs.newInstance();

	    // el cliente puede ser usado para iniciar o senializar workflows
	    return  WorkflowClient.newInstance(service);

	}
}
