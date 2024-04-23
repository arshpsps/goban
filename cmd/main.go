package main

import (
	"goban/internals/dataHandle"
)

func main() {
	// bblt.Run()
	db := dataHandle.Conndb()
	db.GetAllProjects()
}
