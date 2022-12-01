package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRouteController struct {
	log            *logrus.Logger
	db             gorm.DB
	r              gin.RouterGroup
	userController UserControllerInterface
}

func NewUserRouteController(log *logrus.Logger, db gorm.DB, r gin.RouterGroup) *UserRouteController {
	return &UserRouteController{
		log:            log,
		db:             db,
		r:              r,
		userController: NewUserController(db),
	}
}

func (urc *UserRouteController) RegisterUserRoutes() {
	urc.r.POST("/create", urc.userController.Create)
	urc.r.GET("/:id", urc.userController.GetUser)
}
