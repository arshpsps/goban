package kanban

import (
	"fmt"
)

type Card struct{
    Title string
    Description string
    Tags []string
    Status int
}

type Board struct{
    Name string
    Cards []Card
}

type Project struct{
    Name string
    Boards []Board
    Created_on int
}

func Start(args []string) {
	switch args[0] {

	case "create":
		if len(args) != 2 {
			fmt.Println("insufficient args")
		} else {
			create(args[1])
		}

	default:
		fmt.Println("invalid arguments, run 'help' to get a list of valid arguments.")
	}
}

func create(arg string) {
	fmt.Println("create called")
}
