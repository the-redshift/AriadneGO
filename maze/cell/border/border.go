package border

type Border uint8

const (
	PASSAGE Border = iota
	WALL
)

func (b Border) String() string {
	switch b {
	case WALL:
		return "WALL"
	case PASSAGE:
		return "PASSAGE"
	}

	return "INVALID"
}
