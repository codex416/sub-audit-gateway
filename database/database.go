package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var DB *gorm.DB



func InitDatabase() error {


	db, err := gorm.Open(
		sqlite.Open("audit.db"),
		&gorm.Config{},
	)


	if err != nil {

		return err

	}



	DB = db



	err = DB.AutoMigrate(
		&AuditLog{},
	)


	return err

}
