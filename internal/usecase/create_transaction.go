package usecase

import (
	"errors"
	"time"

	"gitlab.com/marcosvto/sys-fin-api/internal/entity"
	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-fin-api/pkg/errorable"

	log "github.com/sirupsen/logrus"
)

type CreateTransactionUseCase struct {
	TransactionRepository database.ITransactionRepository
}

func NewCreateTransactionUseCase(transactionRepository database.ITransactionRepository) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionRepository: transactionRepository,
	}
}

func (u *CreateTransactionUseCase) Execute(input dtos.CreateTransactionInput) (dtos.TransactionOutput, error) {

	transactionAt, _ := time.Parse("2006-01-02", input.TransactionAt)
	transaction := entity.NewTransaction(
		input.Amount,
		input.TransactionType,
		input.Description,
		input.ShortDescription,
		transactionAt,
		input.CategoryId,
		input.WalletId,
	)

	err := u.TransactionRepository.Create(transaction)
	if err != nil {
		log.Error(err)
		return dtos.TransactionOutput{}, errors.New(errorable.FAILED_TO_CREATE_TRANSACTION)
	}

	updatedAt := ""
	if !transaction.UpdatedAt.IsZero() {
		updatedAt = transaction.UpdatedAt.Format("2006-01-02")
	}

	return dtos.TransactionOutput{
		ID:               transaction.ID,
		Description:      transaction.Description,
		ShortDescription: transaction.ShortDescription,
		TransactionType:  transaction.TransactionType,
		Paid:             transaction.Paid,
		TransactionAt:    transaction.TransactionAt.Format("2006-01-02"),
		Amount:           transaction.Amount,
		CategoryId:       transaction.CategoryId,
		CategoryName:     transaction.Category.Name,
		WalletId:         transaction.WalletId,
		CreatedAt:        transaction.CreatedAt.Format("2006-01-02"),
		UpdatedAt:        updatedAt,
	}, nil
}
