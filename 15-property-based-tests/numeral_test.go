package numeral

import "testing"

func TestRomanNumerals(t *testing.T) {
	actual := ConvertToRoman(1)
	expected := "I"

	if actual != expected {
		t.Errorf("actual %q, expected %q", actual, expected)
	}
}
