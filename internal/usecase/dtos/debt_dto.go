package dtos

type CreateDebtInput struct {
	Amount      float64 `json:"amount" `
	CategoryId  int    `json:"category_id"`
	Date        string  `json:"date"`	
}

type DebtOutput struct {
	ID          uint    `json:"id"`
	Amount      float64 `json:"amount" `
	Category    string  `json:"category"`
	CategoryId  int    `json:"category_id"`
	Date        string  `json:"date"`
}

func NewDebtOutput(id uint, amount float64, category string, categoryId int, date string) *DebtOutput {
	return &DebtOutput{
		ID:          id,
		Amount:      amount,
		Category:    category,
		CategoryId:  categoryId,
		Date:        date,
	}
}