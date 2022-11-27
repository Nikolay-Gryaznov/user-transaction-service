package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	"user-transaction-service/pkg/model"
)

type TransactionMakerPostgres struct {
	db *sqlx.DB
}

func NewTransactionMakerPostgres(db *sqlx.DB) *TransactionMakerPostgres {
	return &TransactionMakerPostgres{db: db}
}

func (m *TransactionMakerPostgres) MakeTransaction(id int64, amount int64) (string, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return "", err
	}
	q := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", userTable)
	row := tx.QueryRow(q, id)

	var user model.User
	err = row.Scan(&user.Id, &user.FullName, &user.Balance)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	if user.Balance+amount > 0 {
		q = fmt.Sprintf("UPDATE %s SET balance = $1 WHERE id = $2", userTable)
		_, err := tx.Exec(q, user.Balance+amount, id)
		if err != nil {
			tx.Rollback()
			return "", err
		}
		return m.event(tx, "successful", amount, user)
	} else {
		return m.event(tx, "failed", amount, user)
	}

}

func (m *TransactionMakerPostgres) event(tx *sql.Tx, result string, amount int64, user model.User) (string, error) {
	var balance int64
	if result == "successful" {
		balance = user.Balance + amount
	} else {
		balance = user.Balance
	}
	q := fmt.Sprintf("INSERT INTO %s (event, event_time, user_id) VALUES ($1, $2, $3) RETURNING id",
		historyTable)
	event := fmt.Sprintf("User %s with ID = %d make %s transaction %d, now balance = %d",
		user.FullName, user.Id, result, amount, balance)
	row := tx.QueryRow(q, event, time.Now(), user.Id)
	var id int64
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return "", err
	}
	response := fmt.Sprintf("New event with id = %d", id)
	return response, tx.Commit()
}
