package main

import "testing"

func TestAverage(t *testing.T) {
	avg := average(3, 5)

	if avg != 4 {
		t.Error("Expected", 4, "Got", avg)
	}
}
