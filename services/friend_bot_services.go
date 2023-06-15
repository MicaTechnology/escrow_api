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

const friendBotAmount = 9999.0

var FriendBotService friendBotInterface = &friendBotService{}

type friendBotInterface interface {
	AddFunds(friend_bot.FundRequest) *rest_errors.RestErr
}

type friendBotService struct{}

func (s *friendBotService) AddFunds(fundRequest friend_bot.FundRequest) *rest_errors.RestErr {
	destinationKeypair, err := keypair.Parse(fundRequest.Address)
	if err != nil {
		logger.Error("Error while parsing address", err)
		return rest_errors.NewBadRequestError("Invalid Address")
	}

	size := int(math.Ceil(fundRequest.Amount / friendBotAmount))
	logger.GetLogger().Printf("Merging %d accounts in %s", size, destinationKeypair.Address())

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
