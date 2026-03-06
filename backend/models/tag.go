package models

// Tag 标签（适配现有数据库）
type Tag struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
