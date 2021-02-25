package com.sistecma.temporalio.saga.vuelo;

import java.util.UUID;

class VueloImpl implements VueloI {

	@Override
	public String reservar(String nombre) {
		System.out.println("reservar vuelo con nombre: '" + nombre + "'");
	    return UUID.randomUUID().toString();
	}

	@Override
	public String cancelar(String id, String nombre) {
		System.out.println("cancelar reservación de vuelo con nombre: '" + nombre + "'" + " id: " + id);
		return UUID.randomUUID().toString();
	}

}
