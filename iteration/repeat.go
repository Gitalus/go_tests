package iteration

const repeatCounter = 5

// Repeat 5 times the given character
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCounter; i++ {
		repeated += character
	}
	return repeated
}
