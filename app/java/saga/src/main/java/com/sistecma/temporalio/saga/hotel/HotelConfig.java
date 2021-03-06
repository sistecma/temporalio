package com.sistecma.temporalio.saga.hotel;

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
 * Configuración de microservicio hotel
 */
@Configuration
@PropertySource("classpath:temporalio.properties") // accede a los properties
@Import({ Config.class }) // importa la configuración Config.class del paquete común
class HotelConfig {
	
	// bean worker fábrica para el microservicio hotel
	@Bean
	public WorkerFactory fabricaHotel(@Value("${temporalio.taskqueue.hotel}") // trae el valor de temporalio.properties
	                             String colaTareas, // identificador de la cola del microservicio hotel
			                     WorkflowClient cliente) { // trae como parámetro el bean WorkflowClient de Config.class

		// el worker fábrica puede ser usado para crear workers
	    WorkerFactory fabrica = WorkerFactory.newInstance(cliente);

	    // este es el worker que escucha la cola y hostea la lógica de negocio de hotel
	    Worker worker = fabrica.newWorker(colaTareas);
	    worker.registerActivitiesImplementations(new HotelImpl()); // registra la implementación de actividades al worker
	    	    
	    return fabrica;
	}

	

}
