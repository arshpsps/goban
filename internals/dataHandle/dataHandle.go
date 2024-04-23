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
	db *gorm.DB
}

func Conndb() dbconn {
	var err error
	db := dbconn{}
	db.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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

func (dbconn *dbconn) GetAllProjects() []Project {
	var projects []Project
	dbconn.db.Find(&projects)
	for _, proj := range projects {
		fmt.Println(proj.Name)
	}
	return projects
}

func (dbconn *dbconn) GetProject(id int) Project {
	var project Project
	dbconn.db.First(&project, id)
	return project
}

func (dbconn *dbconn) GetBoard(id int) Board {
	var board Board
	dbconn.db.First(&board, id)
	return board
}

func (dbconn *dbconn) GetCard(id int) Card {
	var card Card
	dbconn.db.First(&card, id)
	return card
}
