### Timer example

This example considers many useful aspects of Temporal platform. 

1- The usage of timers feature.

2- How to query running workflows in Temporal.

3- How to update state in running workflows by using signals feature.

The workflow works as follow:

* It has a timer that allows to specify a timeout for the workflow.

* The timeout is only resetted when a signal is sent to the workflow.

* If the timeout is reached the workflow is cancelled and counter is zeroed

* The workflow allows to query how many times a signal was sent within the timeout interval

