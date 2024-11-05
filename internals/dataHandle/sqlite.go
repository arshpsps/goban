package dataHandle

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBConn struct {
	db *gorm.DB
}

func NewSqliteConndb() DBConn {
	fmt.Println("hey")
	var err error
	db := DBConn{}
	db.db, err = gorm.Open(sqlite.Open("data/gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Println("Connection Successful!")
	return db
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

func (dbconn *DBConn) Insert(item *Project) {
	dbconn.db.Create(item)
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

func UpdateView(inp []string) {
	fmt.Println(inp)
	fmt.Println(len(inp))
}

func (DBConn *DBConn) UpdateCard(card Card) {
	DBConn.db.Save(&card)
}
