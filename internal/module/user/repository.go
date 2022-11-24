package user

import (
	"errors"

	"github.com/bypepe77/api/internal/common/models"
	"gorm.io/gorm"
)

type Params struct {
	key   string
	value string
}

type UserRepositoryInterface interface {
	Create(user *models.User) (*models.User, error)
	GetBy(param *Params) (*models.User, error)
}

type UserRepository struct {
	DB gorm.DB
}

func NewUserRespository(db gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}

func (repository *UserRepository) Create(user *models.User) (*models.User, error) {
	userCreated := repository.DB.Create(&user)
	if userCreated.Error != nil {
		return nil, errors.New("Cannot create user")
	}

	return user, nil
}

func (repository *UserRepository) GetBy(param *Params) (*models.User, error) {
	user := &models.User{}
	where := BuildQueryWithOneValue(param)

	queryUser := repository.DB.Where(where).First(&user)
	if queryUser.Error != nil {
		return nil, errors.New("Cannot find user with this clauses")
	}

	return user, nil
}

func BuildQueryWithOneValue(param *Params) map[string]interface{} {
	m := make(map[string]interface{})
	m[param.key] = param.value

	return m
}
