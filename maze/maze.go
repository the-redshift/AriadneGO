package maze

import (
	"errors"
	"fmt"
)

type Maze struct {
	array [][]int
}

func New(height int, width int) (Maze, error) {
	const min_width_height int = 3
	var m Maze

	if height < min_width_height || width < min_width_height {
		var err_string = fmt.Sprintf("Minimum value for height and width is: %d", min_width_height)
		return m, errors.New(err_string)
	}

	var array = make([][]int, height)
	for i := range array {
		array[i] = make([]int, width)
	}

	m.array = array

	return m, nil
}

func (m Maze) Display() {
	for i := range m.array {
		for j := range m.array[i] {
			fmt.Printf("%d ", m.array[i][j])
		}
		fmt.Printf("\n")
	}
}


