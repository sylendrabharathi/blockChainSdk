package blockchain

import (
	"fmt"
	"log"

	"github.com/hyperledger/fabric-sdk-go/api/apitxn/chclient"
	"encoding/json"
)

// QueryHello query the chaincode to get the state of hello
func (setup *FabricSetup) QueryHello() (interface{}, error) {

	// Prepare arguments
	var args []string
	args = append(args, "invoke")
	args = append(args, "query")
	args = append(args, "hello")

	response, err := setup.client.Query(chclient.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	res := make(map[string]interface{})
	json.Unmarshal(response.Payload, &res)
	log.Println("Response from query payload : ***************** ", res)
	log.Println("Response from query payload : -----------------", string(response.Payload))

	return response.TransactionID.ID, nil
}

func (setup *FabricSetup) QueryLedgerTransactions() (interface{}, error) {

	log.Println("Query Ledger Transactions")
	var args []string
	args = append(args, "invoke")
	args = append(args, "getHistory")
	args = append(args, "hello")

	var res interface{}

	response, err := setup.client.Query(chclient.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	for _, r := range response.Responses {
		var resq interface{}
		json.Unmarshal(r.ProposalResponse.Payload, &resq)
		log.Println("Check ===========")
		log.Println(r.Proposal.TxnID.ID)
	}
	log.Printf("Response from query payload QueryLedgerTransactions : %s", response.Payload)

	err = json.Unmarshal(response.Payload, &res)
	log.Println("Error in unmarshal = " , err.Error())

	return res, nil
}
