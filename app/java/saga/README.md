# Ejemplo de uso de microservicios con Temporal y aplicación de patrón SAGA
El ejemplo trata de una aplicación para la gestión de reservaciones de hotel, vuelo, y vehículo para los viajeros. 

La implementación la realizaremos con Java, Spring Framework y Temporal. Si no sabes que es Temporal puedes revisar el siguiente link: [https://sistecma.github.io/2021/02/04/aplicaciones-invencibles-con-temporal.html](https://sistecma.github.io/2021/02/04/aplicaciones-invencibles-con-temporal.html)

Asumiremos 4 microservicios: 

* Hotel: Microservicio encargado de implementar la lógica de negocio referente al agendamiento de hotel (desde el punto de vista de Temporal lo representaremos con actividades)

* Vuelo: Microservicio encargado de implementar la lógica de negocio referente al agendamiento de vuelos (desde el punto de vista de Temporal lo representaremos con actividades)

* Vehículo: Microservicio encargado de implementar la lógica de negocio referente al agendamiento de transporte terrestre desde el aeropuerto al hotel (desde el punto de vista de Temporal lo representaremos con actividades)

* Viaje: Microservicio orquestador (desde el punto de vista de Temporal lo representaremos con un workflow)

También asumiremos un starter simple que lo que hace es invocar al workflow.

# Ejecución del ejemplo
Primero asegúrate de tener temporal listo y corriendo. Puedes descargarte la versión en docker. [Sigue los pasos aqui](https://github.com/temporalio/temporal)

A continuación levanta cada microservicio en una ventana/terminal nueva por cada uno. Al invocar el starter obtendrás el resultado (éxito o falla).
Asumimos te encuentras en el directorio raiz llamado saga (cada vez que abras una nueva terminal para levantar cada microservicio y estas usando maven)

### Levantar microservicio hotel

mvn exec:java -Dexec.mainClass="com.sistecma.temporalio.saga.hotel.HotelWorker"

### Levantar microservicio vuelo

mvn exec:java -Dexec.mainClass="com.sistecma.temporalio.saga.vuelo.VueloWorker"

### Levantar microservicio vehículo

mvn exec:java -Dexec.mainClass="com.sistecma.temporalio.saga.vehiculo.VehiculoWorker"

### Levantar microservicio orquestador viaje

mvn exec:java -Dexec.mainClass="com.sistecma.temporalio.saga.viaje.ViajeWorker"

### Ejecución de starter

mvn exec:java -Dexec.mainClass="com.sistecma.temporalio.saga.starter.SagaStarter"

Para ver el reverso de transacciones (y el funcionamiento de SAGA) puedes por ejemplo dejar subidos todos los microservicios a excepción de el de vuelo. Al ejecutar el starter notarás que después del tiempo definido como timeout en el código se aplicara reverso.

### Nota
El presente ejemplo fue tomado del repositorio original de Temporal y adaptado para ser usado con Spring Framework (bajo nuestro criterio), modificado y comentado al idioma Espanol para mejor entendimiento. El trabajo original esta en este [repositorio](https://github.com/temporalio/samples-java/tree/master/src/main/java/io/temporal/samples/bookingsaga).

El objetivo de este ejemplo es meramente educativo. 

Por ahora el ejemplo no tiene unit test. En próximos ejemplos abordaremos todos los aspectos de unit test con temporal.


















