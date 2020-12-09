package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.co.uk"

	expected := fastURL
	actual := Racer(slowURL, fastURL)

	if actual != expected {
		t.Errorf("actual %q, expected %q", actual, expected)
	}
}
