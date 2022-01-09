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
	"github.com/fatih/color"
)

type Point struct {
	X uint8
	Y uint8
}

func (p Point) String() string {
	return fmt.Sprintf("{x: %d, y: %d}", p.X, p.Y)
}

type Maze struct {
	Cell_matrix [][]cell.Cell
	Entrance Point
	Exit Point
	Height uint8
	Width uint8
}

func New(Height uint8, Width uint8) (Maze, error) {
	const min_Width_Height uint8 = 3
	var m Maze

	// Init maze's Cell_matrix
	if Height < min_Width_Height || Width < min_Width_Height {
		var err_string = fmt.Sprintf("Minimum value for Height and Width is: %d", min_Width_Height)
		return m, errors.New(err_string)
	}

	var Cell_matrix = make([][]cell.Cell, Height)
	for i := range Cell_matrix {
		Cell_matrix[i] = make([]cell.Cell, Width)
	}

	m.Cell_matrix = Cell_matrix
	m.Width = Width
	m.Height = Height

	// Generate actual values of the maze
	m.generate()

	return m, nil
}

func (m *Maze) generate() {
	rand.Seed(time.Now().UnixNano())
	// Randomize Entrance/Exit points
	m.Entrance = Point{0, uint8(rand.Intn(int(m.Width)))}
	m.Exit = Point{uint8(m.Height - 1), uint8(rand.Intn(int(m.Width)))}

	// Fill out cell matrix with initial cells with complete borders
	for i := range m.Cell_matrix {
		for j := range m.Cell_matrix[i] {
			m.Cell_matrix[i][j] = cell.New()
		}
	}

	// Removing appropriate walls from Entrance/Exit
	m.Cell_matrix[m.Entrance.X][m.Entrance.Y].CarvePassage(direction.NORTH)
	m.Cell_matrix[m.Exit.X][m.Exit.Y].CarvePassage(direction.SOUTH)

	// Create random web of passages from starting point
	m.createPassages(m.Entrance)

	// Reset 'Visited' values as Ariadne uses them to find path
	for i := range m.Cell_matrix {
		for j := range m.Cell_matrix[i] {
			m.Cell_matrix[i][j].Visited = false
		}
	}
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
		potential_x := int8(p.X) + s_x
		potential_y := int8(p.Y) + s_y

		if potential_x < 0 || potential_x >= int8(m.Width) {
			continue
		}
		if potential_y < 0 || potential_y >= int8(m.Height) {
			continue
		}

		shiftedPoint := Point{uint8(potential_x), uint8(potential_y)}
		if m.Cell_matrix[shiftedPoint.X][shiftedPoint.Y].Visited {
			continue
		}

		m.Cell_matrix[shiftedPoint.X][shiftedPoint.Y].Visited = true
		m.Cell_matrix[p.X][p.Y].CarvePassage(direction)
		m.Cell_matrix[shiftedPoint.X][shiftedPoint.Y].CarvePassage(direction)

		// Animate labirynth generation
		// m.Display()

		// God I hope this recursion works
		m.createPassages(shiftedPoint)
		}
}

func (m Maze) Display(path []Point) {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	// Upper border, always solid
	for i := uint8(0); i < m.Width; i++ {
		if i == m.Entrance.Y {
			fmt.Printf("  ")
		}	else {
			fmt.Printf(" _")
		}
	}

	// Then we display borders for each cell
	for i := uint8(0); i < m.Height; i++ {
		fmt.Println()

		for j := uint8(0); j < m.Width; j++ {
			matched := false
			for z := range path {
				pathPoint := path[z]
				if pathPoint.X == i && pathPoint.Y == j {
					matched = true
					break
				}
			}

			if matched {
				color.Set(color.BgRed)
			} else {
				color.Unset()
			}

			fmt.Printf("%s", m.Cell_matrix[i][j])

			if j == m.Width - 1 {
				fmt.Printf("|")
			}
		}
	}

	fmt.Println()

	// Sleep to make animation smoother
	time.Sleep(25 * time.Millisecond)
}


