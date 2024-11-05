package dataHandle

import "gorm.io/gorm"

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

type JsonRoot struct {
	Projects []JsonProject
}

type JsonProject struct {
	Name   string
	Boards []JsonBoard
}

type JsonBoard struct {
	Name  string
	Cards []JsonCard
}

type JsonCard struct {
	Name        string
	Description string
	Status      int
}
