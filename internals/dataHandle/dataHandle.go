package dataHandle

import (
	"encoding/json"
	"fmt"
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

func JsonRead(data string) (JsonData, error) {
	var jsonData JsonData
	if !json.Valid([]byte(data)) {
		fmt.Println("invalid JSON:", data)
		// return *new(JsonData), errors.New(`invalid JSON:, ${data}`)
		return *new(JsonData), fmt.Errorf("invalid JSON:, %s", data)
	}
	json.Unmarshal([]byte(data), &jsonData)
	fmt.Println(jsonData.Projects[0].Boards[0].Cards[0].Tags[0])
	return jsonData, nil
}
