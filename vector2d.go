package geom2d

import . "math"

// A geometric vector for 2D space. A vector has a direction and magnitude.
// Also known as an Euclidean vector. You can perform vector algebra on
// Vector2D. That is you can add, substract, perform dot or cross product.
type Vector2D struct {
	X, Y float64
}

// Number of components. Always 2 for 2D vector.
func (v Vector2D) Length() int {
	return 2
}

// Magnitude of vector.
func (v Vector2D) Norm() float64 {
	return Sqrt(v.X*v.X + v.Y*v.Y)
}

// Get the unit vector of v. 
func (v Vector2D) Unit() Direction {
	len := v.Norm()
	return Direction{v.X/len, v.Y/len}
}

func (v Vector2D) Abs() Vector2D {
	return Vector2D{Abs(v.X), Abs(v.Y)}
}

// Substract vector v from u
func (v Vector2D) Sub(u Vector2D) Vector2D {
	return Vector2D{u.X - v.X, u.Y - v.Y}
}

// Add vector v to vector u.
func (v Vector2D) Add(u Vector2D ) Vector2D {
	return Vector2D{v.X + u.X, v.Y + u.Y}
}

// Multiply vector with a factor.
// Direction will remain unchanged but magnitude will increase.
func (v Vector2D) Mul(factor float64) Vector2D {
	return Vector2D{v.X * factor, v.Y * factor}
}

// The dot product of the two vectors v and u
func (v Vector2D) Dot(u Vector2D) float64 {
	return v.X*u.X + v.Y*u.Y
}

// The cross product of the two vectors v and u
func (v Vector2D) Cross(u Vector2D) float64 {
	return v.X*u.Y + v.Y*u.X
}

func (v Vector2D) Normal() Vector2D {
	return Vector2D{-v.Y, v.X}
}