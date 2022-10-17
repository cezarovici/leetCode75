package main

import (
	"testing"
)

var benchCases = []struct {
	description string
	input       []int
}{
	{"example1", []int{1, 2, 3, 4}},
	{"example2", []int{1, 1, 1, 1, 1}},
	{"example3", []int{3, 1, 2, 10, 1}},
}

func BenchmarkRunningSumDoubleFor(b *testing.B) {
	for _, bc := range benchCases {
		b.Run(bc.description, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				runningSumDoubleFor(bc.input)
			}
		})
	}
}

func BenchmarkDecreasingTheLatest(b *testing.B) {
	for _, bc := range benchCases {
		b.Run(bc.description, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				runningSumDecreasingTheLast(bc.input)
			}
		})
	}
}
