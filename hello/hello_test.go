package hello

import "testing"

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
	}
}

func TestGreet(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions. tomxu.top is awesome" {
		t.Errorf("Greet() = %s; Expected Hello GitHub Actions. tomxu.top is awesome", result)
	}

}
