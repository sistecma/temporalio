### Parent - Child workflow example

This example considers Child workflows usage in Temporal platform. 

Child workflow enables the scheduling of other Workflows from within a Workflow's implementation. The parent Workflow has the ability to monitor and impact the lifecycle of the child Workflow, similar to the way it does for an Activity that it invoked.

To execute this example:

1- You need a Temporal service running.
2- Run the following command to start the worker
```
go run child/worker/main.go
```
3- Run the following command to start the example
```
go run child/starter/main.go
```
