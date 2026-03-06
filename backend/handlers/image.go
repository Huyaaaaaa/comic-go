package handlers

import (
	"net/http"
	"path/filepath"
	"strconv"

	"comic-go/backend/config"
	"comic-go/backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ImageHandler struct {
	DB  *gorm.DB
	Cfg *config.Config
}

func NewImageHandler(db *gorm.DB, cfg *config.Config) *ImageHandler {
	return &ImageHandler{DB: db, Cfg: cfg}
}

func (h *ImageHandler) Serve(c *gin.Context) {
	comicID := c.Param("comic_id")
	filename := c.Param("filename")

	var img models.ComicImage
	if err := h.DB.Where("comic_id = ? AND filename = ?", comicID, filename).First(&img).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "图片不存在"})
		return
	}

	if img.LocalPath != "" {
		c.File(img.LocalPath)
		return
	}

	// 重定向到 CDN
	c.Redirect(http.StatusFound, img.URL)
}

func (h *ImageHandler) Upload(c *gin.Context) {
	comicID, _ := strconv.Atoi(c.Param("comic_id"))
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请上传文件"})
		return
	}

	sort, _ := strconv.Atoi(c.PostForm("sort"))
	ext := filepath.Ext(file.Filename)
	saveName := strconv.Itoa(comicID) + "_" + strconv.Itoa(sort) + ext
	savePath := filepath.Join(h.Cfg.UploadDir, saveName)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存失败"})
		return
	}

	img := models.ComicImage{
		ComicID:   comicID,
		Sort:      sort,
		Filename:  saveName,
		Extension: ext,
		LocalPath: savePath,
	}
	h.DB.Create(&img)
	c.JSON(http.StatusOK, img)
}
