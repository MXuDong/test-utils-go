package patch

import "testing"

func a() int {
	return 1
}

func b() int {
	return 2
}

func TestCore(t *testing.T) {
	Cover(a, b)
	if a() != b() {
		t.Errorf("error, not swarp")
	}
}
