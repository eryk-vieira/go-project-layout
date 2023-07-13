package port

import (
	"github.com/eryk-vieira/go-api-project-layout/internal/core/model"
)

// Primary Port
type UserService interface {
	Create() (*model.User, error)
}

// Secondary Port
type UserRepository interface {
	Create() (*model.User, error)
}
