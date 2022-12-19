package day2

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Note: input is generated from the Advent of Code site
// and is user specific.
// See https://adventofcode.com/2022/day/2/input.

//go:embed input.txt
var input string

//go:embed example.txt
var example string

// TestSolver tests that the solution to the Day 2 challenge is correct.
func TestSolver(t *testing.T) {
	t.Run("Example Part 1", func(t *testing.T) {
		solution, err := NewSolver(example, part1).Solve()
		require.NoError(t, err)
		assert.Equal(t, 15, solution)
	})

	t.Run("Example Part 2", func(t *testing.T) {
		solution, err := NewSolver(example, part2).Solve()
		require.NoError(t, err)
		assert.Equal(t, 12, solution)
	})

	t.Run("Part 1", func(t *testing.T) {
		solution, err := NewSolver(input, part1).Solve()
		require.NoError(t, err)
		assert.Equal(t, 14_531, solution)
	})

	t.Run("Part 2", func(t *testing.T) {
		solution, err := NewSolver(input, part2).Solve()
		require.NoError(t, err)
		assert.Equal(t, 11_258, solution)
	})
}

// BenchmarkSolver runs benchmarks for the Day 2 solution.
func BenchmarkSolver(b *testing.B) {
	b.Run("Part 1", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = NewSolver(input, part1).Solve()
		}
	})
	b.Run("Part 2", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = NewSolver(input, part2).Solve()
		}
	})
}
