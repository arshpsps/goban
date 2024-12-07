package dataHandle

type DataHandler interface {
	Insert(item *Project)
	InsertBoard(item *Board)
	InsertCard(item *Card)
	UpdateCard(card Card)
	GetCardsInBoard(id uint) []Card
	GetAllProjects() []Project
	GetBoardsInProject(id uint) []Board
	GetCard(id int) Card
	GetBoard(id int) Board
	GetProject(id int) Project
}

// TODO: something similar to an interface but for data / struct. example: a Card struct without a gorm.Model field OR a monad of some sort
