package service

import (
	"ZipVault/internal/dto"
	"fmt"
	"strings"
)

func DownloadService(d *dto.DownloadRequest) (bool, error) {
	d.Name = Sanitize(d.Name)
	fmt.Println(d.Name)
	return true, nil

}

func Sanitize(s string) string {
	return strings.ReplaceAll(s, " ", "")

}
