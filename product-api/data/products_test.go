package data

import "testing"

func TestChecksValidations(t *testing.T) {
	p := &Product{
		Name:  "Hass",
		Price: 1.00,
		SKU:   "abs-def-ghi",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
