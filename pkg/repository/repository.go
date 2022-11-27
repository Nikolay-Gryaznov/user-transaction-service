package repository

import (
	"github.com/jmoiron/sqlx"
	model2 "user-transaction-service/pkg/model"
)

type TransactionMaker interface {
	MakeTransaction(id int64, amount int64) (string, error)
}

type UserCreator interface {
	CreateUser(u model2.User) (int64, error)
}

type UserReader interface {
	GetUserById(id int64) (model2.User, error)
	GetUsers() ([]model2.User, error)
}

type HistoryReader interface {
	GetAllHistory() ([]model2.History, error)
	GetHistoryByUserId(id int64) ([]model2.History, error)
}

type Repository struct {
	UserCreator
	UserReader
	HistoryReader
	TransactionMaker
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserCreator:      NewUserCreatorPostgres(db),
		UserReader:       NewUserReaderPostgres(db),
		TransactionMaker: NewTransactionMakerPostgres(db),
		HistoryReader:    NewHistoryReaderPostgres(db),
	}
}
