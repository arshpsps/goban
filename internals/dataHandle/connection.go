package dataHandle

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dsn string = "root:@tcp(127.0.0.1:3306)/goban?charset=utf8mb4&parseTime=True&loc=Local"
	DB  DBConn
)

type DBConn struct {
	db *gorm.DB
}

func NewConndb() DBConn {
	var err error
	db := DBConn{}
	db.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Println("Connection Successful!")
	return db
}
