package usecase

import (
	"errors"
	"log"

	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
)

type DeleteSubscriptionUseCase struct {
	SubscriptionRepository *database.SubscriptionRepository
}

func NewDeleteSubscriptionUsecase(repository *database.SubscriptionRepository) *DeleteSubscriptionUseCase {
	return &DeleteSubscriptionUseCase{
		SubscriptionRepository: repository,
	}
}

func (u *DeleteSubscriptionUseCase) Execute(id int) error {
	_, err := u.SubscriptionRepository.FindById(id)
	if err != nil {
		log.Println(err)
		return errors.New("subscription not found")
	}

	err = u.SubscriptionRepository.DeleteById(id)
	if err != nil {
		log.Println(err)
		return errors.New("failed to delete subscription")
	}

	return nil
}
