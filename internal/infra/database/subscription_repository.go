package database

import (
	"database/sql"

	"gitlab.com/marcosvto/sys-fin-api/internal/entity"
)

type SubscriptionRepository struct {
	DB *sql.DB
}

func NewSubscriptionRepository(db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		DB: db,
	}
}

func (r *SubscriptionRepository) Create(subscription *entity.Subscription) (int, error) {
	query := `
	INSERT INTO
		subscriptions (name, price)
	VALUES
		($1, $2) RETURNING id
	`
	row := r.DB.QueryRow(
		query,
		subscription.Name,
		subscription.Price,
	)

	err := row.Scan(&subscription.ID)
	if err != nil {
		return -1, err
	}

	return subscription.ID, nil
}

func (r *SubscriptionRepository) FindById(id int) (entity.Subscription, error) {
	var subscription entity.Subscription

	query := `
	SELECT
		id, name, price
	FROM
		subscriptions
	WHERE
		id = $1
	`
	row := r.DB.QueryRow(query, id)

	err := row.Scan(
		&subscription.ID,
		&subscription.Name,
		&subscription.Price,
	)
	if err != nil {
		return entity.Subscription{}, err
	}

	return subscription, nil
}

func (r *SubscriptionRepository) FindAll() ([]entity.Subscription, error) {

	query := `
	SELECT
		id, name, price
	FROM
		subscriptions
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var subscriptions []entity.Subscription

	for rows.Next() {
		var subscription entity.Subscription

		err = rows.Scan(
			&subscription.ID,
			&subscription.Name,
			&subscription.Price,
		)
		if err != nil {
			return nil, err
		}

		subscriptions = append(subscriptions, subscription)

	}

	return subscriptions, nil
}

func (r *SubscriptionRepository) DeleteById(id int) error {
	query := `
	DELETE FROM
		subscriptions
	WHERE
		id = $1
	`

	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *SubscriptionRepository) Update(subscription *entity.Subscription) (entity.Subscription, error) {
	query := `
	UPDATE
		subscriptions
	SET
		name = $1,
		price = $2
	WHERE
		id = $3
	`

	_, err := r.DB.Exec(query, subscription.Name, subscription.Price, subscription.ID)
	if err != nil {
		return entity.Subscription{}, err
	}

	sub, err := r.FindById(subscription.ID)
	if err != nil {
		return entity.Subscription{}, err
	}

	return sub, nil
}
