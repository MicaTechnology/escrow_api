package escrows

type Escrow struct {
	Id           string  `json:"id"`
	Address      string  `json:"address"`
	Amount       float64 `json:"amount"`
	ClaimPercent float32 `json:"claim_percent"`
	Tenant       Account `json:"tenant"`
	Landlord     Account `json:"landlord"`
}

type Account struct {
	Address string `json:"address"`
	UserId  string `json:"user_id"`
	secret  string
}

func (a *Account) SetSecret(secret string) {
	// TODO: encrypt secret
	a.secret = secret
}
