import asyncio
import logging
from datetime import timedelta

from temporal.activity_method import activity_method
from temporal.workerfactory import WorkerFactory
from temporal.workflow import workflow_method, Workflow, WorkflowClient

logging.basicConfig(level=logging.INFO)

# ENG
# Task queue name and namespace
# SPA
# El nombre del namespace y el taskqueue 
TASK_QUEUE = "qa-tq"
NAMESPACE = "default"

# ENG
# This python script implements the microservice qa which contains the workflow and the worker registering the workflow.
# SPA
# Este script en python implementa el microservicio qa que contiene el workflow y el worker que registra el workflow.


# ENG- Activities Interface which defines the activities o business logic. The implementation of the activities can be found within the microservice called question 
# SPA- Interface que define las actividades o lógica de negocio. La implementación de las actividades se encuentran en el microservicio question
class QuestionActivities:
    # ENG
    # We instruct the activity method to use the task queue qa-tq and the schedule to close timeout with 10 seconds.
    # ask_questions is the activity method
    # SPA
    # Instruimos el método de actividad para usar el task queue qa-tq y el schedule to close timeout en 10 segundos.
    # ask_questions es el método actividad
    @activity_method(task_queue=TASK_QUEUE, schedule_to_close_timeout=timedelta(seconds=10))
    async def ask_questions(self, question: str, name: str) -> str:
        raise NotImplementedError


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

    
# ENG- Workflow Interface
# SPA- Interface del workflow
class QAWorkflow:
    # ENG
    # We instruct the workflow method to use the task queue qa-tq
    # apply_qa is the workflow method
    # SPA
    # Instruimos el método de workflow para user el task queue qa-tq
    @workflow_method(task_queue=TASK_QUEUE)
    async def apply_qa(self, question: str, name: str) -> str:
        raise NotImplementedError


# ENG- Workflow Implementation
# SPA- Implementación del workflow
class QAWorkflowImpl(QAWorkflow):

    def __init__(self):
        self.question_activities: QuestionActivities = Workflow.new_activity_stub(QuestionActivities)
        self.answer_activities: AnswerActivities = Workflow.new_activity_stub(AnswerActivities)

    async def apply_qa(self, question, name):
       questionstr = await self.question_activities.ask_questions("Ask/Pregunta: ", question)
       answerstr= await self.answer_activities.reply("\nMy name is / Mi nombre es: ", name)
       return questionstr + " " + answerstr

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
    #4- Registramos la implementación del workflow
    #5- Iniciamos el factory

    client = WorkflowClient.new_client(namespace=NAMESPACE)
    factory = WorkerFactory(client, NAMESPACE)
    worker = factory.new_worker(TASK_QUEUE)
    worker.register_workflow_implementation_type(QAWorkflowImpl)
    factory.start()
    print("Worker started")
    
if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    asyncio.ensure_future(client_main())
    loop.run_forever()
