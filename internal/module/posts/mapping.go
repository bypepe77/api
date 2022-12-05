package posts

import (
	"context"
	"errors"
	"time"

	"github.com/bypepe77/api/ent"
	"github.com/bypepe77/api/internal/module/user"
	"github.com/gernest/mention"
)

type PublicPost struct {
	ID        int             `json:"id,omitempty"`
	CreatedAt time.Time       `json:"created_at,omitempty"`
	Author    user.PublicUser `json:"author,omitempty"`
	Content   string          `json:"content,omitempty"`
}

func parseTags(tags []mention.Tag) []string {
	var postTags []string

	for i := 0; i < len(tags); i++ {
		actualTag := tags[i]

		postTags = append(postTags, actualTag.Tag)
	}

	return postTags
}

func toPublicPost(post *ent.Post) (*PublicPost, error) {
	queryUser, err := post.QueryAuthor().First(context.Background())
	if err != nil {
		return nil, errors.New("Cannot fetch user")
	}
	return &PublicPost{
		ID:      post.ID,
		Content: post.Content,
		Author:  user.ToPublicUser(queryUser),
	}, nil
}
