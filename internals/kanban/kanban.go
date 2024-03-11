package kanban

import (
	"fmt"
	"goban/internals/dataHandle"
	"os"
)

var fileLocation string = "../../data/data.json"

func fileOpen() *os.File {
	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	return file
}

func Start(args []string) {
	var fData []byte
	file := fileOpen()
	a, err := file.Read(fData)
	if err != nil {
		fmt.Printf("file reading fkked up: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(a)
	fmt.Println(fData)

	switch args[0] {

	case "create":
		if len(args) != 2 {
			fmt.Println("insufficient args")
		} else {
			create(string(fData))
		}

	default:
		fmt.Println("invalid arguments, run 'help' to get a list of valid arguments.")
	}
}

func create(arg string) {
	fmt.Println("create called")
	fmt.Println(arg)
	dataHandle.JsonRead(arg)
}
