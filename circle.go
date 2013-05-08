package geom2d

import (
	"math"
)

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
func (c Circle) IntersectCircle(k Circle) bool {
	return k.Center.Sub(c.Center).Norm() < c.Radius + k.Radius
}

// Check if circle and rectangle intersect.
// Not optimized much. 
func (c Circle) IntersectRect(r Rect) bool {
	r2 := c.Radius * c.Radius
	
	// translate coordinate system placing circle at centre
	rmax := r.Max.Sub(c.Center)
	rmin := r.Min.Sub(c.Center)
	
	// left of circle center
	if rmax.X < 0 {
		if rmax.Y < 0 {  // lower left corner
			return rmax.SqrNorm() < r2
		} else if rmin.Y > 0 { // upper left corner
			return rmax.X*rmax.X + rmin.Y*rmin.Y < r2
		} else {
			return math.Abs(rmax.X) < c.Radius
		}
	} else if rmin.X > 0 { // riht of circle centre?
		if rmax.Y < 0 {  // lower right corner
			return rmin.X*rmin.X + rmax.Y*rmax.Y < r2
		} else if rmin.Y > 0 { // upper right corner
			return rmin.SqrNorm() < r2
		} else {
			return rmin.X < c.Radius
		}
	} else { // rectangle on circle vertical centerline
		if rmax.Y < 0 {
			return math.Abs(rmax.Y) < c.Radius
		} else if rmin.Y > 0 {
			return rmin.Y < c.Radius
		} else {
			return true
		}
	}
	
	return false
}

func (c Circle) IntersectPolygon(poly Polygon) bool {
	for i, p := range poly[1:] {
		if c.IntersectSegment(Segment{poly[i - 1], p}) {
			return true
		}
	}
	return c.IntersectSegment(Segment{poly[len(poly) - 1], poly[0]})
}

func (c Circle) IntersectSegment(seg Segment) bool {
	return seg.Distance(c.Center) < c.Radius
}

func (c Circle) Intersect(shape Shape) bool {
	return shape.IntersectCircle(c)
}