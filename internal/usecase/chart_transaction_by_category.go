package usecase

import (
	"fmt"
	"time"

	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
)

type ChartTransactionByCategoryUseCase struct {
	TransactionRepository database.ITransactionRepository
}

func NewChartTransactionByCategoryUseCase(transactionRepository database.ITransactionRepository) *ChartTransactionByCategoryUseCase {
	return &ChartTransactionByCategoryUseCase{
		TransactionRepository: transactionRepository,
	}
}

func (u *ChartTransactionByCategoryUseCase) Execute(input dtos.GetChartTransactionByCategoryInput) ([]map[string]any, error) {

	if input.Month == "" || input.Year == "" {
		input.Month = time.Now().AddDate(0, -1, 0).Format("01")
		input.Year = time.Now().Format("2006")
	}

	fmt.Println(input.Month, input.Year)

	res, err := u.TransactionRepository.GetChartTransactionByCategory(input.Month, input.Year)

	if err != nil {
		return nil, err
	}

	return res, nil
}
