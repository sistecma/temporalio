package com.sistecma.temporalio.wallet;

import io.temporal.workflow.QueryMethod;
import io.temporal.workflow.SignalMethod;
import io.temporal.workflow.WorkflowInterface;
import io.temporal.workflow.WorkflowMethod;

@WorkflowInterface
public interface WalletI {

	// Define the task queue name
	public static final String TASK_QUEUE = "wallet-queue";

	// Define our workflow unique id
	public static final String WORKFLOW_ID = "wallet-workflow";

	@WorkflowMethod
	public void createWallet(float initialBalance);

	// Workflow query method. Used to return our greeting as a query value
	@QueryMethod
	public float queryBalance();

	@SignalMethod
	void push(float value);

	@SignalMethod
	void pull(float value);

    @SignalMethod
    void exit();	
}
