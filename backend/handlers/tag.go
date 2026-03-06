package handlers

import (
	"net/http"

	"comic-site/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TagHandler struct {
	DB *gorm.DB
}

func NewTagHandler(db *gorm.DB) *TagHandler {
	return &TagHandler{DB: db}
}

func (h *TagHandler) List(c *gin.Context) {
	var tags []models.Tag
	h.DB.Order("id ASC").Find(&tags)
	c.JSON(http.StatusOK, tags)
}

func (h *TagHandler) Categories(c *gin.Context) {
	var categories []models.Category
	h.DB.Order("id ASC").Find(&categories)
	c.JSON(http.StatusOK, categories)
}
