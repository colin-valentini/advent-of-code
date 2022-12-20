package day1

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Note: input is generated from the Advent of Code site and is user specific.
// See https://adventofcode.com/2022/day/1/input.

//go:embed input.txt
var input string

//go:embed example.txt
var example string

// TestSolver tests that the solution to the Day 1 challenge is correct.
func TestSolver(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		solution, err := NewSolver(example).Solve()
		require.NoError(t, err)
		assert.Equal(t, 24_000, solution)
	})

	t.Run("Input", func(t *testing.T) {
		solution, err := NewSolver(input).Solve()
		require.NoError(t, err)
		assert.Equal(t, 69_626, solution)
	})
}

// BenchmarkSolver runs benchmarks for the Day 1 solution.
func BenchmarkSolver(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = NewSolver(input).Solve()
	}
}
