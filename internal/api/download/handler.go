package download

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DownloadRequest struct {
	Name   string `json:"name"`
	Date   string `json:"date"`
	Folder string `json:"folder"`
}

func DownloadHandler(c *gin.Context) {
	var req DownloadRequest
	err := c.ShouldBindJSON(&req)
	fmt.Println(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
