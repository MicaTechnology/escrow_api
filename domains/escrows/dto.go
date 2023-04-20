package escrows

type Escrow struct {
	Id           string  `json:"id"`
	Address      string  `json:"address"`
	TennatId     string  `json:"tennat_id"`
	Amount       float64 `json:"amount"`
	LandLordId   string  `json:"landlord_id"`
	ClaimPercent float32 `json:"claim_percent"`
}
