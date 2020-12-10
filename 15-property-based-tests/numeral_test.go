package numeral_converter

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description   string
		Arabic        int
		expectedRoman string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			actual := ConvertToRoman(test.Arabic)
			if actual != test.expectedRoman {
				t.Errorf("actual %q, expected %q", actual, test.expectedRoman)
			}
		})
	}
}
