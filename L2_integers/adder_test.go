package integers

import "testing"

func TestAdder(t *testing.T) {
	got := Add(2, 3)
	expected := 5

	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}
