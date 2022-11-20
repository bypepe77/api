package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Slug        string `gorm:"unique_index;not null"`
	Title       string `gorm:"not null"`
	Description string
	Body        string
	Author      User
	AuthorID    uint
	Comments    []Comment `gorm:"foreignkey:PostID"`
	Favorites   []User    `gorm:"many2many:favorites;"`
}

type Comment struct {
	gorm.Model
	Post   Post
	PostID uint
	User   User
	UserID uint
	Body   string
}
