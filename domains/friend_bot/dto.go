package friend_bot

import "github.com/MicaTechnology/escrow_api/utils/rest_errors"

type FundRequest struct {
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}

func (f *FundRequest) Validate() *rest_errors.RestErr {
	if f.Address == "" {
		return rest_errors.NewBadRequestError("Invalid/Missing Address")
	}
	if f.Amount <= 0 {
		return rest_errors.NewBadRequestError("Invalid/Missing Amount")
	}
	return nil
}
