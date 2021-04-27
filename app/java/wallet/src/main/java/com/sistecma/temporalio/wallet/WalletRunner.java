package com.sistecma.temporalio.wallet;

import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import io.temporal.serviceclient.WorkflowServiceStubs;
import io.temporal.worker.Worker;
import io.temporal.worker.WorkerFactory;

public class WalletRunner {

	public static void main(String[] args) {
		// Define the workflow service.
	    WorkflowServiceStubs service = WorkflowServiceStubs.newInstance();

	    /*
	     * Define the workflow client. It is a Temporal service client used to start, signal, and query
	     * workflows
	     */
	    WorkflowClient client = WorkflowClient.newInstance(service);

	    /*
	     * Define the workflow factory. It is used to create workflow workers for a specific task queue.
	     */
	    WorkerFactory factory = WorkerFactory.newInstance(client);

	    /*
	     * Define the workflow worker. Workflow workers listen to a defined task queue and process
	     * workflows and activities.
	     */
	    Worker worker = factory.newWorker(WalletI.TASK_QUEUE);

	    /*
	     * Register the workflow implementation with the worker.
	     * Workflow implementations must be known to the worker at runtime in
	     * order to dispatch workflow tasks.
	     */
	    worker.registerWorkflowImplementationTypes(WalletImpl.class);

	    /*
	     * Start all the workers registered for a specific task queue.
	     * The started workers then start polling for workflows and activities.
	     */
	    factory.start();

	    // Create the workflow options
	    WorkflowOptions workflowOptions =
	        WorkflowOptions.newBuilder().setTaskQueue(WalletI.TASK_QUEUE).setWorkflowId(WalletI.WORKFLOW_ID).build();

	    // Create the workflow client stub. It is used to start the workflow execution.
	    WalletI workflow = client.newWorkflowStub(WalletI.class, workflowOptions);

	    // Start workflow asynchronously and call its createWallet workflow method
	    WorkflowClient.start(workflow::createWallet,10.0f);
	    
	    float temp= workflow.queryBalance();
	    System.out.println("current balance: " + String.valueOf(temp));
	    
	    workflow.push(20.0f);
	    
	    temp= workflow.queryBalance();
	    System.out.println("current balance: " + String.valueOf(temp));
	    
	    workflow.pull(5.1f);
	    
	    temp= workflow.queryBalance();
	    System.out.println("current balance: " + String.valueOf(temp));	 
	    
	    workflow.exit();
	    System.exit(0);
	}

}
