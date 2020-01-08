package main

import "testing"

type test struct {
	data   []int
	answer int
}

func TestAverage(t *testing.T) {
	tests := []test{
		test{[]int{87, 53, 87, 99, 13}, 67},
		test{[]int{85, 29, 93, 14, 48}, 53},
		test{[]int{31, 51, 34, 66, 58}, 48},
		test{[]int{50, 38, 62, 72, 81}, 60},
	}

	for _, v := range tests {
		x := average(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}