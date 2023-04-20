package stellar

import (
	"fmt"
	"log"

	"github.com/MicaTechnology/escrow_api/utils/rest_errors"
	"github.com/stellar/go/protocols/horizon"
	"github.com/stellar/go/txnbuild"
)

func buildTransaction(sourceAccount horizon.Account, operations []txnbuild.Operation) (*txnbuild.Transaction, *rest_errors.RestErr) {
	preconditions := txnbuild.Preconditions{
		TimeBounds: txnbuild.NewTimeout(300),
	}

	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &sourceAccount,
			IncrementSequenceNum: true,
			Operations:           operations,
			BaseFee:              txnbuild.MinBaseFee,
			Preconditions:        preconditions,
		},
	)
	if err != nil {
		return nil, restError(err)
	}

	return tx, nil
}

func submitTransaction(tx *txnbuild.Transaction) *rest_errors.RestErr {
	// Get the base 64 encoded transaction envelope
	txe, err := tx.Base64()
	if err != nil {
		log.Fatalln(err)
	}

	// Send the transaction to the network
	resp, err := getClient().SubmitTransactionXDR(txe)
	if err != nil {
		return restError(err)
	}
	fmt.Printf("Transaction successfull! Hash: %s \n", resp.Hash)
	fmt.Printf("Transaction xdr: %s \n", resp.ResultXdr)

	return nil
}
