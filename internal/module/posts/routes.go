package posts

import (
	"github.com/bypepe77/api/ent"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserPostController struct {
	log            *logrus.Logger
	db             *ent.Client
	r              gin.RouterGroup
	postController PostControllerInterface
}

func NewPostRouteController(log *logrus.Logger, db *ent.Client, r gin.RouterGroup) *UserPostController {
	return &UserPostController{
		log:            log,
		db:             db,
		r:              r,
		postController: NewPostController(db),
	}
}

func (urc *UserPostController) RegisterPostRoutes() {
	urc.r.GET("/test", urc.postController.GetPost)
}
