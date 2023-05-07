package math

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	got := Min(4, 6)
	want := 4

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestMin2(t *testing.T) {
	got := Min(6, 4)
	want := 4

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSubtract(t *testing.T) {
	cases := []struct {
		a    int
		b    int
		want int
	}{
		// {1, 2, -1}, // test coverage
		{3, 1, 2},
		{3, 3, 0},
	}

	for _, c := range cases {
		got := Subtract(c.a, c.b)
		if got != c.want {
			t.Errorf("Min(%d, %d) == %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Min(4, 6)
	}
}

// https://pkg.go.dev/testing#hdr-Examples
func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}
