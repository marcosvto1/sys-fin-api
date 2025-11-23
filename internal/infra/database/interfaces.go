package database

import "gitlab.com/marcosvto/sys-fin-api/internal/entity"

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
	FindById(id int) (entity.Category, error)
}

type ITransactionRepository interface {
	Create(transaction *entity.Transaction) error
	Find(offset, pageNumber int, filter FindTransactionOptions) ([]entity.Transaction, int, error)
	Update(transaction entity.Transaction) error
	DeleteById(id int) error
	FindById(id int) (entity.Transaction, error)
	GetChartTransactionByCategory(month, year string) ([]map[string]any, error)
	GetChartTransactionByType(year string) ([]map[string]any, error)
}

type ISubscriptionRepository interface {
	Create(subscription *entity.Subscription) error
	Update(subscription *entity.Subscription) error
	FindAll() ([]entity.Subscription, error)
	FindById(id int) (entity.Subscription, error)
	DeleteById(id int) error
}

type IMovieRepository interface {
	Create(movie *entity.Movie) (*entity.Movie, error)
	Find(offset, pageNumber int, filter FindMovieOptions) ([]entity.Movie, int, error)
	FindAll() ([]entity.Movie, error)
	FindById(id int) (entity.Movie, error)
	Update(movie *entity.Movie) error
	DeleteById(id int) error
}


type IDebtRepository interface {
	Create(debt *entity.Debt) (*entity.Debt, error)
	// GetDebtByID(id uint) (*entity.Debt, error)
	// Update(debt *entity.Debt) error
	// DeleteById(id uint) error
	// FindAll() ([]*entity.Debt, error)
}