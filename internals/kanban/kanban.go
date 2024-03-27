package kanban

import (
	"fmt"
	"goban/internals/dataHandle"
	"os"
	"strings"
)

var (
	fileLocation string   = "./data/data.json"
	fData        []byte   = make([]byte, 4096)
	file         *os.File = fileOpen()
)

func fileOpen() *os.File {
	file, err := os.OpenFile(fileLocation, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	return file
}

func GrabJsonObj() dataHandle.JsonData {
	_, err := file.Read(fData)
	if err != nil {
		fmt.Printf("file reading fkked up: %v\n", err)
		os.Exit(1)
	}
	data := strings.Trim(string(fData), "\n \x00")

	return dataHandle.JsonRead(data)
}
