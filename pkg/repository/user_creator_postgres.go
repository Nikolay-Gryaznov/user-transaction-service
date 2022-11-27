package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"user-transaction-service/pkg/model"
)

type UserCreatorPostgres struct {
	db *sqlx.DB
}

func NewUserCreatorPostgres(db *sqlx.DB) *UserCreatorPostgres {
	return &UserCreatorPostgres{db: db}
}

func (c *UserCreatorPostgres) CreateUser(u model.User) (int64, error) {
	q := fmt.Sprintf("INSERT INTO %s (full_name, balance) VALUES ($1, $2) RETURNING id", userTable)
	row := c.db.QueryRow(q, u.FullName, u.Balance)
	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
