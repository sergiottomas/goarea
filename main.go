package goarea

//Pi is a numeric proportion defined by relation between
//the perimeter of a circumference and its perimeter
const Pi = 3.1416

//Circle is responsible to calculate the area of circuference
func Circle(radius float64) float64 {
	return radius * Pi
}

//Rect is responsible to calculate the area of the rectangle
func Rect(base, height float64) float64 {
	return base * height
}

func _TriangleEq(base, height float64) float64 {
	return (base * height) / 2
}
