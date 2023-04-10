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
