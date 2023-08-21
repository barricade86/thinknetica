package person

import "testing"

func TestPerson_FindMaxAge(t *testing.T) {
	venom := &Customer{1, "Eddie Brock", 1200}
	carnage := &Employee{2, "Symbiont", "Kletus Cassidy", 500}
	toxin := &Employee{3, "Symbiont", "Patrick Mulligan", 900}

	got := FindMaxAge(venom, carnage, toxin)
	var want uint
	want = 1200
	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}
