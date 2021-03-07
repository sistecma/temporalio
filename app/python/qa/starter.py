import asyncio
import logging

from temporal.workflow import workflow_method, WorkflowClient

logging.basicConfig(level=logging.INFO)

# ENG
# Task queue name and namespace
# SPA
# El nombre del namespace y el taskqueue
TASK_QUEUE = "qa-tq"
NAMESPACE = "default"

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

async def client_main():
    # ENG
    # 1- We create the new client
    # 2- We create el workflow stub
    # 3- We invoke the workflow method and wait for the result.
    # SPA
    # 1- Creamos el nuevo cliente
    # 2- Creamos el workflow subordinado/avatar(objeto remoto que tiene la interface del worklow)
    # 3- Invocamos el método workflow y esperamos por el resultado.
    client = WorkflowClient.new_client(namespace=NAMESPACE)
    qa_workflow: QAWorkflow = client.new_workflow_stub(QAWorkflow)
    result = await qa_workflow.apply_qa("What's your name? / Cual es tu nombre?", "Hernan Moreno")
    print(result)

if __name__ == '__main__':
    asyncio.run(client_main())
