package iteration

import "fmt"
import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("x", 8))
	// Output: xxxxxxxx
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 30)
	}
}