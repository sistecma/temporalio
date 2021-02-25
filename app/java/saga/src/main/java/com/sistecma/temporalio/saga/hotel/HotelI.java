package com.sistecma.temporalio.saga.hotel;

import io.temporal.activity.ActivityInterface;

// interface para la actividad que sostiene la lógica de negocio de hotel
@ActivityInterface
public interface HotelI {
	
	// Reserva el hotel
	public String reservar(String nombre);
	
	// Cancela la reservación del hotel
	public String cancelar(String id, String nombre);
	
}
