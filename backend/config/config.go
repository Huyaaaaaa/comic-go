package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	Port       string
	DBDriver   string
	DBSource   string
	JWTSecret  string
	UploadDir  string
	ImageCDNs  []string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		dbDriver = "sqlite"
	}
	dbSource := os.Getenv("DB_SOURCE")
	if dbSource == "" {
		// 使用绝对路径指向现有数据库
		dbSource = "/Users/huyaaaaaa/project/spider/data/comics.db"
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "comic-site-secret-key-change-in-production"
	}
	uploadDir := os.Getenv("UPLOAD_DIR")
	if uploadDir == "" {
		uploadDir = filepath.Join("..", "downloads")
	}

	return &Config{
		Port:      port,
		DBDriver:  dbDriver,
		DBSource:  dbSource,
		JWTSecret: jwtSecret,
		UploadDir: uploadDir,
		ImageCDNs: []string{
			"https://img.cimg-lux.top",
		},
	}
}
