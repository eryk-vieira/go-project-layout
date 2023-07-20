package repository

import (
	"database/sql"
	"fmt"

	"github.com/eryk-vieira/go-api-project-layout/internal/core/model"
	"github.com/eryk-vieira/go-api-project-layout/internal/core/port"
)

type userRepository struct {
	*sql.DB
}

func NewUserRepository(db *sql.DB) port.UserRepository {
	return &userRepository{
		db,
	}
}

func (db *userRepository) Create(user *model.User) error {
	query := `
      INSERT INTO users(id, name, username) 
      VALUES($1, $2, $3)
  `
	var (
		id       = user.Id
		name     = user.Name
		username = user.Username
	)

	fmt.Println(id)

	_, err := db.Exec(query, id, name, username)

	if err != nil {
		return err
	}

	return nil
}
