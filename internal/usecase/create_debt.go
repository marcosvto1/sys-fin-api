package usecase

import (
	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-fin-api/internal/entity"
	log "github.com/sirupsen/logrus"
)

type CreateDebt struct {
	debtRepository database.IDebtRepository
	categoryRepository database.ICategoryRepository
}

func NewCreateDebt(debtRepo database.IDebtRepository, categoryRepo database.ICategoryRepository) *CreateDebt {
	return &CreateDebt{
		debtRepository: debtRepo,
		categoryRepository: categoryRepo,
	}
}

func (uc *CreateDebt) Execute(debtDTO dtos.CreateDebtInput) (*dtos.DebtOutput, error) {
    category, err := uc.categoryRepository.FindById(debtDTO.CategoryId)
	if err != nil {
		log.Error("Error fetching category: ", err)
		return nil, err
	}
	
	debt := entity.NewDebt(0, debtDTO.Amount, debtDTO.CategoryId, debtDTO.Date)
	createdDebt, err := uc.debtRepository.Create(debt)
	if err != nil {
		log.Error("Error creating debt: ", err)
		return nil, err
	}

	debtOutput := dtos.NewDebtOutput(createdDebt.ID, createdDebt.Amount, category.Name, createdDebt.CategoryId, createdDebt.Date)

	return debtOutput, nil
}