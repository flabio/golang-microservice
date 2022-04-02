package domain

import (
	"strconv"
	"test/errs"
	"test/logger"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	clienta *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {

	sqlInsert := "INSERT INTO accounts(customer_id,opening_date,account_type,amount,status) VALUES(?,?,?,?,?)"
	result, err := d.clienta.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewNotFoundError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last  insert id for  new account: " + err.Error())
		return nil, errs.NewNotFoundError("Unexpected error from database")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}
func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{clienta: dbClient}
}
