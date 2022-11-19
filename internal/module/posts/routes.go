package posts

import (
	"github.com/bypepe77/api/internal/module/posts/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserPostController struct {
	log            *logrus.Logger
	db             gorm.DB
	r              gin.RouterGroup
	postController controller.PostController
}

func NewPostRouteController(log *logrus.Logger, db gorm.DB, r gin.RouterGroup) *UserPostController {
	return &UserPostController{
		log:            log,
		db:             db,
		r:              r,
		postController: *controller.NewPostController(),
	}
}

func (urc *UserPostController) RegisterPostRoutes() {
	urc.r.GET("/test", urc.postController.GetPost)
}
