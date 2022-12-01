package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username       string `gorm:"unique_index;not null"`
	Email          string `gorm:"unique_index;not null"`
	Password       string `gorm:"not null"`
	Bio            string
	Image          string
	FollowersCount int64
	FollowingCount int64
}

type Follow struct {
	Follower    User
	FollowerID  uint `gorm:"primary_key"`
	Following   User
	FollowingID uint `gorm:"primary_key"`
}
