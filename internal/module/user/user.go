package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserControllerInterface interface {
	GetUser(c *gin.Context)
}

type UserController struct {
	DB gorm.DB
}

func NewUserController(db gorm.DB) UserControllerInterface {
	return &UserController{
		DB: db,
	}
}

func (uc *UserController) GetUser(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"data": "Test"})

}
