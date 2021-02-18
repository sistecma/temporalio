package com.sistecma.temporalio.saga.hotel;

import io.temporal.activity.ActivityInterface;

@ActivityInterface
public interface HotelI {
	
	// Reserva el hotel
	public String reservar(String nombre);
	
	// Cancela la reservación del hotel
	public String cancelar(String id, String nombre);
	
}
