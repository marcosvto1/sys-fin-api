package usecase

import (
	"errors"

	"gitlab.com/marcosvto/sys-fin-api/internal/entity"
	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
)

type UpdateSubscriptionUsecase struct {
	SubscriptionRepository *database.SubscriptionRepository
}

func NewUpdateSubscriptionUsecase(repo *database.SubscriptionRepository) *UpdateSubscriptionUsecase {
	return &UpdateSubscriptionUsecase{
		SubscriptionRepository: repo,
	}
}

func (u *UpdateSubscriptionUsecase) Execute(id int, input dtos.UpdateSubscriptionInput) (entity.Subscription, error) {
	entityEmpty := entity.Subscription{}

	entity, err := u.SubscriptionRepository.FindById(id)
	if err != nil {
		return entityEmpty, errors.New("subscription not found")
	}

	entity.Name = input.Name
	entity.Price = input.Price

	updated, err := u.SubscriptionRepository.Update(&entity)
	if err != nil {
		return entityEmpty, errors.New("failed to update subscription")
	}

	return updated, nil
}
