package sync_counter

import (
	"sync"
	"testing"
)

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

	t.Run("it runs safely concurrently", func(t *testing.T) {
		expectedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(expectedCount)

		for i := 0; i < expectedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, expectedCount)
	})
}

func assertCounter(t *testing.T, actual Counter, expected int) {
	t.Helper()
	if actual.Value() != expected {
		t.Errorf("actual %d, expected %d", actual.Value(), expected)
	}
}
