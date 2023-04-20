package services

import (
	"os"
	"strconv"

	"github.com/MicaTechnology/escrow_api/domains/escrows"
	"github.com/MicaTechnology/escrow_api/utils/logger"
	"github.com/MicaTechnology/escrow_api/utils/rest_errors"
	"github.com/MicaTechnology/escrow_api/utils/stellar"
	"github.com/stellar/go/keypair"
)

var (
	EscrowsService escrowsServiceInterface = &escrowsService{}
)

type escrowsServiceInterface interface {
	Create(escrow escrows.Escrow) (*escrows.Escrow, *rest_errors.RestErr)
}

type escrowsService struct{}

func (s *escrowsService) Create(escrow escrows.Escrow) (*escrows.Escrow, *rest_errors.RestErr) {
	micaKeypair, rest_err := stellar.GetKeypair(os.Getenv("SIGNER_SECRET_KEY"))
	if rest_err != nil {
		return nil, rest_err
	}

	tenantSignerKeyPair, rest_err := stellar.CreateAccount(micaKeypair, stellar.MinBalance)
	if rest_err != nil {
		return nil, rest_err
	}
	// TODO: We should keep this in our database
	logger.GetLogger().Printf("Tenant public address: %s", tenantSignerKeyPair.Address())

	landlordSignerKeyPair, rest_err := stellar.CreateAccount(micaKeypair, stellar.MinBalance)
	if rest_err != nil {
		return nil, rest_err
	}
	// TODO: We should keep this in our database
	logger.GetLogger().Printf("Landlord public address: %s", landlordSignerKeyPair.Address())

	escrowKeypair, _ := stellar.CreateAccount(micaKeypair, strconv.FormatFloat(escrow.Amount, 'f', 2, 64))
	// TODO: Save escrowKeypair.Address() in our database

	stellar.SetMultiSign(escrowKeypair, []*keypair.Full{tenantSignerKeyPair, landlordSignerKeyPair, micaKeypair})

	if err := escrow.Save(); err != nil {
		return nil, err
	}
	return &escrow, nil
}
