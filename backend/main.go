package main

import (
	"log"
	"os"
	"path/filepath"

	"comic-go/backend/config"
	"comic-go/backend/handlers"
	"comic-go/backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()
	middleware.SetJWTSecret(cfg.JWTSecret)

	// 确保目录存在
	os.MkdirAll(filepath.Dir(cfg.DBSource), 0755)
	os.MkdirAll(cfg.UploadDir, 0755)

	// 数据库
	db, err := gorm.Open(sqlite.Open(cfg.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	log.Printf("数据库连接成功: %s\n", cfg.DBSource)

	// 禁用自动迁移，使用现有数据库结构
	// db.AutoMigrate(
	// 	&models.Comic{}, &models.Tag{}, &models.Category{},
	// 	&models.ComicImage{}, &models.User{},
	// 	&models.Favorite{}, &models.Rating{},
	// )

	// Handlers
	comicH := handlers.NewComicHandler(db)
	tagH := handlers.NewTagHandler(db)
	userH := handlers.NewUserHandler(db)
	imageH := handlers.NewImageHandler(db, cfg)

	// Router
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	// 静态文件
	r.Static("/uploads", cfg.UploadDir)

	api := r.Group("/api")
	{
		// 公开接口
		api.GET("/comics", comicH.List)
		api.GET("/comics/:id", comicH.Detail)
		api.GET("/comics/:id/images", comicH.Images)
		api.GET("/tags", tagH.List)
		api.GET("/categories", tagH.Categories)
		api.GET("/search", comicH.Search)
		api.GET("/images/:comic_id/:filename", imageH.Serve)

		// 认证接口
		api.POST("/auth/register", userH.Register)
		api.POST("/auth/login", userH.Login)

		// 需要登录
		auth := api.Group("", middleware.AuthRequired())
		{
			auth.POST("/comics/:id/rate", comicH.Rate)
			auth.GET("/user/favorites", userH.Favorites)
			auth.POST("/user/favorites/:id", userH.AddFavorite)
			auth.DELETE("/user/favorites/:id", userH.RemoveFavorite)
			auth.POST("/admin/images/:comic_id", imageH.Upload)
		}
	}

	log.Printf("服务启动在 :%s\n", cfg.Port)
	r.Run(":" + cfg.Port)
}
