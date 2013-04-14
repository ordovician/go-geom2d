package geom2d

// A circle in 2D space. A circle is a simple shape in Euclidean geometry.
// All points have the same distance to the centre.
type Circle struct {
	Center Point
	Radius float64
}

// True if point p is inside circle c. If p is on
// the circle it will not be counted as inside.
func (c Circle) Inside(p Point) bool {
	return p.Sub(c.Center).Norm() < c.Radius
}

// Check if two circles overlap
func (c Circle) Intersect(k Circle) bool {
	return k.Center.Sub(c.Center).Norm() < c.Radius + k.Radius
}
