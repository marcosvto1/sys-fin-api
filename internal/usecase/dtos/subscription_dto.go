package dtos

type SubscriptionDTO struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateSubscriptionInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpdateSubscriptionInput = CreateSubscriptionInput
