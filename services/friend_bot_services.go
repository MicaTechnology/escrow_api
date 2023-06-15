package services

import (
	"math"

	"github.com/MicaTechnology/escrow_api/domains/friend_bot"
	"github.com/MicaTechnology/escrow_api/utils/logger"
	"github.com/MicaTechnology/escrow_api/utils/rest_errors"
	"github.com/MicaTechnology/escrow_api/utils/stellar"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/txnbuild"
)

const friendBotAmount = 9500.0

var FriendBotService friendBotInterface = &friendBotService{}

type friendBotInterface interface {
	AddFunds(fund_request friend_bot.FundRequest) *rest_errors.RestErr
}

type friendBotService struct{}

func (s *friendBotService) AddFunds(fund_request friend_bot.FundRequest) *rest_errors.RestErr {
	destinationKeypair, err := keypair.Parse(fund_request.Address)
	if err != nil {
		logger.Error("Error while parsing address", err)
		return rest_errors.NewInternalServerError("Invalid Address", err)
	}

	size := int(math.Ceil(fund_request.Amount / friendBotAmount))
	logger.GetLogger().Printf("Merge %d accounts in %s", size, destinationKeypair.Address())

	operation := txnbuild.AccountMerge{
		Destination: destinationKeypair.Address(),
	}
	for i := 0; i < size; i++ {
		rest_err := stellar.MergeAccount(operation)
		if rest_err != nil {
			return rest_err
		}
		logger.GetLogger().Printf("Account %d merged", i+1)
	}
	return nil
}
