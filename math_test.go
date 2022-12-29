package math

import "testing"

func TestDivInt(t *testing.T) {
	want := 10 / 2
	got := Div[int](10, 2)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestDivFloat64(t *testing.T) {
	want := float64(10.1) / float64(2.2)
	got := Div[float64](10.1, 2.2)

	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
