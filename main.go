package main

import (
	"ariadne/maze"
)

func main() {
	m, err := maze.New(25, 25)
	if err == nil {
		m.Display()
	}
}
