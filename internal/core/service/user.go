package service

import (
	"github.com/eryk-vieira/go-api-project-layout/internal/core/model"
	"github.com/eryk-vieira/go-api-project-layout/internal/core/port"
)

type userService struct {
	repository port.UserRepository
}

func NewUserService(repository port.UserRepository) port.UserService {
	return &userService{
		repository,
	}
}

func (service *userService) Create() (*model.User, error) {
	user, err := service.repository.Create()

	if err != nil {
		return nil, err
	}

	return user, nil
}
