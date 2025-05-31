package app

import (
	download "ZipVault/internal/zipvault/domain/interface"
	"ZipVault/internal/zipvault/dto/dto"
)

type ListingService struct {
	repo download.Repository
}

func (s ListingService) SearchFileRecords(d *dto.ListingRequest) ([]dto.FileRecord, error) {
	var records, err = s.repo.FindFileRecordsByNameOrDate(d.Name, "", "")
	if err != nil {
		return records, err
	}
	return records, nil
}

func NewListingService(repo download.Repository) *ListingService {
	return &ListingService{repo: repo}
}
