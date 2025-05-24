package api

import (
	"ZipVault/internal/api/download"
	"ZipVault/internal/api/health"
	"github.com/gin-gonic/gin"
	_ "net/http"

	_ "github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.GET("/health", health.HealthHandler)
	router.GET("/download", download.DownloadHandler)
	router.Run(":8080")
}
