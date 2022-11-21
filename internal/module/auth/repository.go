package auth

import "gorm.io/gorm"

type AuthRepositoryInterface interface {
	CreateUser() string
}

type AuthRepository struct {
	DB gorm.DB
}

func NewAuthRepository(db gorm.DB) AuthRepositoryInterface {
	return &AuthRepository{
		DB: db,
	}
}

func (r *AuthRepository) CreateUser() string {
	return "User"
}
