package service

import (
	"user-transaction-service/pkg/repository"
)

type TransactionMakerService struct {
	repo repository.TransactionMaker
}

func NewTransactionMakerService(repo repository.TransactionMaker) *TransactionMakerService {
	return &TransactionMakerService{repo: repo}
}

func (m *TransactionMakerService) MakeTransaction(id int64, amount int64) (string, error) {
	return m.repo.MakeTransaction(id, amount)
}
