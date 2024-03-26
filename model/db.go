package model

import (
	"fmt"
	"log"

	"os"

	// "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	Db *gorm.DB
}

var Mysql MysqlDB

func OpenDB() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME")) // username:password@(host:port)/name_database

	sqlConnection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Mysql = MysqlDB{
		Db: sqlConnection,
	}
	err = AutoMigrate(sqlConnection)
	if err != nil {
		log.Fatal(err)
	}
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&DaftarBuku{},
	)
	if err != nil {
		return err
	}
	return nil
}
