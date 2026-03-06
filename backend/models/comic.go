package models

import (
	"time"

	"gorm.io/gorm"
)

type Comic struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Title        string         `json:"title" gorm:"size:500"`
	Subtitle     string         `json:"subtitle" gorm:"size:500"`
	Author       string         `json:"author" gorm:"size:200"`
	CoverURL     string         `json:"cover_url" gorm:"size:500"`
	Rating       float64        `json:"rating" gorm:"default:0"`
	RatingCount  int            `json:"rating_count" gorm:"default:0"`
	RatingTotal  float64        `json:"-" gorm:"default:0"`
	Favorites    int            `json:"favorites" gorm:"default:0"`
	CategoryID   uint           `json:"category_id"`
	CategoryName string         `json:"category_name" gorm:"size:100"`
	Tags         []Tag          `json:"tags" gorm:"many2many:comic_tags;"`
	Images       []ComicImage   `json:"images,omitempty" gorm:"foreignKey:ComicID"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"size:100;uniqueIndex"`
}
