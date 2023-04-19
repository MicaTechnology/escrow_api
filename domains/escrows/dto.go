package escrows

type Escrow struct {
	Id         string `json:"id"`
	address    string
	TennatId   string `json:"tennat_id"`
	amount     string `json:"amount"`
	LandLordId string `json:"landlord_id"`
}
