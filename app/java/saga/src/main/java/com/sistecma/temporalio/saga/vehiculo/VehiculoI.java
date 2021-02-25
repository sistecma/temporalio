package com.sistecma.temporalio.saga.vehiculo;

import io.temporal.activity.ActivityInterface;

@ActivityInterface
public interface VehiculoI {
	
	// Reserva el vehiculo
	public String reservar(String nombre);
	
	// Cancela la reservación del vehiculo
	public String cancelar(String id, String nombre);
	
}
