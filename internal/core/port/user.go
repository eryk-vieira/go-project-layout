package port

import (
	"github.com/eryk-vieira/go-api-project-layout/internal/core/dto"
	"github.com/eryk-vieira/go-api-project-layout/internal/core/model"
)

// Primary Port
type UserService interface {
	Create(user *dto.CreateUser) (*model.User, error)
}

// Secondary Port
type UserRepository interface {
	Create(user *model.User) error
}
