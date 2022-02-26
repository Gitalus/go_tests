package structs_methods_interface

type Rectangle struct {
	width  float64
	height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

// area can be in lowecase because is only being used on the same package
func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
