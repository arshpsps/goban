package dataHandle

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

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

func (dbconn *DBConn) CreateTables() {
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

func (dbconn *DBConn) GetAllProjects() []Project {
	var projects []Project
	dbconn.db.Find(&projects)
	return projects
}

func (dbconn *DBConn) GetProject(id int) Project {
	var project Project
	dbconn.db.First(&project, id)
	return project
}

func (dbconn *DBConn) GetBoard(id int) Board {
	var board Board
	dbconn.db.First(&board, id)
	return board
}

func (dbconn *DBConn) GetCard(id int) Card {
	var card Card
	dbconn.db.First(&card, id)
	return card
}

func (dbconn *DBConn) GetBoardsInProject(id int) []Board {
	var boards []Board
	dbconn.db.Model(&boards).Where("project_id = ?", id).Find(&boards)
	return boards
}

func (dbconn *DBConn) GetCardsInProject(id int) []Card {
	var cards []Card
	dbconn.db.Model(&cards).Where("board_id = ?", id).Find(&cards)
	return cards
}

func (card Card) UpdateItemInDB(inp []string) {
	fmt.Println(inp)
}
