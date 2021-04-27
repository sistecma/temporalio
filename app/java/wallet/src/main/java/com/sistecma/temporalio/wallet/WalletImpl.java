package com.sistecma.temporalio.wallet;

import io.temporal.workflow.Workflow;

/*
 * In this example we will simulate the concept of a wallet over which we can push or pull "values"
 * Push makes a credit to the current balance inside the wallet
 * Pull makes a debt to the current balance inside the wallet
 */

public class WalletImpl implements WalletI{

	private boolean exit= false;
    private float balance= 0.0f;
    
	@Override
	public void createWallet(float initialBalance) {
		if(initialBalance <= 0) {
			return;
		}
		balance= initialBalance;
		while(true) {
			Workflow.await(() -> exit); // keep the workflow running/awaiting until exit=true
			if(exit){
				return;			
			}
		}
	}

	@Override
	public float queryBalance() {
		return balance;
	}

	@Override
	public void push(float value) {
        balance= balance + value;

	}

	@Override
	public void pull(float value) {
        float temp= balance;
        temp= temp - value;
        if(temp<0) {
        	return;	
        }else {
        	balance= temp;
        }		
	}

	@Override
	public void exit() {
       exit= true;		
	}


}
