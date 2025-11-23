package database

import "database/sql"
import "gitlab.com/marcosvto/sys-fin-api/internal/entity"
import log "github.com/sirupsen/logrus"

type DebtRepository struct {
	DB *sql.DB
}

func NewDebtRepository(db *sql.DB) *DebtRepository {
	return &DebtRepository{DB: db}
}

func (r *DebtRepository) Create(debt *entity.Debt) (*entity.Debt, error) {
	query := "INSERT INTO debts (amount, category_id, date_debit) VALUES ($1, $2, $3) RETURNING id"

	log.Info("Executing query: ", query)
	log.Infof("With values: amount=%f, category_id=%d, date_debit=%s", debt.Amount, debt.CategoryId, debt.Date)		

	var id uint
	err := r.DB.QueryRow(query, debt.Amount, debt.CategoryId, debt.Date).Scan(&id)
	if err != nil {
		log.Error("Error executing query: ", err)
		return nil, err
	}
	debt.ID = id
	return debt, nil
}