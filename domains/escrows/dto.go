package escrows

type Escrow struct {
	Id           string `json:"id"`
	Address      string `json:"address"`
	TennatId     string `json:"tennat_id"`
	Amount       string `json:"amount"`
	LandLordId   string `json:"landlord_id"`
	ClaimPercent string `json:"claim_percent"`
}
