package database

import (
	"fmt"
	"log"

	"pertemuan6/structs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDatabase(dbUsername, dbPassword, dbPort, dbName, dbHost string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("ERROR: Failed to connect to Database -> %v\n", err)
	}
	db.AutoMigrate(structs.PersonGorm{})
	return db
}
