package domain

import (
	"test/dto"
	"test/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAstext() string {
	statusAstext := "active"
	if c.Status == "0" {
		statusAstext = "inactive"
	}
	return statusAstext
}
func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAstext(),
	}
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
