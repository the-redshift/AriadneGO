package cell

import (
	"fmt"
	"errors"
	"ariadne/maze/cell/direction"
	"ariadne/maze/cell/border"
)

type Cell struct {
	Visited bool
	Borders map[direction.Direction]border.Border
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
	c.Borders = b

	return c
}

func (c Cell) Display() {
	fmt.Println("--- Cell Contents ---")
	fmt.Println("[Visited]", c.Visited)
	for k, v := range c.Borders {
		fmt.Printf("[%s] %s\n", k, v)
	}
}

func (c Cell) String() string {
	var output = ""

	if c.Borders[direction.WEST] == border.WALL {
		output += "|"
	} else {
		output += " "
	}

	if c.Borders[direction.SOUTH] == border.WALL {
		output += "_"
	} else {
		output += " "
	}

	return output
}

func (c Cell) CarvePassage(d direction.Direction) error {
	if _, exists := c.Borders[d]; !exists {
		return errors.New("Invalid direction. Accepted values: North, East, West, South.")
	}

	c.Borders[d] = border.PASSAGE
	return nil
}


