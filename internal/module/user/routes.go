package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRouteController struct {
	log *logrus.Logger
	db  gorm.DB
	r   gin.RouterGroup
}

func NewUserRouteController(log *logrus.Logger, db gorm.DB, r gin.RouterGroup) *UserRouteController {
	return &UserRouteController{
		log: log,
		db:  db,
		r:   r,
	}
}

func (urc *UserRouteController) RegisterUserRoutes() {
	fmt.Println("register routes")
	urc.r.GET("/test", printsomething)
}

func printsomething(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Test"})
}
