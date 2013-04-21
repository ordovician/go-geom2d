package geom2d

import "math"
// NOTE: by embedding Point, Vector2D or direction in each other
// we might be able to share functionality the way we wanted
 

// type Point struct {
// 	X, Y float64
// }

type Point Vector2D

// Find the vector between two points
func (p Point) Sub(q Point) Vector2D {
	return Vector2D{q.X - p.X, q.Y - p.Y}
}

// Translate point by adding a vector
func (p Point) Add(q Vector2D) Point {
	return Point{q.X + p.X, q.Y + p.Y}
}

// Absolute value of point
func (p Point) Abs() Point {
	return Point{math.Abs(p.X), math.Abs(p.Y)}
}

// Point p lexiographically smaller than Point q.
// First compare X then Y.
func (p Point) IsMin(q Point) bool {
	return p.X < q.X || (p.X == q.X && p.Y < q.Y)
}

// Point p lexiographically larger than Point q.
// First compare X then Y.
func (p Point) IsMax(q Point) bool {
	return p.X > q.X || (p.X == q.X && p.Y > q.Y)
}

// Project point onto axis.
// Projecting a point onto an axis, is the same as finding
// the spot closest to the point along the axis. 
// The number returned is how far it is from origo to this spot
func (p Point) Project(axis Direction) float64 {
	return Vector2D(p).Dot(Vector2D(axis))
}

