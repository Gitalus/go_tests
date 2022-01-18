package iteration

import "strings"

const repeatCounter = 5

// Repeat 5 times the given character
func Repeat(character string) string {
	var repeated string = strings.Repeat(character, repeatCounter)

	return repeated
}
