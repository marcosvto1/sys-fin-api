package usecase

import (
	"errors"
	"time"

	"gitlab.com/marcosvto/sys-adv-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-adv-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-adv-api/pkg/errorable"

	log "github.com/sirupsen/logrus"
)

type UpdateTransactoinUseCase struct {
	TransactionRepository database.ITransactionRepository
}

func NewUpdateTransactoinUseCase(transactionRepository database.ITransactionRepository) *UpdateTransactoinUseCase {
	return &UpdateTransactoinUseCase{
		TransactionRepository: transactionRepository,
	}
}

func (uc *UpdateTransactoinUseCase) Execute(id int, input dtos.UpdateTransactionInput) error {
	transaction, err := uc.TransactionRepository.FindById(id)
	if err != nil {
		return errors.New(errorable.NOT_FOUND_REGISTER)
	}

	transaction.Description = input.Description
	transaction.Amount = input.Amount
	transaction.CategoryId = input.CategoryId
	transaction.WalletId = input.WalletId
	transactionAt, _ := time.Parse("2006-01-02", input.TransactionAt)
	transaction.TransactionAt = transactionAt
	transaction.UpdatedAt = time.Now()
	transaction.Paid = input.Paid

	err = uc.TransactionRepository.Update(transaction)
	if err != nil {
		log.Error(err)
		return errors.New(errorable.FAILED_TO_UPDATE_TRANSACTION)
	}

	return nil
}
