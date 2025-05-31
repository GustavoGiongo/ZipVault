package mysql

import (
	"ZipVault/internal/zipvault/dto/dto"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	_ "testing"
)

func TestFindByNameAndDate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMySQLRepository(db)

	query := "SELECT name,date,folder FROM files WHERE LOWER(name) LIKE ? AND date BETWEEN ? AND ?"

	mock.ExpectPrepare(regexp.QuoteMeta(query)).
		ExpectQuery().
		WithArgs("%report%", "2024-01-01", "2024-01-31").
		WillReturnRows(sqlmock.NewRows([]string{"name", "date", "folder"}).
			AddRow("report_1.zip", "2024-01-15", "daily"))

	result, err := repo.FindFileRecordsByNameOrDate("report", "2024-01-01", "2024-01-31")

	assert.NoError(t, err)
	assert.Equal(t, []dto.FileRecord{
		{Name: "report_1.zip", Date: "2024-01-15", Folder: "daily"},
	}, result)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindByStartAndEndDate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMySQLRepository(db)

	query := "SELECT name,date,folder FROM files WHERE date BETWEEN ? AND ?"

	mock.ExpectPrepare(regexp.QuoteMeta(query)).
		ExpectQuery().
		WithArgs("2024-01-01", "2024-01-31").
		WillReturnRows(sqlmock.NewRows([]string{"name", "date", "folder"}).
			AddRow("report_1.zip", "2024-01-15", "daily"))

	result, err := repo.FindFileRecordsByNameOrDate("", "2024-01-01", "2024-01-31")

	assert.NoError(t, err)
	assert.Equal(t, []dto.FileRecord{
		{Name: "report_1.zip", Date: "2024-01-15", Folder: "daily"},
	}, result)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindByStartDate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMySQLRepository(db)

	query := "SELECT name,date,folder FROM files WHERE date = ?"

	mock.ExpectPrepare(regexp.QuoteMeta(query)).
		ExpectQuery().
		WithArgs("2024-01-01").
		WillReturnRows(sqlmock.NewRows([]string{"name", "date", "folder"}).
			AddRow("report_1.zip", "2024-01-15", "daily"))

	result, err := repo.FindFileRecordsByNameOrDate("", "2024-01-01", "")

	assert.NoError(t, err)
	assert.Equal(t, []dto.FileRecord{
		{Name: "report_1.zip", Date: "2024-01-15", Folder: "daily"},
	}, result)

	assert.NoError(t, mock.ExpectationsWereMet())
}
