package main

import (
	"math/rand"
	"slices"
	"testing"
)

var testCases = []struct {
	name     string
	a, b     []int
	expected []int
}{
	{
		name:     "есть пересечение",
		a:        []int{1, 2, 3, 4},
		b:        []int{3, 4, 5, 6},
		expected: []int{3, 4},
	},
	{
		name:     "нет пересечения",
		a:        []int{1, 2, 3, 4},
		b:        []int{5, 6, 7, 8},
		expected: []int{},
	},
	{
		name:     "полное пересечение",
		a:        []int{1, 2, 3, 4},
		b:        []int{4, 3, 2, 1},
		expected: []int{1, 2, 3, 4},
	},
	{
		name:     "есть пересечение. несортированное",
		a:        []int{1, 3, 2, 4},
		b:        []int{5, 4, 3, 6},
		expected: []int{3, 4},
	},
}

func assertEqual(t *testing.T, got, want []int) {
	slices.Sort(got)
	slices.Sort(want)
	if !slices.Equal(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestIntersectWithMap(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := intersectWithMap(tc.a, tc.b)
			assertEqual(t, result, tc.expected)
		})
	}
}

func TestIntersectWithSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := intersectWithSort(tc.a, tc.b)
			assertEqual(t, result, tc.expected)
		})
	}
}

func BenchmarkIntersectWithMap(b *testing.B) {
	s1 := generateRandomSlice(100000, 0, 60000)
	s2 := generateRandomSlice(100000, 30000, 100000)
	for i := 0; i < b.N; i++ {

		intersectWithMap(s1, s2)
	}
}

func BenchmarkIntersectWithSort(b *testing.B) {
	s1 := generateRandomSlice(100000, 0, 60000)
	s2 := generateRandomSlice(100000, 30000, 100000)
	for i := 0; i < b.N; i++ {
		intersectWithSort(s1, s2)
	}
}

func generateRandomSlice(size, min, max int) []int {
	slice := make([]int, size)

	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(max-min+1) + min
	}

	return slice
}
