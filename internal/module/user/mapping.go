package user

import "github.com/bypepe77/api/ent"

type PublicUser struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func toPublicUser(user *ent.User) PublicUser {
	return PublicUser{
		Name: user.Name,
		Age:  user.Age,
	}
}
