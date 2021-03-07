import asyncio
import logging
from datetime import timedelta

from temporal.activity_method import activity_method
from temporal.workerfactory import WorkerFactory
from temporal.workflow import WorkflowClient

logging.basicConfig(level=logging.INFO)

# ENG- We define the task queue name and namespace. In this case it's qa-tq because it's for the questions microservice's queue
# SPA_ Definimos el nombre del task queue y namespace. En este caso es qa-tq porque es para el microservicio de questions/preguntas
TASK_QUEUE = "qa-tq"
NAMESPACE = "default"

# ENG- Activities Interface which defines the activities o business logic. The implementation of the activities can be found within the microservice called answers
# SPA- Interface que define las actividades o lógica de negocio. La implementación de las actividades se encuentran en el microservicio answers
class AnswerActivities:
    # ENG
    # We instruct the activity method to use the task queue qa-tq and the schedule to close timeout with 10 seconds.
    # ask_questions is the activity method
    # SPA
    # Instruimos el método de actividad para usar el task queue qa-tq y el schedule to close timeout en 10 segundos.
    # ask_questions es el método actividad
    @activity_method(task_queue=TASK_QUEUE, schedule_to_close_timeout=timedelta(seconds=10))
    async def reply(self, answer: str, name: str) -> str:
        raise NotImplementedError


# ENG- Activities Implementation
# SPA- Implementación de Actividades o lógica de negocio
class AnswerActivitiesImpl:
    async def reply(self, answer: str, name: str):
        return answer + " " + name


async def client_main():
    # ENG
    # 1- We create a new client associated with the namespace
    # 2- We define a workerfactory associated with the namespace
    #3- We create the worker
    #4- We register the workflow implementation
    #5- We start the factory

    # SPA
    # 1- Creamos el nuevo cliente asociado con el namespace
    #2- Definimos el workerfactory asociado con el namespace
    #3- Creamos el worker
    #4- Registramos la implementación de la actividad
    #5- Iniciamos el factory
    client = WorkflowClient.new_client(namespace=NAMESPACE)
    factory = WorkerFactory(client, NAMESPACE)
    worker = factory.new_worker(TASK_QUEUE)
    worker.register_activities_implementation(
        AnswerActivitiesImpl(), "AnswerActivities")
    factory.start()
    print("Worker started")

if __name__ == '__main__':
    loop = asyncio.get_event_loop()  # (1)
    asyncio.ensure_future(client_main())
    loop.run_forever()
