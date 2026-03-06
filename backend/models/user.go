package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"size:100;uniqueIndex"`
	Email     string    `json:"email" gorm:"size:200;uniqueIndex"`
	Password  string    `json:"-" gorm:"size:200"`
	IsAdmin   bool      `json:"is_admin" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
}

type Favorite struct {
	UserID  uint `json:"user_id" gorm:"primaryKey"`
	ComicID uint `json:"comic_id" gorm:"primaryKey"`
}

type Rating struct {
	UserID  uint    `json:"user_id" gorm:"primaryKey"`
	ComicID uint    `json:"comic_id" gorm:"primaryKey"`
	Score   float64 `json:"score"`
}
