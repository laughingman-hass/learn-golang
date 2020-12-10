package numeral

import "testing"

func TestRomanNumerals(t *testing.T) {
	t.Run("1 gets converted t I", func(t *testing.T) {
		actual := ConvertToRoman(1)
		expected := "I"

		if actual != expected {
			t.Errorf("actual %q, expected %q", actual, expected)
		}
	})

	t.Run("2 gets converted to II", func(t *testing.T) {
		actual := ConvertToRoman(2)
		expected := "II"

		if actual != expected {
			t.Errorf("actual %q, expected %q", actual, expected)
		}
	})
}
