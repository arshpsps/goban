package kanban

import (
	"fmt"
	"goban/internals/dataHandle"
	"os"
)

var (
    fileLocation string = "goban/data/data.json"
)

func fileOpen() *os.File{
    file, err := os.Open(fileLocation)
    if err != nil {
        fmt.Errorf("Error: %v", err)
        os.Exit(1)
    }
    return file
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
	dataHandle.JsonRead(data)
}
