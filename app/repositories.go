package app

import (
	infra "github.com/MicaTechnology/escrow_api/infra/memory"
	"github.com/MicaTechnology/escrow_api/repository"
)

func StartRepositories() {
	repository.SetEscrowRepository(&infra.EscrowRepository{})
}
