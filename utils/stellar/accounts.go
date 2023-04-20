package stellar

import (
	"os"

	"github.com/MicaTechnology/escrow_api/utils/logger"
	"github.com/MicaTechnology/escrow_api/utils/rest_errors"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/protocols/horizon"
	"github.com/stellar/go/txnbuild"
)

const MinBalance = "1.0"

func GetKeypair(secret_key string) (*keypair.Full, *rest_errors.RestErr) {
	micaKeypair, err := keypair.ParseFull(secret_key)
	if err != nil {
		logger.Error("Error while parsing mica keypair", err)
		return nil, rest_errors.NewInternalServerError("Error while parsing mica keypair", err)
	}
	return micaKeypair, nil
}

func getAccount(address string) (account horizon.Account, rest_err *rest_errors.RestErr) {
	// Get the current state of Tennant account from the network
	accountRequest := horizonclient.AccountRequest{AccountID: address}
	account, err := getClient().AccountDetail(accountRequest)
	if err != nil {
		rest_err = restError(err)
	}
	return
}

func CreateAccount(creatorKeypair *keypair.Full, amount string) (*keypair.Full, *rest_errors.RestErr) {
	keyPair, err := keypair.Random()
	if err != nil {
		logger.Error("Error while generate keypair", err)
		return nil, rest_errors.NewInternalServerError("Error while creating account", err)
	}

	operation := txnbuild.CreateAccount{
		Destination: creatorKeypair.Address(),
		Amount:      amount,
	}
	account, rest_err := getAccount(creatorKeypair.Address())
	if rest_err != nil {
		return nil, rest_err
	}

	tx, rest_err := buildTransaction(account, []txnbuild.Operation{&operation})
	tx, err = tx.Sign(getPassphrase(), creatorKeypair)
	if err != nil {
		return nil, restError(err)
	}
	submitTransaction(tx)
	return keyPair, nil
}

func FundAccount(address string, amount float64) *rest_errors.RestErr {
	if os.Getenv("ENV") != "production" {
		getClient().Fund(address)
		return nil
	}

	// TODO: Connect with Bitso or another exchange to fund the account
	return nil
}

func SetMultiSign(scrowKeypair *keypair.Full, signers []*keypair.Full) *rest_errors.RestErr {
	escrowAccount, _ := getAccount(scrowKeypair.Address())

	operations := []txnbuild.Operation{
		&txnbuild.SetOptions{
			MasterWeight:    txnbuild.NewThreshold(0),
			LowThreshold:    txnbuild.NewThreshold(2),
			MediumThreshold: txnbuild.NewThreshold(2),
			HighThreshold:   txnbuild.NewThreshold(2),
		},
	}
	for _, signer := range signers {
		operations = append(operations, &txnbuild.SetOptions{
			Signer: &txnbuild.Signer{
				Address: signer.Address(),
				Weight:  1,
			},
		})
	}

	tx, _ := buildTransaction(escrowAccount, operations)
	tx, err := tx.Sign(getPassphrase(), scrowKeypair)
	if err != nil {
		return restError(err)
	}

	submitTransaction(tx)
	logger.Info("Escrow account multisign seted")
	return nil
}
