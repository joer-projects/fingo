package setup

import (
	schema "github.com/joer-projects/fingo/schema/accounting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() error {
	dsn := "host=localhost user=admin password=admin dbname=accounting port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&schema.Transaction{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&schema.TransactionLine{})
	if err != nil {
		return err
	}

	return nil
}
