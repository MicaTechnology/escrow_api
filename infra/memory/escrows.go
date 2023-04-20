package infra

import (
	"github.com/MicaTechnology/escrow_api/domains/escrows"
	"github.com/MicaTechnology/escrow_api/utils/rest_errors"
	"github.com/google/uuid"
)

var EscrowTable = map[string]*escrows.Escrow{}

type EscrowRepository struct{}

func (r *EscrowRepository) Create(escrow *escrows.Escrow) *rest_errors.RestErr {
	escrow.Id = uuid.New().String()
	EscrowTable[escrow.Id] = escrow
	return nil
}

func (r *EscrowRepository) Get(id string) (*escrows.Escrow, *rest_errors.RestErr) {
	result := EscrowTable[id]
	if result == nil {
		return nil, rest_errors.NewNotFoundError("escrow not found")
	}
	return result, nil
}

func (r *EscrowRepository) Close() *rest_errors.RestErr {
	return nil
}
