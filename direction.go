package geom2d

// A unit vector. Like a 2D vector, always with magnitude 1
type Direction struct {
	X, Y float64
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