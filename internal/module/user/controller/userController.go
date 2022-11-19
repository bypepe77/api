package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) GetUser(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"data": "Test"})

}
