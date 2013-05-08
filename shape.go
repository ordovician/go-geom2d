package geom2d

// Common interface to Polygons, Line segments, Circles etc.
type Shape interface {
	Intersect(shape Shape) bool
	IntersectPolygon(poly Polygon) bool
	IntersectRect(r Rect) bool
	IntersectCircle(c Circle) bool
}