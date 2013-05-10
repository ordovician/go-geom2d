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

// Width and height of rectangle
func (r Rect) Size() Vector2D {
	return r.Max.Sub(r.Min).Abs()
}

// half the width and height of rectangle
func (r Rect) HalfSize() Vector2D {
	return r.Size().Mul(0.5)
}

// Check if point p is inside rectangle r. Point is considered inside
// if it is on the border. 
func (r Rect) Inside(p Point) bool {
	return r.Min.IsMin(p) && r.Max.IsMax(p) || p == r.Min || p == r.Max
}

// Check if the rectangles r and s intersect each other
func (r Rect) IntersectRect(s Rect) bool {
	d := s.Center().Sub(r.Center())
	h1 := r.HalfSize()
	h2 := s.HalfSize()
	
	return d.X <= h1.X + h2.X && d.Y <= h1.Y + h2.Y
}

func (r Rect) IntersectCircle(c Circle) bool {
	return c.IntersectRect(r)
}

func (r Rect) IntersectPolygon(poly Polygon) bool {
	return poly.IntersectRect(r)
}

func (r Rect) Intersect(shape Shape) bool {
	return shape.IntersectRect(r)
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

// Transform rectangle r with matrix m. Only translation
// and scaling is legal.
func (r *Rect) Transform(m Matrix3x3) {
	r.Min = m.TransformPoint(r.Min)
	r.Max = m.TransformPoint(r.Max)
}

func (r *Rect) Move(delta Vector2D) {
	r.Min = r.Min.Add(delta)
	r.Max = r.Max.Add(delta)
}

func (r *Rect) MoveTo(pos Point) {
	size := r.Size()
	r.Max = pos.Add(size)
	r.Min = pos
}

func (r *Rect) MoveCenter(pos Point) {
	r.Move(pos.Sub(r.Center()))
}

// The smallest rectangle which can contain both r and point p
func (r Rect) SurroundPoint(p Point) Rect {
	return Rect{r.Min.MinComp(p),
				r.Max.MaxComp(p)}
}

// The smallest rectangle containing both rectangles r and s
func (r Rect) SurroundRect(s Rect) Rect {
	return Rect{r.Min.MinComp(s.Min),
				r.Max.MaxComp(s.Max)}
}

// The rect is the same as its own bounding box
func (r Rect) BoundingBox() Rect {
	return r
}
