package user

import (
	"net/http"

	"github.com/bypepe77/api/ent"
	"github.com/gin-gonic/gin"
)

const (
	ErrorCreatingUser    = "Error creating user"
	ErrorUsernameNull    = "Username can't be blank"
	ErrorGettingUserById = "User not found"
	ErrorWhoWillFollow   = "Username of who you going to follow/unfollow can't be blank"
	ErrorFollowing       = "Error following this user"
	ErrorUnfollowing     = "Error Unfollowing this user"
)

type UserControllerInterface interface {
	Create(c *gin.Context)
	GetByUsername(c *gin.Context)
	Follow(c *gin.Context)
	Unfollow(c *gin.Context)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorUsernameNull})
		return
	}

	user, err := ctrl.repository.Get(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorGettingUserById})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": toPublicUser(user)})
}

func (ctrl *UserController) Follow(c *gin.Context) {
	userToFollow := c.Param("whoIWillFollow")

	if len(userToFollow) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorWhoWillFollow})
		return
	}

	user, err := ctrl.repository.Get(userToFollow)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorFollowing})
		return
	}

	follow, err := ctrl.repository.Follow(user, 1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorFollowing})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": follow})
}

func (ctrl *UserController) Unfollow(c *gin.Context) {
	userToUnfollow := c.Param("whoIwillUnfollow")

	if len(userToUnfollow) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorWhoWillFollow})
		return
	}

	user, err := ctrl.repository.Get(userToUnfollow)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorFollowing})
		return
	}

	follow, err := ctrl.repository.Unfollow(user, 1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorUnfollowing})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": follow})
}

func returnJson(c *gin.Context, status int, name string, data interface{}) {
	c.JSON(status, gin.H{name: data})
	return
}
