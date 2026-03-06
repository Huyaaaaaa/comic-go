package models

// ComicImage 漫画图片（适配现有数据库）
type ComicImage struct {
	ComicID   int    `json:"comic_id" gorm:"primaryKey"`
	Sort      int    `json:"sort" gorm:"primaryKey"`
	Filename  string `json:"filename"`
	Extension string `json:"extension"`
	URL       string `json:"url"`
	LocalPath string `json:"local_path"`
}
