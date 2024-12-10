package models

import (
	"fmt"
)

type Vector struct {
	Loc Coord
	Dir Bearing
}

func (v Vector) String() string {
	return v.Dir.String() + fmt.Sprintf(" - [%d, %d, %d]", v.Loc.X, v.Loc.Y, v.Loc.Z)
}
