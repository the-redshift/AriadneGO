package maze

import (
	"time"
	"math/rand"
	"ariadne/maze"
	"ariadne/maze/cell/direction"
	"ariadne/maze/cell/border"
)

type Ariadne struct {
	Path []maze.Point
  maze maze.Maze
}

func New(m maze.Maze) Ariadne {
	var path []maze.Point
	ariadne := Ariadne{path, m}
	return ariadne
}

func (a *Ariadne) FindPath() {
	rand.Seed(time.Now().UnixNano())
	var path []maze.Point
	a.traverseMaze(a.maze.Entrance, path)
}

func (a *Ariadne) traverseMaze(p maze.Point, path []maze.Point) bool {
	a.maze.Cell_matrix[p.X][p.Y].Visited = true
	path = append(path, p)
	a.maze.Display(path)

	if p == a.maze.Exit {
		a.Path = path
		return true
	}

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

    if potential_x < 0 || potential_x >= int8(a.maze.Width) {
      continue
    }
    if potential_y < 0 || potential_y >= int8(a.maze.Height) {
      continue
    }
		if a.maze.Cell_matrix[p.X][p.Y].Borders[direction] == border.WALL {
			continue
		}

    shiftedPoint := maze.Point{uint8(potential_x), uint8(potential_y)}
    if a.maze.Cell_matrix[shiftedPoint.X][shiftedPoint.Y].Visited {
      continue
    }

		if a.traverseMaze(shiftedPoint, path) {
			return true
		}
	}

	return false
}



