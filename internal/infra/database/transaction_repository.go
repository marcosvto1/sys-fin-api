package database

import (
	"database/sql"
	"fmt"

	"gitlab.com/marcosvto/sys-adv-api/internal/entity"
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
	row := r.DB.QueryRow(`
	INSERT INTO 
		transactions (transaction_type, amount, category_id, wallet_id, transaction_at, created_at, updated_at, description) 
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id
	`, transaction.TransactionType, transaction.Amount, transaction.CategoryId, transaction.WalletId, transaction.TransactionAt, transaction.CreatedAt, transaction.UpdatedAt, transaction.Description)

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
	SELECT t.id, t.description, t.transaction_type, t.amount, t.transaction_at, c.id, c.name, w.id, w.name
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
	SELECT t.id, t.description, t.transaction_type, t.amount, t.transaction_at, c.id, c.name, w.id, w.name
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
