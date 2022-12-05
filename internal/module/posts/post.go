package posts

import (
	"net/http"
	"strconv"

	"github.com/bypepe77/api/ent"
	"github.com/gernest/mention"
	"github.com/gin-gonic/gin"
)

const (
	ErrorCreatingPost = "Error creating post"
	ErrorPostIDNull   = "Post ID can't be blank"
	ErrorFetchingPost = "Error creating post"
)

type PostControllerInterface interface {
	Create(c *gin.Context)
	GetById(c *gin.Context)
}

type PostController struct {
	DB         *ent.Client
	repository PostsRepositoryInterface
}

func NewPostController(db *ent.Client) PostControllerInterface {
	return &PostController{
		DB:         db,
		repository: NewPostRepositoy(db),
	}
}

func (uc *PostController) Create(c *gin.Context) {
	var post *Post
	err := c.BindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorCreatingPost})
		return
	}

	tags := mention.GetTags('#', post.Content)
	parsedTags := parseTags(tags)

	postCreated, err := uc.repository.Create(parsedTags, post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorCreatingPost})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": postCreated})
}

func (uc *PostController) GetById(c *gin.Context) {
	id := c.Param("postID")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorPostIDNull})
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorPostIDNull})
	}

	post, err := uc.repository.GetById(intId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorFetchingPost})
	}

	publicPost, err := toPublicPost(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorFetchingPost})
	}

	c.JSON(http.StatusOK, gin.H{"data": publicPost})
}
