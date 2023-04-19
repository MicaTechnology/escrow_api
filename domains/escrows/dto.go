package escrows

type Escrow struct {
	Address    string `json:"address"`
	TennatId   string `json:"tennat_id"`
	Amount     string `json:"amount"`
	LandLordId string `json:"landlord_id"`
}
