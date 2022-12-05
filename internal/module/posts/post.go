package posts

import (
	"net/http"

	"github.com/bypepe77/api/ent"
	"github.com/gin-gonic/gin"
)

type PostControllerInterface interface {
	Create(c *gin.Context)
}

type PostController struct {
	DB *ent.Client
}

func NewPostController(db *ent.Client) PostControllerInterface {
	return &PostController{
		DB: db,
	}
}

func (uc *PostController) Create(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"data": "This is a post"})

}
