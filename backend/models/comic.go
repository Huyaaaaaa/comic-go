package models

import (
	"time"
)

// Comic 漫画主表（适配现有数据库）
type Comic struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	Title           string    `json:"title"`
	Subtitle        string    `json:"subtitle"`
	Author          string    `json:"author"`
	AuthorID        int       `json:"author_id"`
	CoverURL        string    `json:"cover_url"`
	CoverBase64     string    `json:"cover_base64"`
	Rating          float64   `json:"rating"`
	RatingCount     int       `json:"rating_count"`
	Favorites       int       `json:"favorites"`
	CategoryID      int       `json:"category_id"`
	CategoryName    string    `json:"category_name"`
	CreatedAt       string    `json:"created_at"`
	UpdatedAt       string    `json:"updated_at"`
	HasCoverCached  int       `json:"has_cover_cached"`
	CoverCachedAt   time.Time `json:"cover_cached_at"`
	CoverSize       int       `json:"cover_size"`

	// 关联
	Tags    []Tag         `json:"tags,omitempty" gorm:"many2many:comic_tags;"`
	Authors []ComicAuthor `json:"authors,omitempty" gorm:"foreignKey:ComicID"`
	Images  []ComicImage  `json:"images,omitempty" gorm:"foreignKey:ComicID"`
}

// TableName 指定表名
func (Comic) TableName() string {
	return "comics"
}

// ComicAuthor 漫画作者关联表
type ComicAuthor struct {
	ComicID    int    `json:"comic_id" gorm:"primaryKey"`
	AuthorID   int    `json:"author_id" gorm:"primaryKey"`
	AuthorName string `json:"author_name" gorm:"primaryKey"`
	Position   int    `json:"position"`
}

// Category 分类
type Category struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
