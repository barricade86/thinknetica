package list

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"
)

func TestSort_Ints(t *testing.T) {
	ints := []int{22, 543, 3, 5, 1, 7, 2}
	sort.Ints(ints)
	got := fmt.Sprintf("%v", ints)
	want := "[1 2 3 5 7 22 543]"
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}
}

func TestSort_Strings(t *testing.T) {
	tests := map[string]struct {
		input  string
		wanted string
	}{
		"first": {
			input:  "Rabbit Duck MegaDeth",
			wanted: "Duck MegaDeth Rabbit",
		},
		"second": {
			input:  "Romulan Klingon Kronos",
			wanted: "Klingon Kronos Romulan",
		},
		"third": {
			input:  "Sith Jedi Yoda",
			wanted: "Jedi Sith Yoda",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			inputSlice := strings.Split(test.input, " ")
			sort.Strings(inputSlice)
			got := strings.Join(inputSlice, " ")
			if got != test.wanted {
				t.Errorf("got %q, want %q", got, test.wanted)
			}
		})
	}
}

func sampleData() []int {
	rand.Seed(time.Now().UnixNano())
	var data []int
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Intn(1000))
	}

	return data
}

func sampleFloatData() []float64 {
	rand.Seed(time.Now().UnixNano())
	var data []float64
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Float64())
	}

	return data
}

func BenchmarkSortInt(t *testing.B) {
	for i := 0; i < t.N; i++ {
		data := sampleData()
		sort.Ints(data)
		_ = data
	}
}

func BenchmarkSortFloats(t *testing.B) {
	for i := 0; i < t.N; i++ {
		data := sampleFloatData()
		sort.Float64s(data)
		_ = data
	}
}
