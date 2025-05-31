package repositoryInterface

import "ZipVault/internal/zipvault/dto/dto"

type Repository interface {
	FindFileRecordsByNameOrDate(name, startDate, endDate string) ([]dto.FileRecord, error)
}
