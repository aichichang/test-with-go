package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	got := 4

	if sum != got {
		// %d prints an integer, %q prints a sting
		t.Errorf("want '%d' but got '%d'", got, sum)
	}

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)

	// Output: 6
}
