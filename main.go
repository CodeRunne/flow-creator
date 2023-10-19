package main

import (
	"context"
	"fmt"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/templates"
	"github.com/onflow/flow-go-sdk/access/http"
	"github.com/coderunne/creator/pkg/utility"
)

// NOTE: All new account must be created by an existing account on any of the networks.
// Testnet account must be funded

func main() {
	// initialize context
	ctx := context.Background()

	// Connect To Flow Host (EmulatorHost|TestnetHost|MainnetHost)
	client, err := http.NewClient(utility.Host)
	utility.Handle(err)

	// Get Latest Block
	block, err := client.GetLatestBlock(ctx, true)
	utility.Handle(err)

	// Generate Mnemonic seed
	mnemonic := utility.GenerateMnemonicKey()

	// Generate Private Key
	privatekey, err := crypto.GeneratePrivateKey(crypto.ECDSA_P256, []byte(mnemonic))
	utility.Handle(err)

	// Generate New Flow Account Key
	accKey := flow.NewAccountKey().
		FromPrivateKey(privatekey).
		SetHashAlgo(crypto.SHA3_256).
		SetWeight(flow.AccountKeyWeightThreshold)

	// Get Account
	account, err := client.GetAccount(ctx, flow.HexToAddress(utility.NetworkKey))
	utility.Handle(err)

	// Retrive address and keys from account
	address := account.Address
	sacckey := account.Keys[0]

	// Decode the Private key
	privateKey, err := crypto.DecodePrivateKeyHex(crypto.ECDSA_P256, utility.NetworkPrivateKey)
	utility.Handle(err)

	// Get the signer from the private key
	signer, err := crypto.NewInMemorySigner(privateKey, crypto.SHA3_256)
	utility.Handle(err)

	// Send the account creation transaction
	accountTx, err := templates.CreateAccount([]*flow.AccountKey{accKey}, nil, address)
	utility.Handle(err)

	accountTx.SetProposalKey(address, sacckey.Index, sacckey.SequenceNumber).
		SetReferenceBlockID(block.ID).
		SetPayer(address)

	// Sign Transaction Envelope
	err = accountTx.SignEnvelope(address, sacckey.Index, signer)
	utility.Handle(err)

	// Send transaction
	err = client.SendTransaction(ctx, *accountTx)
	utility.Handle(err)

	// Get transaction result
	tResult, err := client.GetTransactionResult(ctx, accountTx.ID())
	utility.Handle(err)

	// Retrieve/Get address from emitted events
	var myAddress flow.Address
	for _, event := range tResult.Events {
		if event.Type == flow.EventAccountCreated {
			accountCreatedEvent := flow.AccountCreatedEvent(event)
			myAddress = accountCreatedEvent.Address()
		}
	}

	// Save to file
	utility.SaveFile(myAddress.String(), mnemonic)
	
	// Print results to the console
	fmt.Printf("===== Account Created ===== \n Mnemonic Seed: %v \n Address: %v \n", mnemonic, myAddress)

	// Print transaction id
	fmt.Printf("===== Transaction ID ===== \n %v \n", accountTx.ID())
}