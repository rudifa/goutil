// Package: ftime provides tools for time related operations.
package ftime

import (
	"log"
	"os"
	"time"
)

// Stopwatch is a tool for timing operations.
type Stopwatch struct {
	startTime time.Time
	logger    *log.Logger
}

// Chrono creates a new Stopwatch and starts it.
func Chrono() *Stopwatch {
	return &Stopwatch{
		logger:    log.New(os.Stderr, "", 0),
		startTime: time.Now(),
	}
}

// Start (re)starts the Stopwatch.
func (sw *Stopwatch) Start(message string) {
	sw.startTime = time.Now()
	if message != "" {
		sw.logger.Println(message)
	}
}

// Lap returns the elapsed time since the Stopwatch was started.
func (sw *Stopwatch) Lap() time.Duration {
	return time.Since(sw.startTime)
}

// Log logs the message an the elapsed time since the Stopwatch was started.
func (sw *Stopwatch) Log(message string) {
	elapsed := time.Since(sw.startTime)
	sw.logger.Println(message, elapsed)
}
