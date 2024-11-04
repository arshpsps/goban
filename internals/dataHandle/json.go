package dataHandle

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type JsonData struct {
	Projects []Project
}

func JsonRead(data string) JsonData {
	var jsonData JsonData
	if !json.Valid([]byte(data)) {
		log.Fatalf("invalid JSON: %s", data)
	}
	json.Unmarshal([]byte(data), &jsonData)
	return jsonData
}

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

func GrabJsonObj() JsonData {
	_, err := file.Read(fData)
	if err != nil {
		fmt.Printf("file reading fkked up: %v\n", err)
		os.Exit(1)
	}
	data := strings.Trim(string(fData), "\n \x00")

	return JsonRead(data)
}
