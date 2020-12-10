package sync_counter

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("actual %d, expected %d", counter.Value(), 3)
		}
	})
}

func assertCounter(t *testing.T, actual Counter, expected int) {
	t.Helper()
	if actual.Value() != expected {
		t.Errorf("actual %d, expected %d", actual.Value(), expected)
	}
}
