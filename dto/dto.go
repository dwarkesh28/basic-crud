package dto

import (
	"go-crud/models"
	"time"
)

type GetPostData struct {
	ID        uint64    `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoUpdateTime:nano"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:nano"`
	Title     string    `json:"title" db:"title"`
	Body      string    `json:"body" db:"body"`
}

func ToDto(post models.PostData) *GetPostData {
	return &GetPostData{
		ID:        post.ID,
		CreatedAt: time.Unix(0, post.CreatedAt),
		UpdatedAt: time.Unix(0, post.UpdatedAt),
		Title:     post.Title,
		Body:      post.Body,
	}
}

type GetUserData struct {
	ID        uint64    `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoUpdateTime:nano"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:nano"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
}


func ToDtoUser(user models.User) *GetUserData {
	return &GetUserData{
		ID:        user.ID,
		CreatedAt: time.Unix(0, user.CreatedAt),
		UpdatedAt: time.Unix(0, user.UpdatedAt),
		Email:     user.Email,
		Password:  user.Password,
	}
}