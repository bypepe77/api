package user

import (
	"github.com/bypepe77/api/internal/module/user/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRouteController struct {
	log            *logrus.Logger
	db             gorm.DB
	r              gin.RouterGroup
	userController controller.UserController
}

func NewUserRouteController(log *logrus.Logger, db gorm.DB, r gin.RouterGroup) *UserRouteController {
	return &UserRouteController{
		log:            log,
		db:             db,
		r:              r,
		userController: *controller.NewUserController(),
	}
}

func (urc *UserRouteController) RegisterUserRoutes() {
	urc.r.GET("/test", urc.userController.GetUser)
}
