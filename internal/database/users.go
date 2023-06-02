package database

import (
	"CRUD_Go_gin/internal/domain"
	"CRUD_Go_gin/pkg/postgres_database"
	"database/sql"
	"github.com/sirupsen/logrus"
)

type Users struct {
	db     *sql.DB
	logger *logrus.Entry
}

func NewUsers() *Users {
	logger := logrus.WithFields(logrus.Fields{"level": "internal database"})
	var u Users
	db, err := postgres_database.NewConnectionUsers()
	if err != nil {
		logger.Info("Can't init connection to DB users")
	}
	u.db = db
	u.logger = logger
	return &u
}

func (u *Users) CreateUser(user *domain.User) error {
	_, err := u.db.Exec("INSERT INTO users (first_name, last_name, email, password_enc) VALUES ($1, $2, $3, $4)",
		user.FirstName, user.LastName, user.Email, user.PasswordEnc)
	if err != nil {
		u.logger.Info("Can't create user")
	}
	return err
}

func (u *Users) GetUser(email string, password string) (*domain.User, error) {

	var user domain.User
	row := u.db.QueryRow("SELECT * FROM users WHERE email = $1 AND password_enc = $2", email, password)
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.PasswordEnc)
	if err != nil {
		u.logger.Info("Can't get user")
		return nil, err
	}
	return &user, nil
}

func (u *Users) UpdateUser(user domain.User) error {
	_, err := u.db.Exec("UPDATE users SET (first_name, last_name, email, password_enc VALUES ($1, $2, $3, $4) WHERE id = S5",
		user.FirstName, user.LastName, user.Email, user.PasswordEnc, user.Id)
	if err != nil {
		u.logger.Info("Can't update user")
	}
	return err
}

func (u *Users) DeleteUser(id int) error {
	_, err := u.db.Exec("DELETE FROM WHERE id = $1", id)
	if err != nil {
		u.logger.Info("Can't delete user")
	}
	return err
}
