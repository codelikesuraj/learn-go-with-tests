package integers

import "fmt"
import "testing"

func TestAdder(t *testing.T) {
	got := Add(2, 4)
	expected := 6

	if got != expected {
		t.Errorf("expected %d but got %d ", expected, got)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}