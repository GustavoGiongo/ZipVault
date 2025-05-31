package dto

type ListingRequest struct {
	Name      string `form:"name"`
	StartDate string `form:"startDate"`
	EndDate   string `form:"edDate"`
}

type FileRecord struct {
	Name   string
	Date   string
	Folder string
	PATH   string
}
