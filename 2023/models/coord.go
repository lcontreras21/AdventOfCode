package models

import "fmt"

type Coord struct {
	X int
	Y int
	Z int

}

func (c Coord) String() string {
	return fmt.Sprintf("[%d, %d, %d]", c.X, c.Y, c.Z)
}
