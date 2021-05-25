package area

import "math"

const Pi = 3.1416

func Circ(radius float64) float64 {
	return math.Pow(radius, 2) * Pi
}

func Rect(width, height float64) float64 {
	return width * height
}

func _TriangleEq(width, height float64) float64 {
	return width * height / 2
}
