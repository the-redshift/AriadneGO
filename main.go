package main

import (
	"ariadne/maze"
	ariadne "ariadne/maze/ariadne"
)

func main() {
	m, err := maze.New(25, 25)
	if err == nil {
		a := ariadne.New(m)
		a.FindPath()
	}
}
