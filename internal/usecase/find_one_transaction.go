package usecase

import (
	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"

	log "github.com/sirupsen/logrus"
)

type FindOneTransactionUseCase struct {
	TransactionRepository database.ITransactionRepository
}

func NewFindOneTransactionUseCase(userRepository database.ITransactionRepository) *FindOneTransactionUseCase {
	return &FindOneTransactionUseCase{
		TransactionRepository: userRepository,
	}
}

func (uc *FindOneTransactionUseCase) Execute(id int) (dtos.FindOutput[dtos.TransactionOutput], error) {

	log.Info("aa")
	transaction, err := uc.TransactionRepository.FindById(id)
	if err != nil {
		log.Error(err)
		return dtos.FindOutput[dtos.TransactionOutput]{}, err
	}

	updatedAt := ""
	if !transaction.UpdatedAt.IsZero() {
		updatedAt = transaction.UpdatedAt.Format("2006-01-02")
	}

	var output []dtos.TransactionOutput
	output = append(output, dtos.TransactionOutput{
		ID:              transaction.ID,
		Description:     transaction.Description,
		Amount:          transaction.Amount,
		CategoryId:      transaction.Category.ID,
		CategoryName:    transaction.Category.Name,
		WalletId:        transaction.Wallet.ID,
		WalletName:      transaction.Wallet.Name,
		TransactionType: transaction.TransactionType,
		Paid:            transaction.Paid,
		TransactionAt:   transaction.TransactionAt.Format("2006-01-02"),
		CreatedAt:       transaction.CreatedAt.Format("2006-01-02"),
		UpdatedAt:       updatedAt,
	})

	return dtos.FindOutput[dtos.TransactionOutput]{
		Items: output,
		Meta: dtos.MetaFindOutput{
			Paginate: dtos.PaginateOutput{
				PageNumber:     1,
				PageSize:       1,
				TotalPages:     1,
				TotalRegisters: 1,
			},
		},
	}, nil
}
