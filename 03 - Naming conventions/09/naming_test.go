package naming

import "testing"

func TestColor(t *testing.T) {
	arg := "blue"
	want := "#0000FF"
	got := Color("blue")
	if got != want {
		t.Errorf("Color(%q) = %q; wanted %q", arg, got, want)
	}
}
