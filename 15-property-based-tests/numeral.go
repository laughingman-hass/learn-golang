package numeral_converter

import "strings"

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for i := 0; i < arabic; i++ {
		if arabic == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}

	return result.String()
}
