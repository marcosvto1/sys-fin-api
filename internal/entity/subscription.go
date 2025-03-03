package entity

type Subscription struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewSubscription(id int, name string, price float64) *Subscription {
	return &Subscription{
		ID:    id,
		Name:  name,
		Price: price,
	}
}
