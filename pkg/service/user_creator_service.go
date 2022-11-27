package service

import (
	"user-transaction-service/pkg/model"
	"user-transaction-service/pkg/repository"
)

type UserCreatorService struct {
	repo repository.UserCreator
}

func NewUserCreatorService(repo repository.UserCreator) *UserCreatorService {
	return &UserCreatorService{repo: repo}
}

func (s *UserCreatorService) CreateUser(u model.User) (int64, error) {
	return s.repo.CreateUser(u)
}
