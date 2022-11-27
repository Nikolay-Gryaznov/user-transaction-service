package service

import (
	"user-transaction-service/pkg/model"
	"user-transaction-service/pkg/repository"
)

type UserReaderService struct {
	repo repository.UserReader
}

func NewUserReaderService(repo repository.UserReader) *UserReaderService {
	return &UserReaderService{repo: repo}
}

func (s *UserReaderService) GetUserById(id int64) (model.User, error) {
	return s.repo.GetUserById(id)
}
func (s *UserReaderService) GetUsers() ([]model.User, error) {
	return s.repo.GetUsers()
}
