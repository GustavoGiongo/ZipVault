package app

import (
	download "ZipVault/internal/zipvault/domain/download"
	"ZipVault/internal/zipvault/dto/dto"
	"strings"
)

type DownloadService struct {
	repo download.DownloadRepository
}

func (s DownloadService) SearchFileRecords(d *dto.ListingRequest) ([]dto.FileRecord, error) {
	strings.ToLower(d.Name)
	var records, err = s.repo.FindFileRecordsByNameOrDate(d.Name, "", "")
	if err != nil {
		return records, err
	}
	return records, nil
}

func NewDownloadService(repo download.DownloadRepository) *DownloadService {
	return &DownloadService{repo: repo}
}

//func Sanitize(s string) string {
//	return strings.ReplaceAll(s, " ", "")
//}
