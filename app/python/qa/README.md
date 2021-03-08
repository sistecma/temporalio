## README IN ENGLISH AND SPANISH / README EN INGLES Y ESPANIOL

### README IN ENGLISH
#### Example of microservice usage with Temporal and Python
In this example we will illustrate how to use Temporal with python3 in a microservice architecture. If you are not familiar with Temporal, please follow this [link](https://temporal.io/)

We define simple questions and answers as an application. We assume 3 microservices:

Questions microservice- This is the microservice dealing with all the questions. The questions are just simple actitity methods and the worker registers all those activities

Answers microservice- This is the microservice dealing with all the answers. The answers are just simple activity methods and the worker registers all those activities

QA microservice- This is the microservice dealing with the workflow.

Starter- This is the python script to just invoke the example.

Execute the example:

1- Open a terminal and Go to the project root

2- Execute: python3 qa.py

3- Open other terminal with the same path

4- Execute: python3 questions.py

5- Open other terminal with the same path

6- Execute: python3 answers.py

7- Open other terminal with the same path

8- Execute: python3 starter.py

### README EN ESPANOL
#### Ejemplo de microservicios con Temporal y Python
En este ejemplo ilustramos como usar Temporal con python3 en una arquitectura de microservicios.Si no estas familiarizado con Temporal puedes seguir este enlace [link](https://temporal.io/) o visitar mi [blog](https://sistecma.github.io/)

Definimos una simple aplicación llamada preguntas y respuestas (q&a). Para esto asumimos los siguientes microservicios:

Microservicio de Preguntas/Questions- Este el microservicio que se encarga de manejar de las preguntas. Las preguntas son simples métodos que son registrados en el worker dentro del archivo python.

Microservicio de Respuestas/Answers-  Este el microservicio que se encarga de manejar de las respuestas. Las respuestas son simples métodos que son registrados en el worker dentro del archivo python.

Microservicio de QA- Este es el microservicio que lidia con el workflow.

Starter- Este es un script de python que invoca el ejemplo.

Pasos para ejecutar el ejemplo:

1- Abrir la terminal y ir al directorio raiz del proyecto. 

2- Ejecutar: python3 qa.py

3- Abrir la terminal y ir al directorio raiz del proyecto. 

4- Ejecutar: python3 questions.py

5- Abrir la terminal y ir al directorio raiz del proyecto. 

6- Ejecutar: python3 answers.py

7- Abrir la terminal y ir al directorio raiz del proyecto. 

8- Ejecutar: python3 starter.py



