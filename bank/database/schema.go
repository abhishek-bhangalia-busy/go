package database

import (
	"bank/models"

	pg "github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.Bank)(nil),
		(*models.Branch)(nil),
		(*models.Customer)(nil),
		(*models.Account)(nil),
		(*models.AccountToCustomer)(nil),
		(*models.Transaction)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
