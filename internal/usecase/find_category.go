package usecase

import (
	"fmt"

	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
)

type FindCategoriesUseCase struct {
	CategoryRepository database.ICategoryRepository
}

func NewFindCategoriesUseCase(repository database.ICategoryRepository) *FindCategoriesUseCase {
	return &FindCategoriesUseCase{
		CategoryRepository: repository,
	}
}

func (usecase *FindCategoriesUseCase) Execute() ([]dtos.CategoryOutput, error) {
	categories, err := usecase.CategoryRepository.FindAll()
	if err != nil {
		return []dtos.CategoryOutput{}, err
	}

	fmt.Println("asdasd")

	var categoriesOutput []dtos.CategoryOutput
	for _, category := range categories {
		dto := dtos.CategoryOutput{
			ID:   category.ID,
			Name: category.Name,
		}
		categoriesOutput = append(categoriesOutput, dto)
	}

	return categoriesOutput, nil
}
