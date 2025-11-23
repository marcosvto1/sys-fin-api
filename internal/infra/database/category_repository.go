package database

import (
	"database/sql"

	"gitlab.com/marcosvto/sys-fin-api/internal/entity"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: db,
	}
}

func (categoryRepo *CategoryRepository) Create(category *entity.Category) error {
	row := categoryRepo.DB.QueryRow("INSERT INTO categories (name, created_at, updated_at) VALUES($1, $2, $3) RETURNING id, created_at, updated_at", category.Name, category.CreatedAt, category.UpdatedAt)
	err := row.Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (cateoryRepo *CategoryRepository) FindAll() ([]entity.Category, error) {
	rows, err := cateoryRepo.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}

	var categories []entity.Category
	for rows.Next() {
		category := entity.Category{}
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (categoryRepo *CategoryRepository) FindById(id int) (entity.Category, error) {
	row := categoryRepo.DB.QueryRow("SELECT id, name FROM categories WHERE id = $1", id)

	category := entity.Category{}
	err := row.Scan(&category.ID, &category.Name)
	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}