package geom2d

import . "math"

// A 3x3 matrix for transforming 2D polygons, vectors etc.
// The matrix follows column-major order, like OpenGL:
//   m0 m3 m6
//   m1 m4 m7
//   m2 m5 m8
// The rotation and translation the matrix represents is given as:
//   cos(𝛩)  -sin(𝛩)   x
//   sin(𝛩)   cos(𝛩)   y
//     0        0      1 
type Matrix3x3 [9]float64

func NewMatrix(pos Point, ang float64) Matrix3x3 {
	m := Identity()
    m[6] = pos.X
    m[7] = pos.Y
  
    cosine := Cos(ang)
    sine := Sin(ang)  

    m[0] = cosine
    m[3] = -sine
    m[1] = sine
    m[4] = cosine 
  
    return m	
}

// The identity matrix. Zero excep the diagonal which is 1
func Identity() Matrix3x3 {
	var m Matrix3x3
	m[0] = 1
	m[4] = 1
	m[8] = 1
	return m
}

// A translation matrix.
// Will translate points along x and y axis according to vector v
func Translate(v Vector2D) Matrix3x3 {
	m := Identity()
	m[6] = v.X
	m[7] = v.Y
	return m
}

// A scaling matrix. Use it to scale a polygon.
// When multiplying each point in a polygon with
// matrix it will have the result of scaling the polygon along x and y axis
// as given.
func Scale(x, y float64) Matrix3x3 {
	var m Matrix3x3
	m[0] = x
	m[4] = y
	m[8] = 1
	return m
}

// Rotation matrix. Angles radians.
func Rotate(ang float64) Matrix3x3 {
	m := Identity()
  
    cosine := Cos(ang)
    sine := Sin(ang)  

    m[0] = cosine
    m[3] = -sine
    m[1] = sine
    m[4] = cosine 
  
    return m	
}

// A Matrix were every element is set to 1. 
func One() Matrix3x3 {
	return Matrix3x3{1, 1, 1,
					 1, 1, 1,
				     1, 1, 1}
}

// Compare two matricies for equality
func (m Matrix3x3) Equal(n Matrix3x3) bool {
	for i := 0; i < 9; i++ {
		if (m[i] != n[i]) {
			return false
		}
	}
	return true
}

// Get the translation part of a matrix.
// If the matrix is used to represent the position and orientation
// of a shape then this returns the position part.
func (m Matrix3x3) Pos() Point {
	return Point{m[6], m[7]}
}

// Set the translation part of a matrix.
func (m *Matrix3x3) SetPos(pos Point) {
	m[6] = pos.X
	m[7] = pos.Y
}

// Get the rotation part of a matrix.
// If matrix is used to represent position and orientation of
// a shape then this returns a direction vector to indicate
// orientation.
func (m Matrix3x3) Dir() Direction {
	return Direction{m[0], m[1]}
}

func (m Matrix3x3) TransformPoint(p Point) Point {
	return Point(m.MulVec(Vector2D(p)))
}

func (m Matrix3x3) TransformVector(v Vector2D) Vector2D {
	return  Vector2D{v.X*m[0] + v.Y*m[3],
					 v.X*m[1] + v.Y*m[4]}
}

// Get the transpose of a matrix.
// Makes rows into columns and vica versa.
func (m Matrix3x3) GetTranspose() Matrix3x3 {
	return Matrix3x3{m[0], m[3], m[6],
					 m[1], m[4], m[7],
				     m[2], m[5], m[8]}
}

// Add two vectors
func (m Matrix3x3) Add(n Matrix3x3)  Matrix3x3 {
	for i := 0; i < 9; i++ {
		m[i] += n[i]
	}
	return m
}

// Substract a vector from another
func (m Matrix3x3) Sub(n Matrix3x3)  Matrix3x3 {
	for i := 0; i < 9; i++ {
		m[i] -= n[i]
	}
	return m
}

// Matrix multiplication
func (m Matrix3x3) Mul(n Matrix3x3) Matrix3x3 {
	var res Matrix3x3
	
	// row is index of first element in a row of matrix m.
	// row elements are spaced 3 indices apart
	for row := 0; row < 3; row++ {
		// col is index of first element in column of matrix n
		// column elements are space 1 index apart
		for col := 0; col < 9; col += 3 {
			sum := m[row]*n[col] + m[row+3]*n[col+1] + m[row+6]*n[col+2]
			res[col+row] = sum
		}
	}
	return res
}

// Multiply matrix m with vector v
func (m Matrix3x3) MulVec(v Vector2D) Vector2D {
	var u Vector2D
	u.X = v.X*m[0] + v.Y*m[3] + m[6]
	u.Y = v.X*m[1] + v.Y*m[4] + m[7]
	return u
}

// Multiply each element of a matrix by a factor
func (m Matrix3x3) MulFac(fac float64) Matrix3x3 {
	for i := 0; i < 9; i++ {
		m[i] *= fac
	}
	return m
}