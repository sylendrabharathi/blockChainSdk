package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chainHero/heroes-service/blockchain"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Channel parameters
		ChannelID:     "chainhero",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/chainHero/heroes-service/fixtures/artifacts/chainhero.channel.tx",

		// Chaincode parameters
		ChainCodeID:     "heroes-service",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/chainHero/heroes-service/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "Org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}

	log.Println("Initialize")
	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
	}

	log.Println("Instantiate")
	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
	}


	user := map[string]interface{}{"id": 1, "name": "Bharathi", "Gender": "Male"}
	// Invoke the chaincode
	txId, err := fSetup.InvokeHello(user)
	if err != nil {
		fmt.Printf("Unable to invoke hello on the chaincode: %v\n", err)
	} else {
		fmt.Printf("Successfully invoke hello, transaction ID: %s\n", txId)
	}

	log.Println("\n Query")
	// Query the chaincode
	response, err := fSetup.QueryHello()
	if err != nil {
		fmt.Printf("Unable to query hello on the chaincode: %v\n", err)
	} else {
		fmt.Printf("Response from the query hello: %s\n", response)
	}

	user["name"] = "Sylendra Bharathi"
	txId, err = fSetup.InvokeHello(user)
	if err != nil {
		fmt.Printf("Unable to invoke hello on the chaincode: %v\n", err)
	} else {
		fmt.Printf("Successfully invoke hello, transaction ID: %s\n", txId)
	}

	user["name"] = "Sylendra Bharathi C"
	txId, err = fSetup.InvokeHello(user)
	if err != nil {
		fmt.Printf("Unable to invoke hello on the chaincode: %v\n", err)
	} else {
		fmt.Printf("Successfully invoke hello, transaction ID: %s\n", txId)
	}

	log.Println("\n After update user ==========================================")
	// Query again the chaincode
	response, err = fSetup.QueryHello()
	if err != nil {
		fmt.Printf("Unable to query hello on the chaincode: %v\n", err)
	} else {
		fmt.Printf("Response from the query hello: %s\n", response)
	}

	log.Println("\n QueryLedgerTransactions Before ===================================")
	// QueryLedgerTransactions again the chaincode
	response, err = fSetup.QueryLedgerTransactions()
	if err != nil {
		fmt.Printf("Unable to QueryLedgerTransactions on the chaincode: %v\n", err)
	} else {
		fmt.Printf("Response from the query QueryLedgerTransactions: %s\n", response)
	}
}
