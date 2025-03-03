package usecase

import (
	"gitlab.com/marcosvto/sys-fin-api/internal/entity"
	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
)

type FindSubscriptionUC struct {
	SubscriptionRepository database.SubscriptionRepository
}

func NewFindSubscriptionUseCase(subscriptionRepository database.SubscriptionRepository) *FindSubscriptionUC {
	return &FindSubscriptionUC{
		SubscriptionRepository: subscriptionRepository,
	}
}

func (f *FindSubscriptionUC) FindAll() ([]entity.Subscription, error) {
	return f.SubscriptionRepository.FindAll()
}

func (f *FindSubscriptionUC) FindById(id int) (entity.Subscription, error) {
	return f.SubscriptionRepository.FindById(id)
}
