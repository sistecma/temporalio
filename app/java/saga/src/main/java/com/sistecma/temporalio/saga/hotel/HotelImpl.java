package com.sistecma.temporalio.saga.hotel;

import java.util.UUID;

public class HotelImpl implements HotelI {

	@Override
	public String reservar(String nombre) {
		System.out.println("reservar hotel con nombre: '" + nombre + "'");
	    return UUID.randomUUID().toString();
	}

	@Override
	public String cancelar(String id, String nombre) {
		System.out.println("cancelar reservación de hotel con nombre: '" + nombre + "'" + " id: " + id);
		return UUID.randomUUID().toString();
	}

}
