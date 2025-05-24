package dto

type DownloadRequest struct {
	Name string `form:"name" binding:"required"`
	Date string `form:"date" binding:"required"`
}
