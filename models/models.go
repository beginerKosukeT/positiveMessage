package models

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	UserName     string
	EmailAddress string
	Password     string
	IconNumber   int
}

// Post model
type Post struct {
	gorm.Model
	PostTitle     string
	Body          string
	AuthorsUserId int
	NumberOfLikes int
}

// Post join model(投稿と投稿者情報をまとめて表示するための構造体)
type PostOverview struct {
	Post
	User
}

// Like model
type Like struct {
	gorm.Model
	PostId int
	UserId int
}

// Save model
type Save struct {
	gorm.Model
	PostId int
	UserId int
}
