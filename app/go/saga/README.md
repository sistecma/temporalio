# Ejemplo de uso de microservicios con Temporal y aplicación de patrón SAGA
El ejemplo trata de una aplicación para la gestión de reservaciones de hotel, vuelo, y vehículo para los viajeros. 

La implementación la realizaremos con Golang y Temporal. Si no sabes que es Temporal puedes revisar el siguiente link: [https://sistecma.github.io/2021/02/04/aplicaciones-invencibles-con-temporal.html](https://sistecma.github.io/2021/02/04/aplicaciones-invencibles-con-temporal.html)

Asumiremos 4 microservicios: 

* Hotel: Microservicio encargado de implementar la lógica de negocio referente al agendamiento de hotel (desde el punto de vista de Temporal lo representaremos con actividades)

* Vuelo: Microservicio encargado de implementar la lógica de negocio referente al agendamiento de vuelos (desde el punto de vista de Temporal lo representaremos con actividades)

* Vehículo: Microservicio encargado de implementar la lógica de negocio referente al agendamiento de transporte terrestre desde el aeropuerto al hotel (desde el punto de vista de Temporal lo representaremos con actividades)

* Viaje: Microservicio orquestador (desde el punto de vista de Temporal lo representaremos con un workflow)

También asumiremos un starter simple que lo que hace es invocar al workflow.

# Ejecución del ejemplo
Primero asegúrate de tener temporal listo y corriendo. Puedes descargarte la versión en docker. [Sigue los pasos aqui](https://github.com/temporalio/temporal)

A continuación levanta cada microservicio en una ventana/terminal nueva por cada uno. Al invocar el starter obtendrás el resultado (éxito o falla).
Asumimos te encuentras en el directorio raiz llamado saga (cada vez que abras una nueva terminal para levantar cada microservicio)

Levantar microservicio hotel

cd hotel

go run main.go

Levantar microservicio vuelo

cd vuelo

go run main.go

Levantar microservicio vehículo

cd vehículo

go run main.go

Levantar microservicio orquestador viaje

cd viaje

go run main.go

Ejecución de starter

cd starter

go run main.go

Para ver el reverso de transacciones (y el funcionamiento de SAGA) puedes por ejemplo dejar subidos todos los microservicios a excepción de el de vuelo. Al ejecutar el starter notarás que después del tiempo definido como timeout en el código se aplicara reverso.

