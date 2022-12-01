package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Slug           string `gorm:"unique_index;not null"`
	Body           string
	Author         User
	AuthorID       uint
	CommentCount   int64
	FavoritesCount int64
}

type Comment struct {
	gorm.Model
	Post   Post
	PostID uint
	User   User
	UserID uint
	Body   string
}
