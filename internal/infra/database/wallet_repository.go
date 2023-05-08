package database

import (
	"database/sql"

	"gitlab.com/marcosvto/sys-adv-api/internal/entity"
)

type WalletRepository struct {
	DB *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{
		DB: db,
	}
}

func (r *WalletRepository) Create(wallet *entity.Wallet) error {
	err := r.DB.QueryRow("INSERT INTO wallets (name, amount, user_id) VALUES ($1, $2, $3) RETURNING id", &wallet.Name, &wallet.Amount, &wallet.UserId).Scan(&wallet.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *WalletRepository) FindAll() ([]entity.Wallet, error) {
	rows, err := r.DB.Query("SELECT id, name, amount FROM wallets")
	if err != nil {
		return nil, err
	}

	var wallets []entity.Wallet
	for rows.Next() {
		wallet := entity.Wallet{}
		err = rows.Scan(&wallet.ID, &wallet.Name, &wallet.Amount)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet)
	}

	return wallets, nil
}
