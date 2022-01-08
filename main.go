package main

import (
	"log"
	"ariadne/maze"
)

func main() {
	var m, err = maze.New(5, 5)

	if err != nil {
		log.Fatal(err)
	}

	m.Display()
}
