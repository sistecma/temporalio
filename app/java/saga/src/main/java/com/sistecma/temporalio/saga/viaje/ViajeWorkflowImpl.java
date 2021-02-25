/*
 *  Copyright (c) 2020 Temporal Technologies, Inc. All Rights Reserved
 *
 *  Copyright 2012-2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 *  Modifications copyright (C) 2017 Uber Technologies, Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"). You may not
 *  use this file except in compliance with the License. A copy of the License is
 *  located at
 *
 *  http://aws.amazon.com/apache2.0
 *
 *  or in the "license" file accompanying this file. This file is distributed on
 *  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 *  express or implied. See the License for the specific language governing
 *  permissions and limitations under the License.
 */

package com.sistecma.temporalio.saga.viaje;

import java.time.Duration;

import com.sistecma.temporalio.saga.hotel.HotelI;
import com.sistecma.temporalio.saga.vehiculo.VehiculoI;
import com.sistecma.temporalio.saga.vuelo.VueloI;

import io.temporal.activity.ActivityOptions;
import io.temporal.common.RetryOptions;
import io.temporal.failure.ActivityFailure;
import io.temporal.workflow.Saga;
import io.temporal.workflow.Workflow;

public class ViajeWorkflowImpl implements ViajeWorkflowI {

 // Uso clases estaticas para configuración extener a fin de no romper la condición deterministica del workflow
 public static class Hotel{
	 public static String TASKQUEUE;
	 public static Duration TIMEOUT;
	 public static int ATTEMPTS;
 }

 public static class Vehiculo{
	 public static String TASKQUEUE;
	 public static Duration TIMEOUT;
	 public static int ATTEMPTS;
 }

 public static class Vuelo{
	 public static String TASKQUEUE;
	 public static Duration TIMEOUT;
	 public static int ATTEMPTS;
 }

  private final ActivityOptions optionsHotel =
      ActivityOptions.newBuilder()
          .setTaskQueue(Hotel.TASKQUEUE)
          .setScheduleToCloseTimeout(Hotel.TIMEOUT)
          .setRetryOptions(RetryOptions.newBuilder().setMaximumAttempts(Hotel.ATTEMPTS).build())
          .build();

  private final ActivityOptions optionsVehiculo =
	      ActivityOptions.newBuilder()
	          .setTaskQueue(Vehiculo.TASKQUEUE)
	          .setScheduleToCloseTimeout(Vehiculo.TIMEOUT)
	          .setRetryOptions(RetryOptions.newBuilder().setMaximumAttempts(Vehiculo.ATTEMPTS).build())
	          .build();

  private final ActivityOptions optionsVuelo =
	      ActivityOptions.newBuilder()
	          .setTaskQueue(Vuelo.TASKQUEUE)
	          .setScheduleToCloseTimeout(Vuelo.TIMEOUT)
	          .setRetryOptions(RetryOptions.newBuilder().setMaximumAttempts(Vuelo.ATTEMPTS).build())
	          .build();  
  
  private final HotelI hotel =
      Workflow.newActivityStub(HotelI.class, optionsHotel);
  
  private final VehiculoI vehiculo =
	      Workflow.newActivityStub(VehiculoI.class, optionsVehiculo); 
  
  private final VueloI vuelo =
	      Workflow.newActivityStub(VueloI.class, optionsVuelo); 
  
  @Override
  public boolean agendar(String name) {
    // Configuración de SAGA
    Saga.Options sagaOptions = new Saga.Options.Builder().setParallelCompensation(true).build();
    Saga saga = new Saga(sagaOptions);
    try {
        String hotelReservationID = hotel.reservar(name);
        saga.addCompensation(hotel::cancelar, hotelReservationID, name);
    	
        String carReservationID = vehiculo.reservar(name);
        saga.addCompensation(vehiculo::cancelar, carReservationID, name);

        String flightReservationID = vuelo.reservar(name);
        saga.addCompensation(vuelo::cancelar, flightReservationID, name);
    } catch (ActivityFailure e) {
      saga.compensate();
      return false;
    }
    return true;
  }
}
