package user

import (
	"context"

	"github.com/bypepe77/api/ent"
)

type PublicUser struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Age            int    `json:"age,omitempty"`
	FollowingCount int    `json:"following_count"`
	FollowerCount  int    `json:"follower_count"`
}

func ToPublicUser(user *ent.User) PublicUser {
	return PublicUser{
		Name:           user.Name,
		Age:            user.Age,
		FollowingCount: user.QueryFollowing().CountX(context.Background()),
		FollowerCount:  user.QueryFollowers().CountX(context.Background()),
	}
}
