package database

import (
	"database/sql"

	"gitlab.com/marcosvto/sys-adv-api/internal/entity"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: db,
	}
}

func (this *CategoryRepository) Create(category *entity.Category) error {
	row := this.DB.QueryRow("INSERT INTO category (name, created_at, updated_at) VALUES($1, $2, $3) RETURNING id, created_at, updated_at", category.Name, category.CreatedAt, category.UpdatedAt)
	err := row.Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
