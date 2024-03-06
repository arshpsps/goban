package main

import (
	"goban/internals/bblt"
	// "goban/internals/kanban"
	// "log"
	// "os"
	// tea "github.com/charmbracelet/bubbletea"
)

func main() {
	bblt.Run()
}

// func pain() {
// 	if len(os.Args) < 2 {
// 		log.Fatalf("insufficient args: %v", os.Args)
// 	}
//
// 	// file, err := os.OpenFile("../data/data.json", os.O_RDWR, 0777)
// 	// if err != nil {
// 	//     log.Fatalf("error opening file: %s", err)
// 	// }
//
// 	kanban.Start(os.Args[1:])
// }
