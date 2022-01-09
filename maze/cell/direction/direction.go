package direction

type Direction uint8

const (
	NORTH Direction = iota
	EAST
	WEST
	SOUTH
)

func (d Direction) String() string {
	switch d {
	case NORTH:
		return "N"
	case EAST:
		return "E"
	case WEST:
		return "W"
	case SOUTH:
		return "S"
	}

	return "INVALID"
}

func (d Direction) Opposite() Direction {
	oppositeDirection := map[Direction]Direction {
		NORTH: SOUTH,
		EAST: WEST,
		WEST: EAST,
		SOUTH: NORTH,
	}

	return oppositeDirection[d]
}

func (d Direction) ShiftCoordinates() (int8, int8) {
	switch d {
	case NORTH:
		return -1, 0
	case EAST:
		return 0, 1
	case WEST:
		return 0, -1
	case SOUTH:
		return 1, 0
	}

	return -1, -1
}
