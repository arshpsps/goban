package main

import (
    "os"
    "log"
    "goban/internals"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatalf("insufficient args: %v", os.Args)
    }
    kanban.Start(os.Args[1:])
}
