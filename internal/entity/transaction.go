package entity

import "time"

type Transaction struct {
	ID              int       `json:"id"`
	TransactionType string    `json:"transaction_type"`
	Amount          float64   `json:"amount"`
	CategoryId      int       `json:"category_id"`
	Category        Category  `json:"category"`
	TransactionAt   string    `json:"transaction_at"`
	WalletId        int       `json:"wallet_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func NewTransaction(transactionType string, amount float64, categoryId int, transactionAt string, walletId int) *Transaction {
	return &Transaction{
		TransactionType: transactionType,
		Amount:          amount,
		CategoryId:      categoryId,
		TransactionAt:   transactionAt,
	}
}
