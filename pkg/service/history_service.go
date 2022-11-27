package service

import (
	"user-transaction-service/pkg/model"
	"user-transaction-service/pkg/repository"
)

type HistoryReaderService struct {
	repo repository.HistoryReader
}

func NewHistoryReaderService(repo repository.HistoryReader) *HistoryReaderService {
	return &HistoryReaderService{repo: repo}
}

func (r *HistoryReaderService) GetAllHistory() ([]model.History, error) {
	return r.repo.GetAllHistory()
}
func (r *HistoryReaderService) GetHistoryByUserId(id int64) ([]model.History, error) {
	return r.repo.GetHistoryByUserId(id)
}
