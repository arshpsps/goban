package dataHandle

type Card struct {
	Title       string
	Description string
	Tags        []string
	Status      int
}

type Board struct {
	Name  string
	Cards []Card
}

type Project struct {
	Name       string
	Boards     []Board
	Created_on int
}
