package com.sistecma.temporalio.saga.viaje;

import java.time.Duration;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Import;
import org.springframework.context.annotation.PropertySource;

import com.sistecma.temporalio.saga.comun.Config;

import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import io.temporal.worker.Worker;
import io.temporal.worker.WorkerFactory;

/*
 * Configuración de microservicio viaje que implementa SAGA
 */
@Configuration
@PropertySource("classpath:temporalio.properties") // accede a los properties
@Import({ Config.class }) // importa la configuración Config.class del paquete común
public class ViajeConfig {

    @Value("${temporalio.taskqueue.viaje}") // trae el valor de temporalio.properties
    private String colaTareas; // cola propia del microservicio viaje
    
    @Value("${temporalio.taskqueue.hotel}")
    private String hotelTaskQueue; // cola del microservicio hotel para enrutamiento
    
    @Value("${temporalio.taskqueue.vehiculo}")
    private String vehiculoTaskQueue; // cola del microservicio vehiculo para enrutamiento
    	                             
    @Value("${temporalio.taskqueue.vuelo}")
    private String vueloTaskQueue;  // colar del microservicio vuelo para enrutamiento

    @Value("${temporalio.timeout.hotel}")
    private long hotelTimeout; // timeout considerado para el hotel
    
    @Value("${temporalio.timeout.vehiculo}")
    private long vehiculoTimeout; // timeout considerado para el vehiculo

    @Value("${temporalio.timeout.vuelo}")
    private long vueloTimeout; // timeout considerado para el vuelo

    @Value("${temporalio.attempts.hotel}")
    private int hotelAttempts; // intentos considerado para el hotel
    
    @Value("${temporalio.attempts.vehiculo}")
    private int vehiculoAttempts; // intentos considerado para el vehiculo

    @Value("${temporalio.attempts.vuelo}")
    private int vueloAttempts; // intentos considerado para el vuelo
    
	@Bean
	public WorkerFactory factoryViaje(
	                             WorkflowClient cliente) {// trae como parámetro el bean WorkflowClient de Config.class

		// el worker fábrica puede ser usado para crear workers
	    WorkerFactory fabrica = WorkerFactory.newInstance(cliente);

	    // este es el worker que escucha la cola y hostea el orquestador de los microservicios
	    Worker worker = fabrica.newWorker(colaTareas);
        
	    // enrutamiento de los microservicios
	    ViajeWorkflowImpl.Hotel.TASKQUEUE= hotelTaskQueue;
	    ViajeWorkflowImpl.Vehiculo.TASKQUEUE= vehiculoTaskQueue;
	    ViajeWorkflowImpl.Vuelo.TASKQUEUE= vueloTaskQueue;

	    // timeouts que configuran el StartToClose timeout
	    ViajeWorkflowImpl.Hotel.TIMEOUT= Duration.ofSeconds(hotelTimeout);
	    ViajeWorkflowImpl.Vehiculo.TIMEOUT= Duration.ofSeconds(vehiculoTimeout);
	    ViajeWorkflowImpl.Vuelo.TIMEOUT= Duration.ofSeconds(vueloTimeout);
	    
	    // attempts o intentos considerados por si hay fallas (temporal puede realizar reintentos arbitrariamente si hay fallas)
	    ViajeWorkflowImpl.Hotel.ATTEMPTS= hotelAttempts;
	    ViajeWorkflowImpl.Vehiculo.ATTEMPTS= vehiculoAttempts;
	    ViajeWorkflowImpl.Vuelo.ATTEMPTS= vueloAttempts;
	    
	    worker.registerWorkflowImplementationTypes(ViajeWorkflowImpl.class); // registra el worflow

	    return fabrica;
	}

    // bean que nos da el workflow con SAGA listo para ser usado. En este caso será usado en SagaStarter
	@Bean
	public ViajeWorkflowI workflow(@Value("${temporalio.taskqueue.viaje}") 
                    String colaTareas,
                    WorkflowClient cliente) {
	    WorkflowOptions options = WorkflowOptions.newBuilder().setTaskQueue(colaTareas).build(); // asigna la cola 
	    ViajeWorkflowI viaje = cliente.newWorkflowStub(ViajeWorkflowI.class, options);
	    return viaje;	    
	}

}
