package connection

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseInfo struct {
	Host string
	DBName string
	User string
	Password string
}

func DbInit() (*gorm.DB, error) {
	d := dbInfo()

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		d.User, d.Password, d.Host, d.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("DB connection error: %v", err.Error())
	}
	
	return db, err
}


func dbInfo() DatabaseInfo {
	d := DatabaseInfo{
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	}

	return d
}