package usecase

import (
	"errors"

	"gitlab.com/marcosvto/sys-fin-api/internal/entity"
	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-fin-api/pkg/errorable"

	log "github.com/sirupsen/logrus"
)

type CreateWalletUseCase struct {
	WalletRepository database.IWalletRepository
	UserRepository   database.IUserRepository
}

func NewCreateWalletUseCase(walletRepository database.IWalletRepository, userRepository database.IUserRepository) *CreateWalletUseCase {
	return &CreateWalletUseCase{
		WalletRepository: walletRepository,
		UserRepository:   userRepository,
	}
}

func (uc *CreateWalletUseCase) Execute(input dtos.CreateWalletInput) (dtos.WalletOutput, error) {

	_, err := uc.UserRepository.FindById(input.UserId)
	if err != nil {
		log.Error(err)
		return dtos.WalletOutput{}, errors.New(errorable.NOT_FOUND_REGISTER)
	}

	wallet := entity.NewWallet(-1, input.Name, input.InitialAmount, input.UserId)
	err = uc.WalletRepository.Create(wallet)
	if err != nil {
		log.Error(err)
		return dtos.WalletOutput{}, errors.New(errorable.FAILED_TO_CREATE_WALLET)
	}

	return dtos.WalletOutput{
		ID:     wallet.ID,
		Name:   wallet.Name,
		UserId: wallet.UserId,
		Amount: wallet.Amount,
	}, nil
}
