package services

import (
	"github.com/MicaTechnology/escrow_api/domains/escrows"
	"github.com/MicaTechnology/escrow_api/utils/rest_errors"
)

var (
	ItemsService escrowsServiceInterface = &escrowsService{}
)

type escrowsServiceInterface interface {
	Create(escrow escrows.Escrow) (*escrows.Escrow, *rest_errors.RestErr)
}

type escrowsService struct{}

func (s *escrowsService) Create(escrow escrows.Escrow) (*escrows.Escrow, *rest_errors.RestErr) {
	if err := escrow.Save(); err != nil {
		return nil, err
	}
	return &escrow, nil
}
