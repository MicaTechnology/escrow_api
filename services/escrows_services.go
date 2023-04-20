package services

import (
	"os"
	"strconv"

	"github.com/MicaTechnology/escrow_api/domains/escrows"
	"github.com/MicaTechnology/escrow_api/repository"
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
	Get(id string) (*escrows.Escrow, *rest_errors.RestErr)
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
	escrow.Tenant.Address = tenantSignerKeyPair.Address()
	escrow.Tenant.SetSecret(tenantSignerKeyPair.Seed())

	landlordSignerKeyPair, rest_err := stellar.CreateAccount(micaKeypair, stellar.MinBalance)
	if rest_err != nil {
		return nil, rest_err
	}
	escrow.Landlord.Address = landlordSignerKeyPair.Address()
	escrow.Landlord.SetSecret(landlordSignerKeyPair.Seed())

	escrowKeypair, _ := stellar.CreateAccount(micaKeypair, strconv.FormatFloat(escrow.Amount, 'f', 2, 64))
	escrow.Address = escrowKeypair.Address()

	stellar.SetMultiSign(escrowKeypair, []*keypair.Full{tenantSignerKeyPair, landlordSignerKeyPair, micaKeypair})
	logger.GetLogger().Printf("Mica public key %s", micaKeypair.Address())
	if err := repository.GetEscrowRepository().Create(&escrow); err != nil {
		return nil, err
	}

	return &escrow, nil
}

func (s *escrowsService) Get(id string) (*escrows.Escrow, *rest_errors.RestErr) {
	escrow, err := repository.GetEscrowRepository().Get(id)
	if err != nil {
		return nil, err
	}
	return escrow, nil
}
