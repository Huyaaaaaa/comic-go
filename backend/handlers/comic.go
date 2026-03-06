package handlers

import (
	"net/http"
	"strconv"

	"comic-go/backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ComicHandler struct {
	DB *gorm.DB
}

func NewComicHandler(db *gorm.DB) *ComicHandler {
	return &ComicHandler{DB: db}
}

func (h *ComicHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "40"))
	sort := c.DefaultQuery("sort", "id")
	tagID := c.Query("tag_id")
	categoryID := c.Query("category_id")
	search := c.Query("search")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 40
	}

	query := h.DB.Model(&models.Comic{})

	if tagID != "" {
		query = query.Joins("JOIN comic_tags ON comic_tags.comic_id = comics.id").
			Where("comic_tags.tag_id = ?", tagID)
	}
	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}
	if search != "" {
		query = query.Where("title LIKE ? OR subtitle LIKE ? OR author LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	// 排序
	switch sort {
	case "rating":
		query = query.Order("rating DESC")
	case "favorites":
		query = query.Order("favorites DESC")
	case "newest":
		query = query.Order("created_at DESC")
	case "updated":
		query = query.Order("updated_at DESC")
	default:
		query = query.Order("id DESC")
	}

	var total int64
	query.Count(&total)

	var comics []models.Comic
	query.Preload("Tags").Offset((page - 1) * pageSize).Limit(pageSize).Find(&comics)

	c.JSON(http.StatusOK, gin.H{
		"comics":     comics,
		"total":      total,
		"page":       page,
		"page_size":  pageSize,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}

func (h *ComicHandler) Detail(c *gin.Context) {
	id := c.Param("id")
	var comic models.Comic
	if err := h.DB.Preload("Tags").Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Order("sort ASC")
	}).First(&comic, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "漫画不存在"})
		return
	}
	c.JSON(http.StatusOK, comic)
}

func (h *ComicHandler) Images(c *gin.Context) {
	id := c.Param("id")
	var images []models.ComicImage
	h.DB.Where("comic_id = ?", id).Order("sort ASC").Find(&images)
	c.JSON(http.StatusOK, images)
}

func (h *ComicHandler) Rate(c *gin.Context) {
	comicID, _ := strconv.Atoi(c.Param("id"))
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	var req struct {
		Score float64 `json:"score" binding:"required,min=1,max=10"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "评分需在 1-10 之间"})
		return
	}

	rating := models.Rating{UserID: userID.(uint), ComicID: comicID, Score: req.Score}
	h.DB.Where("user_id = ? AND comic_id = ?", userID, comicID).
		Assign(models.Rating{Score: req.Score}).FirstOrCreate(&rating)

	// 更新漫画平均评分
	var avg struct{ Avg float64 }
	var count int64
	h.DB.Model(&models.Rating{}).Where("comic_id = ?", comicID).Count(&count)
	h.DB.Model(&models.Rating{}).Where("comic_id = ?", comicID).Select("COALESCE(AVG(score),0) as avg").Scan(&avg)
	h.DB.Model(&models.Comic{}).Where("id = ?", comicID).Updates(map[string]interface{}{
		"rating":       avg.Avg,
		"rating_count": count,
	})

	c.JSON(http.StatusOK, gin.H{"rating": avg.Avg, "count": count})
}

func (h *ComicHandler) Search(c *gin.Context) {
	c.Request.URL.RawQuery += "&search=" + c.Query("q")
	h.List(c)
}
