package user

import (
	"context"
	"fmt"

	"github.com/bypepe77/api/ent"
	"github.com/bypepe77/api/ent/user"
)

type Params struct {
	key   string
	value string
}

type UserRepositoryInterface interface {
	Create(postUser *User) (*ent.User, error)
	Get(username string) (*ent.User, error)
	Follow(userToFollow *ent.User, userWhoIsFollowingID int) (*ent.User, error)
	Unfollow(userToFollow *ent.User, userWhoIsUnfollowingID int) (*ent.User, error)
}

type UserRepository struct {
	DB *ent.Client
}

func NewUserRespository(db *ent.Client) UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}

func (repository *UserRepository) Create(postUser *User) (*ent.User, error) {
	user, err := repository.DB.User.Create().
		SetName(postUser.Name).
		SetUsername(postUser.Username).
		SetAge(postUser.Age).
		SetPassword(postUser.Password).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	return user, nil
}

func (repository *UserRepository) Get(username string) (*ent.User, error) {
	user, err := repository.DB.User.
		Query().
		Where(user.Username(username)).
		Only(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return user, nil
}

func (repository *UserRepository) Follow(userToFollow *ent.User, userWhoIsFollowingID int) (*ent.User, error) {
	follow, err := userToFollow.Update().
		AddFollowerIDs(userWhoIsFollowingID).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed following user: %w", err)
	}

	return follow, nil
}

func (repository *UserRepository) Unfollow(userToFollow *ent.User, userWhoIsUnfollowingID int) (*ent.User, error) {
	follow, err := userToFollow.Update().
		RemoveFollowerIDs(userWhoIsUnfollowingID).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed unfollowing user: %w", err)
	}

	return follow, nil
}
