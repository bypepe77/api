package user

import (
	"github.com/bypepe77/api/ent"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserRouteController struct {
	log            *logrus.Logger
	db             *ent.Client
	r              gin.RouterGroup
	userController UserControllerInterface
}

func NewUserRouteController(log *logrus.Logger, db *ent.Client, r gin.RouterGroup) *UserRouteController {
	return &UserRouteController{
		log:            log,
		db:             db,
		r:              r,
		userController: NewUserController(db),
	}
}

func (urc *UserRouteController) RegisterUserRoutes() {
	urc.r.POST("/create", urc.userController.Create)
	urc.r.GET("/:username", urc.userController.GetUser)
}
