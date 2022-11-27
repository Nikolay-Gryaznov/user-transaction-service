package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"user-transaction-service/pkg/model"
)

type UserReaderPostgres struct {
	db *sqlx.DB
}

func NewUserReaderPostgres(db *sqlx.DB) *UserReaderPostgres {
	return &UserReaderPostgres{db: db}
}

func (r *UserReaderPostgres) GetUserById(id int64) (model.User, error) {
	var u model.User
	q := fmt.Sprintf("SELECT * FROM %s u WHERE u.id = $1", userTable)
	row := r.db.QueryRow(q, id)
	if err := row.Scan(&u.Id, &u.FullName, &u.Balance); err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (r *UserReaderPostgres) GetUsers() ([]model.User, error) {
	var users []model.User
	q := fmt.Sprintf("SELECT * FROM %s", userTable)
	err := r.db.Select(&users, q)
	if err != nil {
		return []model.User{}, err
	}
	return users, nil
}
