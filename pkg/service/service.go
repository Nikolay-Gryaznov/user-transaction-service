package service

import (
	model2 "user-transaction-service/pkg/model"
	"user-transaction-service/pkg/repository"
)

type TransactionMaker interface {
	MakeTransaction(id int64, amount int64) (string, error)
}

type HistoryReader interface {
	GetAllHistory() ([]model2.History, error)
	GetHistoryByUserId(id int64) ([]model2.History, error)
}

type UserCreator interface {
	CreateUser(u model2.User) (int64, error)
}

type UserReader interface {
	GetUserById(id int64) (model2.User, error)
	GetUsers() ([]model2.User, error)
}

type Service struct {
	UserCreator
	UserReader
	TransactionMaker
	HistoryReader
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		UserCreator:      NewUserCreatorService(r.UserCreator),
		UserReader:       NewUserReaderService(r.UserReader),
		TransactionMaker: NewTransactionMakerService(r.TransactionMaker),
		HistoryReader:    NewHistoryReaderService(r.HistoryReader),
	}
}
