package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunningSum(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		want        []int
	}{
		{"single", []int{1}, []int{1}},
		{"negative", []int{-1, 1}, []int{-1, 0}},
		{"example1", []int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{"example2 all equals", []int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{"example3", []int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, runningSumDecreasingTheLast(tc.input), runningSumDoubleFor(tc.input), "running sum")
		})
	}

}
