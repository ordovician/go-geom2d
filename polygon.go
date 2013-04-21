package geom2d

import (
	"math"
)

type Polygon []Point

// Get a list of all the normals on each polygon edge
func (poly Polygon) sepAxis() []Direction {
	out := make([]Direction, len(poly))
	last := poly[len(poly) - 1]
	
	for i, dst := range poly {
		src := last
		if i > 0 {
			src = poly[i - 1]
		}
		diff := dst.Sub(src)
		out[i] = diff.Unit().Normal()
	}
	
	return out
}

// Project all the points of the polygon onto axis
func (poly Polygon) Project(axis Direction) []float64 {
	out := make([]float64, len(poly))
	for i, p := range poly {
		out[i] = p.Project(axis)
	}
	return out
}

// find the smallest of all values
func minValue(values []float64) float64 {
	result := math.MaxFloat64
	for _, v := range values {
		result = math.Min(result, v)
	}
	return result
}

// find largest of all values
func maxValue(values []float64) float64 {
	result := math.MaxFloat64
	for _, v := range values {
		result = math.Max(result, v)
	}
	return result
}

// Check if the polygons a and b overlap. 
func (a Polygon) Intersect(b Polygon) bool {
	sepAxis := a.sepAxis()
	sepAxis = append(sepAxis, b.sepAxis()...)
	
	for _, axis := range sepAxis {
		aProj := a.Project(axis)
		bProj := b.Project(axis)
		aMin := minValue(aProj)
		aMax := maxValue(aProj)
		bMin := minValue(bProj)
		bMax := maxValue(bProj)
		
		if aMin > bMax || aMax < bMin {
			return false			
		}

	}

	return true
}