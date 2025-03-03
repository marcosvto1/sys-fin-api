package usecase

import (
	"errors"

	"gitlab.com/marcosvto/sys-fin-api/internal/entity"
	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
)

type FindSubscriptionUsecase struct {
	SubscriptionRepository database.SubscriptionRepository
}

func NewFindSubscriptionUsecase(subscriptionRepository database.SubscriptionRepository) *FindSubscriptionUsecase {
	return &FindSubscriptionUsecase{
		SubscriptionRepository: subscriptionRepository,
	}
}

func (f *FindSubscriptionUsecase) FindAll() ([]entity.Subscription, error) {
	return f.SubscriptionRepository.FindAll()
}

func (f *FindSubscriptionUsecase) FindById(id int) (entity.Subscription, error) {
	subscription, err := f.SubscriptionRepository.FindById(id)
	if err != nil {
		return entity.Subscription{}, errors.New(
			"subscription not found",
		)
	}

	return subscription, nil
}
