package dtos

type CreateWalletInput struct {
	UserId        int     `json:"user_id"`
	Name          string  `json:"name"`
	InitialAmount float64 `json:"initial_amount"`
}

type CreateWalletOutput struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	UserId int     `json:"user_id"`
	Amount float64 `json:"amount"`
}
