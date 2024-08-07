package dtos

type CreateTransactionInput struct {
	TransactionType  string  `json:"transaction_type"`
	Description      string  `json:"description"`
	ShortDescription string  `json:"short_description"`
	CategoryId       int     `json:"category_id"`
	WalletId         int     `json:"wallet_id"`
	TransactionAt    string  `json:"transaction_at"`
	Amount           float64 `json:"amount"`
	Paid             bool    `json:"paid"`
}

type UpdateTransactionInput = CreateTransactionInput

type TransactionOutput struct {
	ID               int     `json:"id"`
	TransactionType  string  `json:"transaction_type"`
	Description      string  `json:"description"`
	ShortDescription string  `json:"short_description"`
	CategoryId       int     `json:"category_id"`
	CategoryName     string  `json:"category_name"`
	WalletId         int     `json:"wallet_id"`
	WalletName       string  `json:"wallet_name"`
	Amount           float64 `json:"amount"`
	TransactionAt    string  `json:"transaction_at"`
	Paid             bool    `json:"paid"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

type FindTransactionInput struct {
	Id         int    `json:"id"`
	PageNumber int    `json:"page_number"`
	PageSize   int    `json:"page_size"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	CategoryId int    `json:"category_id"`
	WalletId   int    `json:"wallet_id"`
	Paid       bool   `json:"paid"`
}

type GetChartTransactionByCategoryInput struct {
	Month string `json:"month"`
	Year  string `json:"year"`
}

type GetChartTransactionByTypeInput struct {
	Year string `json:"year"`
}
