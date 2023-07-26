package models

type PostData struct {
	ID        uint64 `json:"id" db:"id"`
	CreatedAt int64  `json:"created_at" gorm:"autoUpdateTime:nano"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:nano"`
	Title     string `json:"title" db:"title"`
	Body      string `json:"body" db:"body"`
}
