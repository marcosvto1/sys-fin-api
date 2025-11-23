package entity


type Debt struct {
	ID          uint    `json:"id"`
	Amount      float64 `json:"amount" `
	Date     string   	`json:"date_debit"`
	CategoryId  int     `json:"category_id,omitempty"`
}

func NewDebt(id uint, amount float64, categoryId int, date string) *Debt {
	return &Debt{
		ID:          id,
		Amount:      amount,
		Date:        date,
		CategoryId:  categoryId,
	}
}