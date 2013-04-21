package geom2d

// Common interface to Polygons, Line segments, Circles etc.
type Shape interface {
	Intersect(poly Polygon) bool
}