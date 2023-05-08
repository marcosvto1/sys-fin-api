package dtos

type CreateTransactionInput struct {
	TransactionType string  `json:"transaction_type"`
	Description     string  `json:"description"`
	CategoryId      int     `json:"category_id"`
	WalletId        int     `json:"wallet_id"`
	TransactionAt   string  `json:"transaction_at"`
	Amount          float64 `json:"amount"`
}

type TransactionOutput struct {
	ID              int     `json:"id"`
	TransactionType string  `json:"transaction_type"`
	Description     string  `json:"description"`
	CategoryId      int     `json:"category_id"`
	CategoryName    string  `json:"category_name"`
	WalletId        int     `json:"wallet_id"`
	WalletName      string  `json:"wallet_name"`
	Amount          float64 `json:"amount"`
	TransactionAt   string  `json:"transaction_at"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

type FindTransactionInput struct {
	Id         int    `json:"id"`
	PageNumber int    `json:"page_number"`
	PageSize   int    `json:"page_size"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	CategoryId int    `json:"category_id"`
	WalletId   int    `json:"wallet_id"`
}
