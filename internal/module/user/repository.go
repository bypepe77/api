package user

import (
	"context"
	"fmt"

	"github.com/bypepe77/api/ent"
	"github.com/bypepe77/api/ent/user"
)

type Params struct {
	key   string
	value string
}

type UserRepositoryInterface interface {
	Create() (*ent.User, error)
	Get(username string) (*ent.User, error)
}

type UserRepository struct {
	DB *ent.Client
}

func NewUserRespository(db *ent.Client) UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}

func (repository *UserRepository) Create() (*ent.User, error) {
	user, err := repository.DB.User.Create().
		SetAge(30).
		SetName("Mario").
		Save(context.Background())
	if err != nil {
		fmt.Print("error", err)
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	return user, nil
}

func (repository *UserRepository) Get(username string) (*ent.User, error) {
	user, err := repository.DB.User.
		Query().
		Where(user.Name(username)).
		Only(context.Background())
	if err != nil {
		fmt.Print("error", err)
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return user, nil
}
