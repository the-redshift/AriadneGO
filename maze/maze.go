package maze

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"ariadne/maze/cell"
	"ariadne/maze/cell/direction"
)

type Point struct {
	x uint8
	y uint8
}

func (p Point) String() string {
	return fmt.Sprintf("{x: %d, y: %d}", p.x, p.y)
}

type Maze struct {
	cell_matrix [][]cell.Cell
	entrance Point
	exit Point
	height uint8
	width uint8
}

func New(height uint8, width uint8) (Maze, error) {
	const min_width_height uint8 = 3
	var m Maze

	// Init maze's cell_matrix
	if height < min_width_height || width < min_width_height {
		var err_string = fmt.Sprintf("Minimum value for height and width is: %d", min_width_height)
		return m, errors.New(err_string)
	}

	var cell_matrix = make([][]cell.Cell, height)
	for i := range cell_matrix {
		cell_matrix[i] = make([]cell.Cell, width)
	}

	m.cell_matrix = cell_matrix
	m.width = width
	m.height = height

	// Generate actual values of the maze
	m.generate()

	return m, nil
}

func (m *Maze) generate() {
	// Randomize entrance/exit points
	rand.Seed(time.Now().UnixNano())
	m.entrance = Point{0, uint8(rand.Intn(int(m.width)))}
	m.exit = Point{uint8(m.height - 1), uint8(rand.Intn(int(m.width)))}

	// Fill out cell matrix with initial cells with complete borders
	for i := range m.cell_matrix {
		for j := range m.cell_matrix[i] {
			m.cell_matrix[i][j] = cell.New()
		}
	}

	// Removing appropriate walls from entrance/exit
	m.cell_matrix[m.entrance.x][m.entrance.y].CarvePassage(direction.NORTH)
	m.cell_matrix[m.exit.x][m.exit.y].CarvePassage(direction.SOUTH)

}

func (m Maze) Display() {
	fmt.Println("[Entrance]", m.entrance)
	fmt.Println("[Exit]", m.exit)

	// Upper border, always solid
	for i := uint8(0); i < m.width; i++ {
		if i == m.entrance.y {
			fmt.Printf("  ")
		}	else {
			fmt.Printf(" _")
		}
	}

	// Then we display borders for each cell
	for i := uint8(0); i < m.height; i++ {
		fmt.Println()

		for j := uint8(0); j < m.width; j++ {
			fmt.Printf("%s", m.cell_matrix[i][j])

			if j == m.width - 1 {
				fmt.Printf("|")
			}
		}
	}

	fmt.Println()
}


