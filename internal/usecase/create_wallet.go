package usecase

import (
	"errors"

	"gitlab.com/marcosvto/sys-adv-api/internal/entity"
	"gitlab.com/marcosvto/sys-adv-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-adv-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-adv-api/pkg/errorable"

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

func (uc *CreateWalletUseCase) Execute(input dtos.CreateWalletInput) (dtos.CreateWalletOutput, error) {

	_, err := uc.UserRepository.FindById(input.UserId)
	if err != nil {
		log.Error(err)
		return dtos.CreateWalletOutput{}, errors.New(errorable.NOT_FOUND_REGISTER)
	}

	wallet := entity.NewWallet(-1, input.Name, input.InitialAmount, input.UserId)
	err = uc.WalletRepository.Create(wallet)
	if err != nil {
		log.Error(err)
		return dtos.CreateWalletOutput{}, errors.New(errorable.FAILED_TO_CREATE_WALLET)
	}

	return dtos.CreateWalletOutput{
		Id:     wallet.ID,
		Name:   wallet.Name,
		UserId: wallet.UserId,
		Amount: wallet.Amount,
	}, nil
}