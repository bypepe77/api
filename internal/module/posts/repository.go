package posts

import (
	"context"
	"fmt"

	"github.com/bypepe77/api/ent"
	"github.com/bypepe77/api/ent/post"
)

type PostsRepositoryInterface interface {
	Create(tags []string, post *Post) (*ent.Post, error)
	GetById(PostId int) (*ent.Post, error)
}

type PostRepository struct {
	DB *ent.Client
}

func NewPostRepositoy(db *ent.Client) PostsRepositoryInterface {
	return &PostRepository{
		DB: db,
	}
}

func (r *PostRepository) Create(tags []string, post *Post) (*ent.Post, error) {
	postCreated, err := r.DB.Post.
		Create().
		SetAuthorID(1).
		SetContent(post.Content).
		SetHashtags(tags).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed creating Post: %w", err)
	}

	return postCreated, nil
}

func (r *PostRepository) GetById(PostId int) (*ent.Post, error) {
	post, err := r.DB.Post.
		Query().
		Where(post.ID(PostId)).
		Only(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed creating Post: %w", err)
	}

	return post, nil
}
