package day2

import (
	_ "embed"
	"testing"

	"colin-valentini.com/advent-of-code/aoc2022/challenge"
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
		solution, err := NewSolver(example, challenge.Part1).Solve()
		require.NoError(t, err)
		assert.Equal(t, 15, solution)
	})

	t.Run("Example Part 2", func(t *testing.T) {
		solution, err := NewSolver(example, challenge.Part2).Solve()
		require.NoError(t, err)
		assert.Equal(t, 12, solution)
	})

	t.Run("Part 1", func(t *testing.T) {
		solution, err := NewSolver(input, challenge.Part1).Solve()
		require.NoError(t, err)
		assert.Equal(t, 14_531, solution)
	})

	t.Run("Part 2", func(t *testing.T) {
		solution, err := NewSolver(input, challenge.Part2).Solve()
		require.NoError(t, err)
		assert.Equal(t, 11_258, solution)
	})
}

// BenchmarkSolver runs benchmarks for the Day 2 solution.
func BenchmarkSolver(b *testing.B) {
	b.Run("Part 1", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = NewSolver(input, challenge.Part1).Solve()
		}
	})
	b.Run("Part 2", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = NewSolver(input, challenge.Part2).Solve()
		}
	})
}
