package property_based

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	// A Builder is used to efficiently build a string using Write methods. It minimizes memory copying.
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	// for arabic > 0 {
	// 	switch {
	// 	case arabic > 9:
	// 		result.WriteString("X")
	// 		arabic -= 10
	// 	case arabic > 8:
	// 		result.WriteString("IX")
	// 		arabic -= 9
	// 	case arabic > 4:
	// 		result.WriteString("V")
	// 		arabic -= 5
	// 	case arabic > 3:
	// 		result.WriteString("IV")
	// 		arabic -= 4
	// 	default:
	// 		result.WriteString("I")
	// 		arabic--
	// 	}
	// }

	return result.String()
}
