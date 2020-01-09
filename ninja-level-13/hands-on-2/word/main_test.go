package word

import (
	"fmt"
	"testing"

	"github.com/cewitte/hello-go/ninja-level-13/hands-on-2/quote"
)

func TestCount(t *testing.T) {
	count := Count(quote.SunAlso)

	if count != 1349 {
		t.Error("Expected", 1349, "Got", count)
	}
}

func ExampleCount() {
	fmt.Println(Count(quote.SunAlso))
	// Output: 1349
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
