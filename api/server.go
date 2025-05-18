package api

import (
	"github.com/gin-gonic/gin"
	_ "net/http"

	_ "github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.GET("/health", HealthHandler)
	router.Run(":8080")
}
