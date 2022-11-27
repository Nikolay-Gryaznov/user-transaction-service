package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"user-transaction-service/pkg/model"
)

type HistoryReaderPostgres struct {
	db *sqlx.DB
}

func NewHistoryReaderPostgres(db *sqlx.DB) *HistoryReaderPostgres {
	return &HistoryReaderPostgres{db: db}
}

func (r *HistoryReaderPostgres) GetAllHistory() ([]model.History, error) {
	var histories []model.History
	q := fmt.Sprintf("SELECT * FROM %s", historyTable)
	err := r.db.Select(&histories, q)
	if err != nil {
		return []model.History{}, err
	}
	return histories, nil
}
func (r *HistoryReaderPostgres) GetHistoryByUserId(id int64) ([]model.History, error) {
	var histories []model.History
	q := fmt.Sprintf("SELECT * FROM %s h WHERE h.user_id = $1", historyTable)
	err := r.db.Select(&histories, q, id)
	if err != nil {
		return []model.History{}, err
	}
	return histories, nil
}
