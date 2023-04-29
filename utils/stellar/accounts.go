package stellar

import (
	"os"
	"strconv"

	"github.com/MicaTechnology/escrow_api/domains/escrows"
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

func MicaKeypair() (*keypair.Full, *rest_errors.RestErr) {
	if os.Getenv("ENV") == "production" {
		micaKeypair, err := keypair.ParseFull(os.Getenv("SIGNER_SECRET_KEY"))
		if err != nil {
			logger.Error("Error while parsing mica keypair", err)
			return nil, rest_errors.NewInternalServerError("Error while parsing mica keypair", err)
		}
		return micaKeypair, nil
	}

	keyPair, err := keypair.Random()
	if err != nil {
		logger.Error("Error while generate keypair", err)
		return nil, rest_errors.NewInternalServerError("Error while creating account", err)
	}
	getClient().Fund(keyPair.Address())

	return keyPair, nil
}

func getAccount(address string) (account horizon.Account, rest_err *rest_errors.RestErr) {
	accountRequest := horizonclient.AccountRequest{AccountID: address}
	account, err := getClient().AccountDetail(accountRequest)
	if err != nil {
		rest_err = restError(err)
	}
	return
}

func GenKeypair() (*keypair.Full, *rest_errors.RestErr) {
	keyPair, err := keypair.Random()
	if err != nil {
		logger.Error("Error while generate keypair", err)
		return nil, rest_errors.NewInternalServerError("Error while creating account", err)
	}
	return keyPair, nil
}

func CreateEscrowAccount(creatorKeypair *keypair.Full, escrow escrows.Escrow) (*keypair.Full, *rest_errors.RestErr) {
	logger.Info("Creating escrow account")
	escrowKeypair, rest_err := GenKeypair()
	if rest_err != nil {
		return nil, rest_err
	}

	account, rest_err := getAccount(creatorKeypair.Address())
	if rest_err != nil {
		return nil, rest_err
	}

	operations := []txnbuild.Operation{
		&txnbuild.CreateAccount{
			Destination: escrow.Tenant.Address,
			Amount:      MinBalance,
		},
		&txnbuild.CreateAccount{
			Destination: escrow.Landlord.Address,
			Amount:      MinBalance,
		},
		&txnbuild.CreateAccount{
			Destination: escrowKeypair.Address(),
			Amount:      strconv.FormatFloat(escrow.Amount, 'f', 2, 64),
		},
	}

	tx, rest_err := buildTransaction(account, operations)
	if rest_err != nil {
		return nil, rest_err
	}

	tx, err := tx.Sign(getPassphrase(), creatorKeypair)
	if err != nil {
		return nil, restError(err)
	}
	rest_err = submitTransaction(tx)
	if rest_err != nil {
		return nil, rest_err
	}
	return escrowKeypair, nil
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
	logger.Info("Setting multisign")
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

	tx, rest_err := buildTransaction(escrowAccount, operations)
	if rest_err != nil {
		return rest_err
	}

	tx, err := tx.Sign(getPassphrase(), scrowKeypair)
	if err != nil {
		return restError(err)
	}

	rest_err = submitTransaction(tx)
	if rest_err != nil {
		return rest_err
	}
	return nil
}

func ReleaseFunds(escrow *escrows.Escrow) *rest_errors.RestErr {
	logger.Info("Releasing funds")
	// TODO: Use a method to get operations
	var operations []txnbuild.Operation
	if escrow.GetClaimPercent() == 1 {
		operations = []txnbuild.Operation{
			&txnbuild.AccountMerge{
				Destination: escrow.Landlord.Address,
			},
		}
	} else {
		operations = []txnbuild.Operation{
			&txnbuild.Payment{
				Destination: escrow.Landlord.Address,
				Amount:      strconv.FormatFloat(escrow.GetClaimAmount(), 'f', 2, 64),
				Asset:       txnbuild.NativeAsset{},
			},
			&txnbuild.AccountMerge{
				Destination: escrow.Tenant.Address,
			},
		}
	}

	escrowAccount, rest_err := getAccount(escrow.Address)
	if rest_err != nil {
		return rest_err
	}

	tx, rest_err := buildTransaction(escrowAccount, operations)
	if rest_err != nil {
		return rest_err
	}

	tenantKeyPair, rest_err := GetKeypair(escrow.Tenant.GetSecret())
	if rest_err != nil {
		return rest_err
	}
	landlordKeyPair, rest_err := GetKeypair(escrow.Landlord.GetSecret())
	if rest_err != nil {
		return rest_err
	}

	tx, err := tx.Sign(getPassphrase(), tenantKeyPair, landlordKeyPair)
	if err != nil {
		return restError(err)
	}

	rest_err = submitTransaction(tx)
	if rest_err != nil {
		return rest_err
	}
	return nil
}
