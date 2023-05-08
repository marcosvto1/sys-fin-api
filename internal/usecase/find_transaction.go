package usecase

import (
	"gitlab.com/marcosvto/sys-adv-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-adv-api/internal/usecase/dtos"
)

type FindTransactionUseCase struct {
	TransactionRepository database.ITransactionRepository
}

func NewFindTransactionUseCase(userRepository database.ITransactionRepository) *FindTransactionUseCase {
	return &FindTransactionUseCase{
		TransactionRepository: userRepository,
	}
}

func (uc *FindTransactionUseCase) Execute(input dtos.FindTransactionInput) (dtos.FindOutput[dtos.TransactionOutput], error) {
	offset := (input.PageNumber - 1) * input.PageSize

	transaction, total, err := uc.TransactionRepository.Find(offset, input.PageSize, database.FindTransactionOptions{
		Month:      input.Month,
		Year:       input.Year,
		CategoryId: input.CategoryId,
		WalletId:   input.WalletId,
	})
	if err != nil {
		return dtos.FindOutput[dtos.TransactionOutput]{}, err
	}

	var output []dtos.TransactionOutput
	for _, transaction := range transaction {

		updatedAt := ""
		if !transaction.UpdatedAt.IsZero() {
			updatedAt = transaction.UpdatedAt.Format("2006-01-02")
		}

		output = append(output, dtos.TransactionOutput{
			ID:              transaction.ID,
			Description:     transaction.Description,
			Amount:          transaction.Amount,
			CategoryId:      transaction.Category.ID,
			CategoryName:    transaction.Category.Name,
			WalletId:        transaction.WalletId,
			WalletName:      transaction.Wallet.Name,
			TransactionType: transaction.TransactionType,
			TransactionAt:   transaction.TransactionAt.Format("2006-01-02"),
			CreatedAt:       transaction.CreatedAt.Format("2006-01-02"),
			UpdatedAt:       updatedAt,
		})
	}

	return dtos.FindOutput[dtos.TransactionOutput]{
		Items: output,
		Meta: dtos.MetaFindOutput{
			Paginate: dtos.PaginateOutput{
				PageNumber:     input.PageNumber,
				PageSize:       input.PageSize,
				TotalPages:     total / input.PageSize,
				TotalRegisters: total,
			},
		},
	}, nil
}
