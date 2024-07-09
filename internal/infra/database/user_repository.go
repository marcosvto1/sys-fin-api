package database

import (
	"database/sql"
	"fmt"

	"gitlab.com/marcosvto/sys-fin-api/internal/entity"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(user *entity.User) error {
	err := r.DB.QueryRow("INSERT INTO users (name, email, password) VALUES($1, $2, $3) RETURNING id", user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindById(id int) (entity.User, error) {
	var user entity.User
	err := r.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) Find(offset, pageSize, id int) ([]entity.User, int, error) {
	fmt.Printf("buscando %v %v\n", offset, pageSize)
	count := 0

	rows, err := r.DB.Query("SELECT id, name, email FROM users OFFSET $1 LIMIT $2", offset, pageSize)
	if err != nil {
		fmt.Println(err)
		return nil, count, err
	}

	err = r.DB.QueryRow("select count(id) from users").Scan(&count)
	if err != nil {
		return nil, count, err
	}

	fmt.Println(count)

	var users []entity.User
	for rows.Next() {
		var user = entity.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, count, err
		}
		users = append(users, user)
	}

	return users, count, nil
}

func (r *UserRepository) FindByEmail(email string) (entity.User, error) {
	var user = entity.User{}

	err := r.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
