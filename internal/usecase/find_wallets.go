package usecase

import (
	"gitlab.com/marcosvto/sys-adv-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-adv-api/internal/usecase/dtos"
)

type FindWalletsUseCase struct {
	WalletRepository database.IWalletRepository
}

func NewFindWalletsUseCase(repository database.IWalletRepository) *FindWalletsUseCase {
	return &FindWalletsUseCase{
		WalletRepository: repository,
	}
}

func (usecase *FindWalletsUseCase) Execute() ([]dtos.WalletOutput, error) {
	wallets, err := usecase.WalletRepository.FindAll()
	if err != nil {
		return []dtos.WalletOutput{}, err
	}

	var walletsOutput []dtos.WalletOutput
	for _, wallet := range wallets {
		dto := dtos.WalletOutput{
			ID:   wallet.ID,
			Name: wallet.Name,
		}
		walletsOutput = append(walletsOutput, dto)
	}

	return walletsOutput, nil
}
