package database

import "gitlab.com/marcosvto/sys-adv-api/internal/entity"

type IUserRepository interface {
	Create(user *entity.User) error
	Find(offset, pageNumber, id int) ([]entity.User, int, error)
	FindByEmail(email string) (entity.User, error)
	FindById(id int) (entity.User, error)
}

type IWalletRepository interface {
	Create(wallet *entity.Wallet) error
	FindAll() ([]entity.Wallet, error)
}

type ICategoryRepository interface {
	Create(category *entity.Category) error
	FindAll() ([]entity.Category, error)
}

type ITransactionRepository interface {
	Create(transaction *entity.Transaction) error
	Find(offset, pageNumber int, filter FindTransactionOptions) ([]entity.Transaction, int, error)
	Update(transaction entity.Transaction) error
	DeleteById(id int) error
	FindById(id int) (entity.Transaction, error)
}
