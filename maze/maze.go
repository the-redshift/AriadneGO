package maze

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"os"
	"os/exec"
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
	rand.Seed(time.Now().UnixNano())
	// Randomize entrance/exit points
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

	// Create random web of passages from starting point
	m.createPassages(m.entrance)
}

func (m *Maze) createPassages(p Point) {
	rand.Seed(time.Now().UnixNano())

	directions := []direction.Direction{
		direction.NORTH,
		direction.EAST,
		direction.WEST,
		direction.SOUTH,
	}
	rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })

	for _, direction := range directions {
		s_x, s_y := direction.ShiftCoordinates()
		potential_x := int8(p.x) + s_x
		potential_y := int8(p.y) + s_y

		if potential_x < 0 || potential_x >= int8(m.width) {
			continue
		}
		if potential_y < 0 || potential_y >= int8(m.height) {
			continue
		}

		shiftedPoint := Point{uint8(potential_x), uint8(potential_y)}
		if m.cell_matrix[shiftedPoint.x][shiftedPoint.y].Visited {
			continue
		}

		m.cell_matrix[shiftedPoint.x][shiftedPoint.y].Visited = true
		m.cell_matrix[p.x][p.y].CarvePassage(direction)
		m.cell_matrix[shiftedPoint.x][shiftedPoint.y].CarvePassage(direction)

		// Animate labirynth generation
		m.Display()

		// God I hope this recursion works
		m.createPassages(shiftedPoint)
		}
}

func (m Maze) Display() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

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

	// Sleep to make animation smoother
	time.Sleep(25 * time.Millisecond)
}


