package structs_methods_interface

func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

// area can be in lowecase because is only being used on the same package
func Area(width, height float64) float64 {
	return width * height
}
