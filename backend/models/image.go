package models

type ComicImage struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	ComicID   uint   `json:"comic_id" gorm:"index"`
	Sort      int    `json:"sort"`
	Filename  string `json:"filename" gorm:"size:200"`
	Extension string `json:"extension" gorm:"size:10"`
	URL       string `json:"url" gorm:"size:500"`
	LocalPath string `json:"local_path" gorm:"size:500"`
}
