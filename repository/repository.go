package repository

import (
	"github.com/MicaTechnology/escrow_api/utils/rest_errors"
)

type Repository interface {
	Close() *rest_errors.RestErr
}
