package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthRouteController struct {
	log            *logrus.Logger
	db             gorm.DB
	r              gin.RouterGroup
	authController AuthControllerInterface
}

func NewAuthRouteController(log *logrus.Logger, db gorm.DB, r gin.RouterGroup) *AuthRouteController {
	return &AuthRouteController{
		log:            log,
		db:             db,
		r:              r,
		authController: NewAuthController(db),
	}
}

func (ctrl *AuthRouteController) RegisterPostRoutes() {
	ctrl.r.GET("/test", ctrl.authController.CreateUser)
}
