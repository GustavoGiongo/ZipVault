package http

import (
	"ZipVault/internal/zipvault/app"
	"ZipVault/internal/zipvault/dto/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DownloadHandler struct {
	Service *app.DownloadService
}

func NewDownloadHandler(svc *app.DownloadService) *DownloadHandler {
	return &DownloadHandler{Service: svc}
}

func (dh *DownloadHandler) HandleDownload(c *gin.Context) {
	var req dto.DownloadRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameters"})
		return
	}
	location, err := dh.Service.SearchFileLocation(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		panic(err)
	}
	if location != "" {
		c.JSON(http.StatusOK, gin.H{
			"File location:": location,
		})
	}
	c.JSON(http.StatusNoContent, gin.H{"message:": "File was not Found"})
}
