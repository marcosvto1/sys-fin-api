package entity

type Wallet struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	UserId int     `json:"user_id"`
	User   User    `json:"user"`
}

func NewWallet(id int, name string, amount float64, userId int) *Wallet {
	return &Wallet{
		ID:     id,
		Name:   name,
		Amount: amount,
		UserId: userId,
		User:   User{},
	}
}
