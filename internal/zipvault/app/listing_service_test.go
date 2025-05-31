package app

import (
	"ZipVault/internal/zipvault/dto/dto"
	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) FindFileRecordsByNameOrDate(name, startDate, endDate string) ([]dto.FileRecord, error) {
	args := m.Called(name, startDate, endDate)
	return args.Get(0).([]dto.FileRecord), args.Error(1)
}

func TestListingService_SearchFileRecords(t *testing.T) {
	mockRepo := new(MockRepository)

	expected := []dto.FileRecord{
		{Name: "report_1.zip", Date: "2024-01-01", Folder: "daily"},
	}

	mockRepo.
		On("FindFileRecordsByNameOrDate", "report", "", "").
		Return(expected, nil)

	service := ListingService{repo: mockRepo}

	req := &dto.ListingRequest{Name: "report"}
	result, err := service.SearchFileRecords(req)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}
