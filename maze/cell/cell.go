package cell

import (
	"fmt"
	"ariadne/maze/cell/direction"
	"ariadne/maze/cell/border"
)

type Cell struct {
	Visited bool
	borders map[direction.Direction]border.Border
}

func New() Cell {
	var c Cell
	b := map[direction.Direction]border.Border {
		direction.NORTH: border.WALL,
		direction.EAST: border.WALL,
		direction.WEST: border.WALL,
		direction.SOUTH: border.WALL,
	}

	c.Visited = false
	c.borders = b

	return c
}

func (c Cell) Display() {
	fmt.Println("--- Cell Contents ---")
	fmt.Println("[Visited]", c.Visited)
	for k, v := range c.borders {
		fmt.Printf("[%s] %s\n", k, v)
	}
}


