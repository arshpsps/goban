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

type dbconn struct {
	db       *gorm.DB
	project  Project
	cards    []Card
	boards   []Board
	projects []Project
	board    Board
	card     Card
}

func Conndb() *gorm.DB {
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Println("Connection Successful!")
	return db
}

func (dbconn *dbconn) CreateTables() {
	var err error
	err = dbconn.db.AutoMigrate(Project{})
	if err != nil {
		log.Fatalf("Failed to migrate structs to tables: %s", err)
	}
	err = dbconn.db.AutoMigrate(Board{})
	if err != nil {
		log.Fatalf("Failed to migrate structs to tables: %s", err)
	}
	err = dbconn.db.AutoMigrate(Card{})
	if err != nil {
		log.Fatalf("Failed to migrate structs to tables: %s", err)
	}
}

func (dbconn *dbconn) GetProjects() {
	dbconn.db.Find(&dbconn.projects)
	for _, proj := range dbconn.projects {
		fmt.Print(proj.Name)
	}
}
