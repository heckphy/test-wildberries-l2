package main

import (
	"testing"
	"time"
)

func TestSig(t *testing.T) {
	c := sig(2 * time.Second)

	select {
	case _, ok := <-c:
		if !ok {
			t.Log("Channel closed as expected")
		}
	case <-time.After(3 * time.Second):
		// Прошло больше времени, чем ожидалось
		t.Error("Timeout: Done channel not closed after 3 seconds")
	}
}

func TestOrChannel(t *testing.T) {
	start := time.Now()
	expected := 1 * time.Second

	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	end := time.Since(start)
	tolerance := 100 * time.Millisecond
	roundedEnd := end.Round(tolerance)

	if roundedEnd < expected-tolerance || roundedEnd > expected+tolerance {
		t.Errorf("time passed %v, expected %v +/- %v", end, expected, tolerance)
	}
}
