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
	GetByUsername(c *gin.Context)
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

func (ctrl *UserController) Create(c *gin.Context) {
	var postUser *User
	err := c.BindJSON(&postUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorCreatingUser})
		return
	}

	user, err := ctrl.repository.Create(postUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorCreatingUser})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (ctrl *UserController) GetByUsername(c *gin.Context) {
	username := c.Param("username")

	if len(username) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorIDNull})
		return
	}

	user, err := ctrl.repository.Get(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorGettingUserById})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": toPublicUser(user)})
}
