package models

type Bearing int

const (
	North Bearing = iota
	South
	East
	West
	NorthWest
	SouthWest
	NorthEast
	SouthEast
)

func (b Bearing) String() string {
	to_string := ""
	switch b {
	case North:
		to_string = "North"
	case South:
		to_string = "South"
	case East:
		to_string = "East"
	case West:
		to_string = "West"
	case NorthWest:
		to_string = "NorthWest"
	case SouthWest:
		to_string = "SouthWest"
	case NorthEast:
		to_string = "NorthEast"
	case SouthEast:
		to_string = "SouthEast"
	}
	return to_string
}

func AllDirs() []Bearing {
	return []Bearing{North, South, West, East,
		NorthWest, SouthWest, NorthEast, SouthEast}
}

func MoveCoord(loc Coord, dir Bearing, amount int) Coord {
	switch dir {
	case North:
		loc.X = loc.X - amount
	case South:
		loc.X = loc.X + amount
	case East:
		loc.Y = loc.Y + amount
	case West:
		loc.Y = loc.Y - amount

	case NorthWest:
		loc.X = loc.X - amount
		loc.Y = loc.Y - amount
	case SouthWest:
		loc.X = loc.X + amount
		loc.Y = loc.Y - amount
	case NorthEast:
		loc.X = loc.X - amount
		loc.Y = loc.Y + amount
	case SouthEast:
		loc.X = loc.X + amount
		loc.Y = loc.Y + amount
	}
	return loc
}

func (b Bearing) TurnClockwiseBy90() Bearing {
	turned := North 
	switch b {
	case North:
        turned = East
	case South:
        turned = West 
	case East:
        turned = South 
	case West:
        turned = North 
	}
	return turned 
}
