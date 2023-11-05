package models

type PostData struct {
	ID        uint64 `json:"id" gorm:"primaryKey"`
	UserID    string `json:"user_id"`
	CreatedAt int64  `json:"created_at" gorm:"autoUpdateTime:nano"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:nano"`
	Title     string `json:"title" db:"title"`
	Body      string `json:"body" db:"body"`
	User      User   `gorm:"foreignKey:UserID;references:UserID"`
}

type User struct {
	ID        uint64 `json:"id" gorm:"primaryKey"`
	UserID    string `json:"uuid"`
	CreatedAt int64  `json:"created_at" gorm:"autoUpdateTime:nano"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:nano"`
	Email     string `gorm:"unique" json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string
}
