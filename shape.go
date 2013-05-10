package geom2d

// Common interface to Polygons, Line segments, Circles etc.
type Shape interface {
	Inside(q Point) bool
	Intersect(shape Shape) bool
	IntersectPolygon(poly Polygon) bool
	IntersectRect(r Rect) bool
	IntersectCircle(c Circle) bool
	BoundingBox() Rect
}