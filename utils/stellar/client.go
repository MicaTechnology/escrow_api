package stellar

import (
	"os"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/network"
)

func getClient() *horizonclient.Client {
	// TODO: Add a switch to select between testnet and public network
	if os.Getenv("ENV") != "production" {
		return horizonclient.DefaultTestNetClient
	}

	return horizonclient.DefaultPublicNetClient
}

func getPassphrase() string {
	if os.Getenv("ENV") != "production" {
		return network.TestNetworkPassphrase
	}

	return network.PublicNetworkPassphrase
}
