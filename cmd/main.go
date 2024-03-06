package main

import (
    "os"
    "log"
    "goban/internals/kanban"
    // tea "github.com/charmbracelet/bubbletea"
)

type model struct{

}

func main() {
    if len(os.Args) < 2 {
        log.Fatalf("insufficient args: %v", os.Args)
    }

    // file, err := os.OpenFile("../data/data.json", os.O_RDWR, 0777)
    // if err != nil {
    //     log.Fatalf("error opening file: %s", err)
    // }

    kanban.Start(os.Args[1:])
}
