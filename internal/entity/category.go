package entity

import "time"

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCategory(name string) *Category {
	return &Category{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}
}
