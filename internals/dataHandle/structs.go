package dataHandle

type GlobalCard[t Card| JsonCard] struct {
    modelToUse t
}

type Card struct {
	Title       string
	Description string
	ID          uint
	BoardID     uint `gorm:"foreignKey"`
	Status      int
}

type Board struct {
	Name      string
	ID        uint
	ProjectID uint `gorm:"foreignKey"`
}

type Project struct {
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
