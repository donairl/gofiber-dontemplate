package models

import (
	"time"

	"github.com/donairl/gofiber-dontemplate/lib/database"
)

type Post struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	UserID    uint `gorm:"index"` // assuming you want to establish a relationship with users table
}

func (p Post) TableName() string {
	return "posts"
}

// Get all posts row
func GetAllPosts() ([]Post, error) {
	var posts []Post
	err := database.Connection.Find(&posts).Error // Find all posts
	if err != nil {
		return nil, err // return nil, err
	}
	return posts, nil // return posts, nil
}

// Get a post by ID row
func GetPostByID(id uint) (Post, error) {
	var post Post
	err := database.Connection.Where("id =?", id).First(&post).Error // Find a post with id
	if err != nil {
		return post, err // return post, err
	}
	return post, nil // return post, nil
}

func (p *Post) Create() error {
	err := database.Connection.Create(p).Error // create new post
	if err != nil {
		return err // return err
	}
	return nil // return nil if no error occurred
}

func (p *Post) Update() error {
	err := database.Connection.Save(p).Error // update a post with id = p.ID
	if err != nil {
		return err
	}
	return nil
}
