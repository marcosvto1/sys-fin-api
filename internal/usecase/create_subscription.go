package usecase

import (
	"gitlab.com/marcosvto/sys-fin-api/internal/entity"
	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
)

type CreateSubscriptionUC struct {
	SubscriptionRepository database.SubscriptionRepository
}

func NewCreateSubscriptionUC(subscriptionRepository database.SubscriptionRepository) *CreateSubscriptionUC {
	return &CreateSubscriptionUC{
		SubscriptionRepository: subscriptionRepository,
	}
}

func (f *CreateSubscriptionUC) Create(input dtos.CreateSubscriptionInput) (entity.Subscription, error) {
	subscription := entity.Subscription{
		Name:  input.Name,
		Price: input.Price,
	}

	id, err := f.SubscriptionRepository.Create(&subscription)
	if err != nil {
		return entity.Subscription{}, err
	}

	subscription.ID = id

	return subscription, nil
}
