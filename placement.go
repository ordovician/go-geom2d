package geom2d

// A representation of something with a position and direction.
// That is similar to a ray and plane, but those are geometric objects
// supporting different operations.
type Placement struct {
	Pos Point
	Dir Direction
}