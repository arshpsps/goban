package dataHandle

type DataHandler interface {
	Insert(item *Project)
	UpdateCard(card Card)
	GetCardsInProject(id int) []Card
	GetAllProjects() []Project
	GetBoardsInProject(id int) []Board
	GetCard(id int) Card
	GetBoard(id int) Board
	GetProject(id int) Project
}

// TODO: something similar to an interface but for data / struct. example: a Card struct without a gorm.Model field OR a monad of some sort
