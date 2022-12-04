package user

import (
	"github.com/bypepe77/api/ent"
)

type Params struct {
	key   string
	value string
}

type UserRepositoryInterface interface {
}

type UserRepository struct {
	DB *ent.Client
}

func NewUserRespository(db *ent.Client) UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}
