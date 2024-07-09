package usecase

import (
	"strconv"
	"time"

	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
)

type ChartTransactionByTypeUseCase struct {
	TransactionRepository database.ITransactionRepository
}

func NewChartTransactionByTypeUseCase(transactionRepository database.ITransactionRepository) *ChartTransactionByTypeUseCase {
	return &ChartTransactionByTypeUseCase{
		TransactionRepository: transactionRepository,
	}
}

type Values struct {
	Output float64 `json:"output"`
	Input  float64 `json:"input"`
}

func (u *ChartTransactionByTypeUseCase) Execute(input dtos.GetChartTransactionByTypeInput) ([]map[string]any, error) {

	if input.Year == "" {
		input.Year = time.Now().Format("2006")
	}

	res, err := u.TransactionRepository.GetChartTransactionByType(input.Year)

	var output []map[string]any
	for i := 1; i <= 12; i++ {
		res := GetIn(strconv.Itoa(i), res)
		var values Values

		if len(res) == 2 {
			if res[0]["transaction_type"] == "output" {
				values.Output = res[0]["amount"].(float64)
			}

			if res[0]["transaction_type"] == "input" {
				values.Input = res[0]["amount"].(float64)
			}

			if res[1]["transaction_type"] == "output" {
				values.Output = res[0]["amount"].(float64)
			}

			if res[1]["transaction_type"] == "input" {
				values.Input = res[0]["amount"].(float64)
			}
		} else if len(res) == 1 {
			if res[0]["transaction_type"] == "output" {
				values.Output = res[0]["amount"].(float64)
			}

			if res[0]["transaction_type"] == "input" {
				values.Input = res[0]["amount"].(float64)
			}
		}

		output = append(output, map[string]any{
			"month":    strconv.Itoa(i),
			"monthFmt": time.Month(i).String(),
			"output":   values.Output,
			"input":    values.Input,
			"balance":  values.Input - values.Output,
		})
	}

	if err != nil {
		return nil, err
	}

	return output, nil
}

func GetIn(month string, res []map[string]any) []map[string]any {
	var founded []map[string]any

	for i, _ := range res {
		if res[i]["month"] == month {
			founded = append(founded, res[i])
		}
	}

	return founded
}
