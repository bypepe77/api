package user

import (
	"net/http"

	"github.com/bypepe77/api/ent"
	"github.com/gin-gonic/gin"
)

const (
	ErrorCreatingUser    = "Error creating user"
	ErrorIDNull          = "ID can't be blank"
	ErrorGettingUserById = "User not found"
)

type UserControllerInterface interface {
	Create(c *gin.Context)
	GetUser(c *gin.Context)
}

type UserController struct {
	DB         *ent.Client
	repository UserRepositoryInterface
}

func NewUserController(db *ent.Client) UserControllerInterface {
	return &UserController{
		DB:         db,
		repository: NewUserRespository(db),
	}
}

func (uc *UserController) Create(c *gin.Context) {
	user, err := uc.repository.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorCreatingUser})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("username")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorIDNull})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "user"})
}
