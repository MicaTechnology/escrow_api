package stellar

import (
	"github.com/MicaTechnology/escrow_api/utils/logger"
	"github.com/MicaTechnology/escrow_api/utils/rest_errors"
	"github.com/stellar/go/clients/horizonclient"
)

func restError(err error) *rest_errors.RestErr {
	if p := horizonclient.GetError(err); p != nil {
		logger.GetLogger().Printf("  Info: %s\n", p.Problem)
		if results, ok := p.Problem.Extras["result_codes"]; ok {
			logger.GetLogger().Printf("  Extras: %s\n", results)
		}
	}
	return rest_errors.NewInternalServerError("error connecting to stellar", err)
}
