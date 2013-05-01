package geom2d

import (
	"math"
)

type Radian float64
type Degree float64

// Convert radians to degrees
func Deg(radians Radian) Degree {
	return Degree(radians * 180/math.Pi)
}

// Convert degrees to radians
func Rad(degrees Degree) Radian {
	return Radian(degrees * math.Pi/180)
}

// A unit vector. Like a 2D vector, always with magnitude 1
type Direction struct {
	X, Y float64
}

func DirectionWithAngle(angle Radian) Direction{
	ang := float64(angle)
	return Direction{math.Cos(ang), math.Sin(ang)}
}

// Magnitude of direction. Will always return 1
func (d Direction) Norm() float64 {
	return 1.0
}

// Square magnitude of direction. Will always return 1.
func (d Direction) SqrNorm() float64 {
	return 1.0
}

// Get the unit vector. Will return self.
func (d Direction) Unit() Direction {
	return d
}

// Create a vector with same direction as d but with magnitude factor.
func (d Direction) Mul(factor float64) Vector2D {
	return Vector2D{d.X * factor, d.Y * factor}
}

func (d Direction) Normal() Direction {
	return Direction{-d.Y, d.X}
}

func (d Direction) Angle() Radian {
	return Radian(math.Atan2(d.Y, d.X))
}