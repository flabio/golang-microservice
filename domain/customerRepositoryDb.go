package domain

import (
	"database/sql"
	"test/errs"
	"test/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	//var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)
	findAllSql := "select * from customers"
	err = d.client.Select(&customers, findAllSql)
	//rows, err = d.client.Query(findAllSql)
	/*else {
		findAllSql := "select * from customers where status=?"
		err = d.client.Select(&customers, findAllSql, status)

		//rows, err = d.client.Query(findAllSql,status)
	}*/

	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, err
	}
	/*
		err = sqlx.StructScan(rows, &customers)
		if err != nil {
			logger.Error("Error while querying customer table" + err.Error())
			return nil, err
		}*/
	/*for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateofBirth, &c.Status)
		if err != nil {
			logger.Error("Error while querying customer table" + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}*/
	return customers, nil
}
func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select * from customers where customer_id=?"

	//row := d.client.QueryRow(customerSql, 1)
	var c Customer
	err := d.client.Get(&c, customerSql, id)
	//err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while querying customer table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

	}
	return &c, nil
}
func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {

	return CustomerRepositoryDb{client: dbClient}
}
