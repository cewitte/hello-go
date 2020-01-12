package mymath

import (
	"fmt"
	"testing"
)

type testTable struct {
	data   []int
	result float64
}

func TestCenteredAverage(t *testing.T) {
	table := []testTable{
		{[]int{1, 4, 6, 8, 100}, 6},
		{[]int{0, 8, 10, 1000}, 9},
		{[]int{9000, 4, 10, 8, 6, 12}, 9},
		{[]int{123, 744, 140, 200}, 170},
	}

	for _, v := range table {
		if CenteredAvg(v.data) != v.result {
			t.Error("Expected:", v.result, "Got:", CenteredAvg(v.data))
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{123, 744, 140, 200}))
	// Output: 170
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{9000, 4, 10, 8, 6, 12})
	}
}
