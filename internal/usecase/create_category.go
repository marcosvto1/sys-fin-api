package usecase

import (
	"gitlab.com/marcosvto/sys-adv-api/internal/entity"
	"gitlab.com/marcosvto/sys-adv-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-adv-api/internal/usecase/dtos"
)

type CreateCategoryUseCase struct {
	CategoryRepository database.ICategoryRepository
}

func NewCreateCategoryUseCase(categoryRepository database.ICategoryRepository) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		CategoryRepository: categoryRepository,
	}
}

func (this *CreateCategoryUseCase) Execute(input dtos.CreateCategoryInput) (dtos.CategoryOutput, error) {
	category := entity.NewCategory(input.Name)

	if err := this.CategoryRepository.Create(category); err != nil {
		return dtos.CategoryOutput{}, err
	}

	return dtos.CategoryOutput{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Format("2006-01-02"),
		UpdatedAt: "",
	}, nil
}
