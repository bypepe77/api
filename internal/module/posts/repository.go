package posts

import "gorm.io/gorm"

type PostsRepositoryInterface interface {
	GetPost() string
}

type PostRepository struct {
	DB gorm.DB
}

func NewPostRepositoy(db gorm.DB) PostsRepositoryInterface {
	return &PostRepository{
		DB: db,
	}
}

func (r *PostRepository) GetPost() string {
	return "Post"
}
