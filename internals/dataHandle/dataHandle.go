package dataHandle

import (
	"encoding/json"
	"log"
)

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
