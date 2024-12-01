package repositories

import (
	"database/sql"

	"github.com/pteus/books-api/internal/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByUsername(username string) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) Create(user *models.User) error {
	_, err := u.db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	return err
}

func (u *userRepository) FindByUsername(username string) (*models.User, error) {
	row := u.db.QueryRow("SELECT id, username, password FROM users where username = $1", username)

	user := &models.User{}
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		return nil, err
	}

	return user, nil
}
