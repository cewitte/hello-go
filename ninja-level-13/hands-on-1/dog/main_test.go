package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	d := Years(10)

	if d != 70 {
		t.Error("Expected:", 70, "Got", d)
	}
}

func TestYearsTwo(t *testing.T) {
	d := YearsTwo(10)

	if d != 70 {
		t.Error("Expected:", 70, "Got", d)
	}
}

func ExamleYears() {
	fmt.Println(Years(10))
	// Output: 70
}

func ExamleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output: 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
