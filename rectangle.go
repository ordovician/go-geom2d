package geom2d

//import . "math"

// A 2D rectangle. It is defined by its lower left corner (Min) and
// its upper right corner (Max).
type Rect struct {
	Min, Max Point
}

// A point representing the centre of the rectangle
func (r Rect) Center() Point {
	return Point{(r.Max.X + r.Min.X) * 0.5, (r.Max.Y + r.Min.Y) * 0.5}
}

func (r Rect) Size() Vector2D {
	return r.Max.Sub(r.Min).Abs()
}

func (r Rect) HalfSize() Vector2D {
	return r.Size().Mul(0.5)
}

func (r Rect) Inside(p Point) bool {
	return r.Min.IsMin(p) && r.Max.IsMax(p) || p == r.Min || p == r.Max
}

// Check if the rectangles r and s intersect each other
func (r Rect) Intersect(s Rect) bool {
	d := s.Center().Sub(r.Center())
	h1 := r.HalfSize()
	h2 := s.HalfSize()
	
	return d.X <= h1.X + h2.X && d.Y <= h1.Y + h2.Y
}

func (r Rect) TopLeft() Point {
	return Point{r.Min.X, r.Max.Y}
}

func (r Rect) BottomRight() Point {
	return Point{r.Max.X, r.Min.Y}
}

func (r Rect) TopRight() Point {
	return r.Max
}

func (r Rect) BottomLeft() Point {
	return r.Min
}

// Create a new rectangle transformed by matrix m. Only translation
// and scaling is legal.
func (r Rect) Transform(m Matrix3x3) Rect {
	return Rect{m.TransformPoint(r.Min), m.TransformPoint(r.Max)}
}