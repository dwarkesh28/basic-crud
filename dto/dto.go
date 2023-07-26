package dto

import (
	"go-crud/models"
	"time"
)

type GetData struct {
	ID        uint64    `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoUpdateTime:nano"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:nano"`
	Title     string    `json:"title" db:"title"`
	Body      string    `json:"body" db:"body"`
}

func ToDto(post models.PostData) *GetData {
	return &GetData{
		ID:        post.ID,
		CreatedAt: time.Unix(0, post.CreatedAt),
		UpdatedAt: time.Unix(0, post.UpdatedAt),
		Title:     post.Title,
		Body:      post.Body,
	}
}
