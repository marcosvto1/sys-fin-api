package usecase

import (
	"gitlab.com/marcosvto/sys-adv-api/internal/infra/database"
)

type DeleteTransactionUseCase struct {
	TransactionRepository database.ITransactionRepository
}

func NewDeleteTransactionUsecase(repository database.ITransactionRepository) *DeleteTransactionUseCase {
	return &DeleteTransactionUseCase{
		TransactionRepository: repository,
	}
}

func (uc *DeleteTransactionUseCase) Execute(id int) error {
	_, err := uc.TransactionRepository.FindById(id)
	if err != nil {
		return err
	}

	err = uc.TransactionRepository.DeleteById(id)
	if err != nil {
		return err
	}

	return nil
}
