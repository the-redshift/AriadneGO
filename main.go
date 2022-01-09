package main

import (
	"ariadne/maze/cell"
	"ariadne/maze/cell/direction"
)

func main() {
	var c = cell.New()
	c.Display()
	c.CarvePassage(direction.WEST)
	c.Display()
}
