package dataHandle

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn string = "root:@tcp(127.0.0.1:3306)/goban?charset=utf8mb4&parseTime=True&loc=Local"

type Card struct {
	gorm.Model
	Title       string
	Description string
	ID          uint
	BoardID     uint `gorm:"foreignKey"`
	Status      int
}

type Board struct {
	gorm.Model
	Name      string
	ID        uint
	ProjectID uint `gorm:"foreignKey"`
}

type Project struct {
	gorm.Model
	Name string
	ID   uint
}

func Conndb() error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Println("uwu")
	db.AutoMigrate(Project{})
	db.AutoMigrate(Board{})
	db.AutoMigrate(Card{})
	return nil
}
