package thing

import (
	"bytes"
	"reflect"
	"testing"
)

func TestThing_FindCreatureWithMaxAge(t *testing.T) {
	blade := Customer{4, "Blade", 600}
	frost := Customer{5, "Blood Sucker", 900}
	var creature any
	creature = blade
	var anotherCreature any
	anotherCreature = frost
	got := FindCreatureMaxAge(creature, anotherCreature)
	want := frost
	if reflect.DeepEqual(got, want) != true {
		t.Fatalf("получили %v, ожидалось %v", got, want)
	}
}

func TestThing_PassOnlyStrings(t *testing.T) {
	venom := &Customer{1, "Eddie Brock", 1200}
	gowron, duras, martok := "gowron", "duras", "martok"
	writer := bytes.NewBuffer([]byte{})
	got := PassOnlyStrings(writer, gowron, duras, martok, venom)
	want := []string{"gowron", "duras", "martok"}
	if reflect.DeepEqual(got, want) != true {
		t.Fatalf("получили %v, ожидалось %v", got, want)
	}
}
