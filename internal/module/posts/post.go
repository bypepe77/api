package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostControllerInterface interface {
	GetPost(c *gin.Context)
}

type PostController struct {
	DB gorm.DB
}

func NewPostController(db gorm.DB) PostControllerInterface {
	return &PostController{
		DB: db,
	}
}

func (uc *PostController) GetPost(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"data": "This is a post"})

}
