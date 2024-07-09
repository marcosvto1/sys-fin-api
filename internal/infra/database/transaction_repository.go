package database

import (
	"database/sql"
	"fmt"

	"gitlab.com/marcosvto/sys-fin-api/internal/entity"
)

type FindTransactionOptions struct {
	Month      string
	Year       string
	WalletId   int
	CategoryId int
}

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (r *TransactionRepository) Create(transaction *entity.Transaction) error {
	query := `
	INSERT INTO
		transactions (transaction_type, amount, category_id, wallet_id, transaction_at, created_at, updated_at, description, short_description, paid)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id
	`
	row := r.DB.QueryRow(
		query,
		transaction.TransactionType,
		transaction.Amount,
		transaction.CategoryId,
		transaction.WalletId,
		transaction.TransactionAt,
		transaction.CreatedAt,
		transaction.UpdatedAt,
		transaction.Description,
		transaction.ShortDescription,
		transaction.Paid,
	)

	err := row.Scan(&transaction.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepository) Find(offset, pageSize int, filter FindTransactionOptions) ([]entity.Transaction, int, error) {
	fmt.Printf("buscando %v %v\n", offset, pageSize)
	var err error

	count := 0

	sql := `
	SELECT t.id, t.description, t.transaction_type, t.amount, t.transaction_at, t.paid, c.id, c.name, w.id, w.name
	FROM transactions t
	JOIN wallets w ON w.id = wallet_id
	JOIN categories c ON c.id = category_id
	`

	values := []interface{}{
		offset,
		pageSize,
	}

	indice := 2

	if filter.Month != "" {
		indice += 1
		sql = sql + fmt.Sprintf(` WHERE to_char(transaction_at, 'MM') = $%d`, indice)

		indice += 1
		sql = sql + fmt.Sprintf(` AND to_char(transaction_at, 'YYYY') =  $%d`, indice)
		values = append(values, filter.Month)
		values = append(values, filter.Year)
	}

	if filter.WalletId != -1 {
		indice += 1
		sql = sql + fmt.Sprintf(` AND wallet_id = $%d`, indice)
		values = append(values, filter.WalletId)
	}

	if filter.CategoryId != -1 {
		indice += 1
		sql = sql + fmt.Sprintf(` AND category_id = $%d`, indice)
		values = append(values, filter.CategoryId)
	}

	sql = sql + ` OFFSET $1 LIMIT $2`

	fmt.Println(sql)
	fmt.Println(values...)

	stmt, err := r.DB.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return nil, count, err
	}

	rows, err := stmt.Query(values...)
	if err != nil {
		fmt.Println(err)
		return nil, count, err
	}

	err = r.DB.QueryRow("select count(id) from transactions").Scan(&count)
	if err != nil {
		return nil, count, err
	}

	var transactions []entity.Transaction
	for rows.Next() {
		var transaction = entity.Transaction{}

		err = rows.Scan(
			&transaction.ID,
			&transaction.Description,
			&transaction.TransactionType,
			&transaction.Amount,
			&transaction.TransactionAt,
			&transaction.Paid,
			&transaction.Category.ID,
			&transaction.Category.Name,
			&transaction.Wallet.ID,
			&transaction.Wallet.Name,
		)
		if err != nil {
			return nil, count, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, count, nil
}

func (repo *TransactionRepository) DeleteById(id int) error {
	fmt.Println("entrou")
	stmt, err := repo.DB.Prepare("DELETE FROM transactions WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Query(id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *TransactionRepository) FindById(id int) (entity.Transaction, error) {
	var transaction entity.Transaction

	sql := `
	SELECT t.id, t.description, t.transaction_type, t.amount, t.transaction_at, t.paid , c.id, c.name, w.id, w.name
	FROM transactions t
	JOIN wallets w ON w.id = t.wallet_id
	JOIN categories c ON c.id = t.category_id  WHERE t.id = $1
	`
	stmt, err := repo.DB.Prepare(sql)
	if err != nil {
		return transaction, err
	}

	err = stmt.QueryRow(id).Scan(
		&transaction.ID,
		&transaction.Description,
		&transaction.TransactionType,
		&transaction.Amount,
		&transaction.TransactionAt,
		&transaction.Paid,
		&transaction.Category.ID,
		&transaction.Category.Name,
		&transaction.Wallet.ID,
		&transaction.Wallet.Name,
	)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (repo *TransactionRepository) Update(transaction entity.Transaction) error {
	sql := `
	UPDATE transactions
	SET description=$1, transaction_type=$2, amount=$3, category_id=$4, wallet_id=$5, transaction_at=$6, paid=$7
	WHERE id = $8
	`
	stmt, err := repo.DB.Prepare(sql)
	if err != nil {
		return nil
	}

	res, err := stmt.Exec(
		transaction.Description,
		transaction.TransactionType,
		transaction.Amount,
		transaction.CategoryId,
		transaction.WalletId,
		transaction.TransactionAt,
		transaction.Paid,
		transaction.ID,
	)
	if err != nil {
		return err
	}

	if affected, err := res.RowsAffected(); affected != 1 || err != nil {
		return err
	}

	return nil
}

// CHARTS

func (repo *TransactionRepository) GetChartTransactionByCategory(month, year string) ([]map[string]any, error) {

	sql := `
	SELECT SUM(amount) as amount, c.name as category_name
	FROM transactions t
	JOIN categories c ON c.id = category_id
	WHERE EXTRACT(MONTH FROM t.transaction_at) = $1 AND EXTRACT(YEAR FROM t.transaction_at) = $2
	GROUP BY c.name
	`

	stmt, err := repo.DB.Prepare(sql)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(month, year)
	if err != nil {
		return nil, err
	}

	var chart []map[string]any
	for rows.Next() {
		var category string
		var amount float64
		err = rows.Scan(&amount, &category)
		if err != nil {
			return nil, err
		}
		chart = append(chart, map[string]any{
			"amount":        amount,
			"category_name": category,
		})
	}

	return chart, nil
}

func (repo *TransactionRepository) GetChartTransactionByType(year string) ([]map[string]any, error) {

	sql := `
	SELECT sum(t.amount) as Total, t.transaction_type, EXTRACT(MONTH FROM t.transaction_at) as mon
	FROM transactions t
	JOIN categories c ON c.id = t.category_id
	WHERE EXTRACT(MONTH FROM t.transaction_at) >= 01 AND EXTRACT(MONTH FROM t.transaction_at) <= 12 AND EXTRACT(YEAR FROM t.transaction_at) = $1
	GROUP BY t.transaction_type, mon
	`

	stmt, err := repo.DB.Prepare(sql)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(year)
	if err != nil {
		return nil, err
	}

	var chart []map[string]any
	for rows.Next() {
		var month string
		var transactionType string
		var amount float64
		err = rows.Scan(&amount, &transactionType, &month)
		if err != nil {
			return nil, err
		}
		chart = append(chart, map[string]any{
			"amount":           amount,
			"transaction_type": transactionType,
			"month":            month,
		})
	}

	return chart, nil
}
