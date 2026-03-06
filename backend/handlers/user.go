package handlers

import (
	"net/http"

	"comic-site/middleware"
	"comic-site/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required,min=3,max=50"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
		return
	}

	user := models.User{Username: req.Username, Email: req.Email, Password: string(hash)}
	if err := h.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "用户名或邮箱已存在"})
		return
	}

	token, _ := middleware.GenerateToken(user.ID, user.IsAdmin)
	c.JSON(http.StatusCreated, gin.H{"token": token, "user": user})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := h.DB.Where("username = ? OR email = ?", req.Username, req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return
	}

	token, _ := middleware.GenerateToken(user.ID, user.IsAdmin)
	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}

func (h *UserHandler) Favorites(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var comics []models.Comic
	h.DB.Joins("JOIN favorites ON favorites.comic_id = comics.id").
		Where("favorites.user_id = ?", userID).
		Preload("Tags").Find(&comics)
	c.JSON(http.StatusOK, comics)
}

func (h *UserHandler) AddFavorite(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	comicID := c.Param("id")
	fav := models.Favorite{UserID: userID}
	h.DB.Raw("SELECT ? as comic_id", comicID).Scan(&fav.ComicID)
	if err := h.DB.Create(&fav).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "已收藏"})
		return
	}
	// 更新漫画收藏数
	h.DB.Model(&models.Comic{}).Where("id = ?", comicID).
		UpdateColumn("favorites", gorm.Expr("favorites + 1"))
	c.JSON(http.StatusOK, gin.H{"message": "收藏成功"})
}

func (h *UserHandler) RemoveFavorite(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	comicID := c.Param("id")
	result := h.DB.Where("user_id = ? AND comic_id = ?", userID, comicID).Delete(&models.Favorite{})
	if result.RowsAffected > 0 {
		h.DB.Model(&models.Comic{}).Where("id = ?", comicID).
			UpdateColumn("favorites", gorm.Expr("CASE WHEN favorites > 0 THEN favorites - 1 ELSE 0 END"))
	}
	c.JSON(http.StatusOK, gin.H{"message": "已取消收藏"})
}
