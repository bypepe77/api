package user

import (
	"net/http"

	"github.com/bypepe77/api/internal/common/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	ErrorCreatingUser    = "Error creating user"
	ErrorIDNull          = "ID can't be blank"
	ErrorGettingUserById = "User not found"
)

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
}

type UserController struct {
	DB         gorm.DB
	repository UserRepositoryInterface
}

func NewUserController(db gorm.DB) UserControllerInterface {
	return &UserController{
		DB:         db,
		repository: NewUserRespository(db),
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var requestBody *models.User

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorCreatingUser})
	}

	user, err := uc.repository.Create(requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorCreatingUser})
	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}

func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorIDNull})
		return
	}
	params := &Params{
		key:   "id",
		value: id,
	}
	user, err := uc.repository.GetBy(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorGettingUserById})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
