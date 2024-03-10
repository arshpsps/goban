package kanban

import (
	"fmt"
	"goban/internals/dataHandle"
)

var data string = `{
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
	dataHandle.JsonRead(data)
}
