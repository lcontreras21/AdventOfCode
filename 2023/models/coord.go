package models

import (
	"fmt"
)

type Coord struct {
	X int
	Y int
	Z int

}

func (c Coord) String() string {
	return fmt.Sprintf("[%d, %d, %d]", c.X, c.Y, c.Z)
}

func (c Coord) From(values [3]int) Coord {
    return Coord{X: values[0], Y: values[1], Z: values[2]}
}

type CoordFloat64 struct {
	X float64
	Y float64
	Z float64

}

func (c CoordFloat64) String() string {
	return fmt.Sprintf("[%f, %f, %f]", c.X, c.Y, c.Z)
}

func (c CoordFloat64) From(values [3]float64) CoordFloat64 {
    return CoordFloat64{X: values[0], Y: values[1], Z: values[2]}
}

func (c *CoordFloat64) ToArray() []float64 {
    return []float64{c.X, c.Y, c.Z}
}

func (c *CoordFloat64) Add(other CoordFloat64) (result CoordFloat64) {
    result.X = c.X + other.X
    result.Y = c.Y + other.Y
    result.Z = c.Z + other.Z
    return
}

func (c *CoordFloat64) Sub(other CoordFloat64) (result CoordFloat64) {
    result.X = c.X - other.X
    result.Y = c.Y - other.Y
    result.Z = c.Z - other.Z
    return
}

func (c *CoordFloat64) Dot(other CoordFloat64) (float64) {
    return (c.X * other.X) + (c.Y * other.Y) + (c.Z * other.Z)
}

func (c *CoordFloat64) SkewMatrixOperator() (matrix []CoordFloat64) {
    matrix = []CoordFloat64{
        {X:  0,   Y:-c.Z, Z: c.Y},
        {X:  c.Z, Y: 0,   Z: -c.X},
        {X: -c.Y, Y: c.X, Z: 0},
    }

    return
}

func (c *CoordFloat64) CrossProduct(other CoordFloat64) (result CoordFloat64) {
    c_skewed := c.SkewMatrixOperator() 
    result.X = c_skewed[0].Dot(other)
    result.Y = c_skewed[1].Dot(other)
    result.Z = c_skewed[2].Dot(other)

    return
}
