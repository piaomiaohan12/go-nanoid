package gonanoid

import (
	"testing"
)

// Test the distribution so we are sure, that the collisions won't happen frequently
// Test inspired by AI's javscript nanoid implementation
func TestGenerate(t *testing.T) {
	COUNTER := make(map[byte]int)
	ALPHABET := "abcdefghijklmnopqrstuvwxyz"
	COUNT := 100 * 1000
	SIZE := 5

	Alphabet(ALPHABET)
	Size(SIZE)
	for i := 0; i < COUNT; i++ {
		id := Generate()
		for u := 0; u < len(id); u++ {
			COUNTER[id[u]]++
		}
	}

	for char, count := range COUNTER {
		distribution := (float64(count) * float64(len(ALPHABET))) / float64((COUNT * SIZE))
		if !isInRange(distribution, 0.95, 1.05){
			t.Errorf("distribution error, char %v has %v distribution", char, distribution)
		}
	}
}

// Test if setting the size of ID works
func TestSetSize(t *testing.T) {
	var sizes = []int{4, 10, 20, 22, 30, 40, 60}
	for i := range sizes{
		Size(i)
		id := Generate()
		if len(id) != i {
			t.Errorf("Nanoid generated with false size: %d, except: %d", len(id), i)
		}
	}
}

// Helping function to find if number is in given range
func isInRange(num float64, from float64, to float64) bool {
	return num > from && num < to
}

func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Generate()
	}
}