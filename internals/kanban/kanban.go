package kanban

import (
	"encoding/json"
	"fmt"
	// "encoding/json"
)

func jsonRead(data string) {
	var pj JData
	if !json.Valid([]byte(data)) {
		fmt.Println("invalid JSON:", data)
		return
	}
	json.Unmarshal([]byte(data), &pj)
	fmt.Println(pj.Projects[0].Boards[0].Cards[0].Tags[0])
}

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

type JData struct {
	Projects []Project
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
	fmt.Println(arg)
	data := `{
    "projects": [
        {
            "name": "project1",
            "created_on": 20240209,
            "boards": [

                {
                    "name": "board1",
                    "cards": [
                        {
                            "title": "card1",

                            "description": "this is the first card",
                            "status": 1,
                            "tags": ["tag1"]
                        }
                    ]
                }

            ]

        }
    ]
}
`
	jsonRead(data)
	// c := Card{"c1", "c1d", []string{"c1t1", "c1t2"}, 1}
	// fmt.Printf("%+v\n", c)
}
