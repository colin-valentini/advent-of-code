package day5

import (
	_ "embed"
	"testing"

	"colin-valentini.com/advent-of-code/aoc2022/challenge"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Note: input is generated from the Advent of Code site and is user specific.
// See https://adventofcode.com/2022/day/5/input.

//go:embed input.txt
var input string

//go:embed example.txt
var example string

func TestSolver(t *testing.T) {
	t.Run("Example Part 1", func(t *testing.T) {
		solution, err := NewSolver(example, challenge.Part1).Solve()
		require.NoError(t, err)
		assert.Equal(t, "CMZ", solution)
	})

	t.Run("Example Part 2", func(t *testing.T) {
		solution, err := NewSolver(example, challenge.Part2).Solve()
		require.NoError(t, err)
		assert.Equal(t, "MCD", solution)
	})

	t.Run("Input Part 1", func(t *testing.T) {
		solution, err := NewSolver(input, challenge.Part1).Solve()
		require.NoError(t, err)
		assert.Equal(t, "VWLCWGSDQ", solution)
	})

	t.Run("Input Part 2", func(t *testing.T) {
		solution, err := NewSolver(input, challenge.Part2).Solve()
		require.NoError(t, err)
		assert.Equal(t, "TCGLQSLPW", solution)
	})
}

func BenchmarkSolver(b *testing.B) {
	b.Run("Input", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = NewSolver(input, challenge.Part1).Solve()
		}
	})
}
