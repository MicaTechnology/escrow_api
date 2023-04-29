package repository

import (
	"github.com/MicaTechnology/escrow_api/domains/escrows"
	"github.com/MicaTechnology/escrow_api/utils/rest_errors"
)

type EscrowRepository interface {
	Repository
	Create(escrow *escrows.Escrow) *rest_errors.RestErr
	Get(id string) (*escrows.Escrow, *rest_errors.RestErr)
	Update(escrow *escrows.Escrow) *rest_errors.RestErr
}

var implementation EscrowRepository

func SetEscrowRepository(repository EscrowRepository) {
	implementation = repository
}

func GetEscrowRepository() EscrowRepository {
	return implementation
}
