package com.sistecma.temporalio.saga.vehiculo;

import java.util.UUID;

class VehiculoImpl implements VehiculoI {

	@Override
	public String reservar(String nombre) {
		System.out.println("reservar vehiculo con nombre: '" + nombre + "'");
	    return UUID.randomUUID().toString();
	}

	@Override
	public String cancelar(String id, String nombre) {
		System.out.println("cancelar reservación de vehiculo con nombre: '" + nombre + "'" + " id: " + id);
		return UUID.randomUUID().toString();
	}

}
