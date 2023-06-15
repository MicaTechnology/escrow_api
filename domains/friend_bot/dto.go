package friend_bot

type FundRequest struct {
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}
