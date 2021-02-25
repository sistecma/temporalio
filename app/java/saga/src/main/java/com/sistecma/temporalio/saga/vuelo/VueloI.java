package com.sistecma.temporalio.saga.vuelo;

import io.temporal.activity.ActivityInterface;

@ActivityInterface
public interface VueloI {
	
	// Reserva el vehiculo
	public String reservar(String nombre);
	
	// Cancela la reservación del vehiculo
	public String cancelar(String id, String nombre);
	
}
