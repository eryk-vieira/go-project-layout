package service

import (
	"github.com/eryk-vieira/go-api-project-layout/internal/core/dto"
	"github.com/eryk-vieira/go-api-project-layout/internal/core/model"
	"github.com/eryk-vieira/go-api-project-layout/internal/core/port"
	"github.com/google/uuid"
)

type userService struct {
	repository port.UserRepository
}

func NewUserService(repository port.UserRepository) port.UserService {
	return &userService{
		repository,
	}
}

func (service *userService) Create(*dto.CreateUser) (*model.User, error) {
	err := service.repository.Create(&model.User{
		Id:       uuid.NewString(),
		Name:     "Eryk",
		Username: "Erykeepo 3",
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}
