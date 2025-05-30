package models

type CreateClinicRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateClinicResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	DBName string `json:"db_name"`
}
