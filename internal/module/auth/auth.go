package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthControllerInterface interface {
	CreateUser(c *gin.Context)
}

type AuthController struct {
	DB         gorm.DB
	repository AuthRepositoryInterface
}

func NewAuthController(db gorm.DB) AuthControllerInterface {
	return &AuthController{
		repository: NewAuthRepository(db),
	}
}

func (ctrl *AuthController) CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "This is a create user"})
}
