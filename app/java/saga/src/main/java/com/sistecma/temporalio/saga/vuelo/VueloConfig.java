package com.sistecma.temporalio.saga.vuelo;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Import;
import org.springframework.context.annotation.PropertySource;

import com.sistecma.temporalio.saga.comun.Config;

import io.temporal.client.WorkflowClient;
import io.temporal.worker.Worker;
import io.temporal.worker.WorkerFactory;

/*
 * Configuración de microservicio vuelo
 */
@Configuration
@PropertySource("classpath:temporalio.properties") // accede a los properties
@Import({ Config.class }) // importa la configuración Config.class del paquete común
class VueloConfig {
	@Bean
	public WorkerFactory factoryVuelo(@Value("${temporalio.taskqueue.vuelo}") // trae el valor de temporalio.properties
	                             String colaTareas, // identificador de la cola del microservicio vuelo
			                     WorkflowClient cliente) { // trae como parámetro el bean WorkflowClient de Config.class

		// el worker fábrica puede ser usado para crear workers
	    WorkerFactory fabrica = WorkerFactory.newInstance(cliente);

	    // este es el worker que escucha la cola y hostea la lógica de negocio de vehiculo
	    Worker worker = fabrica.newWorker(colaTareas);
	    worker.registerActivitiesImplementations(new VueloImpl()); // registra la implementación de actividades al worker
	    	    
	    return fabrica;
	}

	

}
