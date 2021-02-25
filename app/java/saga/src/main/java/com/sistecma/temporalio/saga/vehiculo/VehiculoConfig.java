package com.sistecma.temporalio.saga.vehiculo;

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
 * Configuración de microservicio vehiculo
 */
@Configuration
@PropertySource("classpath:temporalio.properties") // accede a los properties
@Import({ Config.class }) // importa la configuración Config.class del paquete común
class VehiculoConfig {
	
	// bean worker fábrica para el microservicio vehiculo
	@Bean
	public WorkerFactory factoryVehiculo(@Value("${temporalio.taskqueue.vehiculo}")// trae el valor de temporalio.properties
	                             String colaTareas, // identificador de la cola del microservicio vehiculo
			                     WorkflowClient cliente) { // trae como parámetro el bean WorkflowClient de Config.class

		// el worker fábrica puede ser usado para crear workers
	    WorkerFactory factory = WorkerFactory.newInstance(cliente);

	    // este es el worker que escucha la cola y hostea la lógica de negocio de vehiculo
	    Worker worker = factory.newWorker(colaTareas);
	    worker.registerActivitiesImplementations(new VehiculoImpl()); // registra la implementación de actividades al worker
	    	    
	    return factory;
	}

	

}
