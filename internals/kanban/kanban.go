package kanban

import (
	"fmt"
	"goban/internals/dataHandle"
	"os"
	"strings"
)

var fileLocation string = "./data/data.json"

func fileOpen() *os.File {
	file, err := os.OpenFile(fileLocation, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	return file
}

func GrabJsonObj() dataHandle.JsonData {
	fData := make([]byte, 4096)
	file := fileOpen()
	_, err := file.Read(fData)
	if err != nil {
		fmt.Printf("file reading fkked up: %v\n", err)
		os.Exit(1)
	}
	data := strings.Trim(string(fData), "\n \x00")

	return dataHandle.JsonRead(data)
}
