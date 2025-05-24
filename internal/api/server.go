package api

import (
	"ZipVault/api/download"
	"ZipVault/api/health"
	"github.com/gin-gonic/gin"
	_ "net/http"

	_ "github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.GET("/health", health.HealthHandler)
	router.POST("/download", download.DownloadHandler)
	router.Run(":8080")
}
