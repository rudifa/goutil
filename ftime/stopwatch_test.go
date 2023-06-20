package ftime_test

import (
	"testing"
	"time"

	"github.com/rudifa/goutil/ftime"
)

// TestStopwatch tests the Stopwatch type.
func TestStopwatch(t *testing.T) {

	// Create a new Stopwatch.
	sw := ftime.Chrono()

	// Sleep for a while.
	time.Sleep(1 * time.Second)

	// Log the elapsed time.
	sw.Log("Elapsed time:")

	// Sleep for a while.
	time.Sleep(1 * time.Second)

	// Log the elapsed time.
	sw.Log("Elapsed time:")

	// Test the Lap() time against expected.
	lap := sw.Lap()

	if lap < 2*time.Second {
		t.Errorf("Lap() returned %v, expected at least 2 seconds.", lap)
	}

}
