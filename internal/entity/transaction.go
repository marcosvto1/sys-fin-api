package entity

import "time"

type Transaction struct {
	ID               int       `json:"id"`
	Description      string    `json:"description"`
	ShortDescription string    `json:"short_description"`
	TransactionType  string    `json:"transaction_type"`
	Amount           float64   `json:"amount"`
	CategoryId       int       `json:"category_id"`
	Category         Category  `json:"category"`
	TransactionAt    time.Time `json:"transaction_at"`
	WalletId         int       `json:"wallet_id"`
	Wallet           Wallet    `json:"wallet"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Paid             bool      `json:"paid"`
}

func NewTransaction(amount float64, transactionType, description, short_description string, transactionAt time.Time, categoryId, walletId int) *Transaction {
	return &Transaction{
		TransactionType:  transactionType,
		Description:      description,
		ShortDescription: short_description,
		Amount:           amount,
		CategoryId:       categoryId,
		WalletId:         walletId,
		TransactionAt:    transactionAt,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Time{},
		Paid:             false,
	}
}
