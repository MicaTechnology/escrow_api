package escrows

type Account struct {
	Address string `json:"address"`
	UserId  string `json:"user_id"`
	secret  string
}

type Escrow struct {
	Id           string  `json:"id"`
	Address      string  `json:"address"`
	Amount       float64 `json:"amount"`
	claimPercent float32
	claimAmount  float64
	Tenant       Account `json:"tenant"`
	Landlord     Account `json:"landlord"`
}

func (a *Account) SetSecret(secret string) {
	// TODO: encrypt secret
	a.secret = secret
}

func (a *Account) GetSecret() string {
	// TODO: decrypt secret
	return a.secret
}

func (e *Escrow) SetClaimAmount(percent float32) {
	e.claimAmount = e.Amount * float64(percent)
	e.claimPercent = percent
}

func (e *Escrow) GetClaimAmount() float64 {
	return e.claimAmount
}

func (e *Escrow) GetClaimPercent() float32 {
	return e.claimPercent
}
