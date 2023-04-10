package dtos

type CreateCategoryInput struct {
	Name string `json:"name"`
}

type CategoryOutput struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
