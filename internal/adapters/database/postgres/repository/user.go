package repository

import (
	"database/sql"

	"github.com/eryk-vieira/go-api-project-layout/internal/core/model"
	"github.com/eryk-vieira/go-api-project-layout/internal/core/port"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) port.UserRepository {
	return &userRepository{
		db,
	}
}

func (*userRepository) Create() (*model.User, error) {
	return nil, nil
}
