package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct {
}

func NewPostController() *PostController {
	return &PostController{}
}

func (uc *PostController) GetPost(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"data": "This is a post"})

}
