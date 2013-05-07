package geom2d

type Segment struct {
	Source, Target Point
}

func (s Segment) IsDegenerate() bool {
	return s.Source.Eq(s.Target)
}

func (s Segment) pointPlacement(p Point) float64 {
	min := s.Min()
	// max(s) - min(s) cross p - min(s)
	return s.Max().Sub(min).Cross(p.Sub(min))
}

// Is point p below line defined by segment
func (s Segment) IsAbove(p Point) bool {
	return s.pointPlacement(p) < 0
}

// Is  point p above line defined by segment
func (s Segment) IsBelow(p Point) bool {
	return s.pointPlacement(p) > 0
}

// Is p on line defined by segment
func (s Segment) IsOn(p Point) bool {
	return s.pointPlacement(p) == 0
}

// Is p at one of the segment end points
func (s Segment) IsOnEnd(p Point) bool {
	return p.Eq(s.Source) || p.Eq(s.Target)
}

// Check if s and t intersect each other
func (seg0 Segment) IntersectSegment(seg1 Segment) bool {
	d := seg0.Target.Sub(seg0.Source)
	v := seg1.Vector()
	
	num := seg0.Source.Sub(seg1.Source)
	denom := d.Cross(v)
	
	if denom == 0 {
		return false
	}
	
	t := v.Cross(num) / denom
	s := d.Cross(num) / denom
	
	return !(t > 1 || t < 0 || s > 1 || s < 0)
}

// Check is segment s intersects rectangle r
func (s Segment) IntersectRect(r Rect) bool {
	return	s.IntersectSegment(Segment{r.BottomLeft() , r.TopLeft()}) 		||
			s.IntersectSegment(Segment{r.TopLeft()    , r.TopRight()})   	||
			s.IntersectSegment(Segment{r.BottomLeft() , r.BottomRight()})	||
			s.IntersectSegment(Segment{r.BottomRight(), r.TopRight()})		
} 


// The lexiographically smallest of the two endpoints of the segment
func (s Segment) Min() Point {
	return s.Source.Min(s.Target)
}

// The lexiographically largest of the two endpoints of the segment
func (s Segment) Max() Point {
	return s.Source.Max(s.Target)
}

// Swap endpoints
func (s Segment) Opposite() Segment {
	return Segment{s.Target, s.Source}
}

// The segment as a vector from source to target endpoint
func (s Segment) Vector() Vector2D {
	return s.Target.Sub(s.Source)
}