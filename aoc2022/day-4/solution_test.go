package day4

import (
	_ "embed"
	"fmt"
	"testing"

	"colin-valentini.com/advent-of-code/aoc2022/challenge"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Note: input is generated from the Advent of Code site and is user specific.
// See https://adventofcode.com/2022/day/4/input.

//go:embed input.txt
var input string

//go:embed example.txt
var example string

func TestSolver(t *testing.T) {
	t.Run("Example Part 1", func(t *testing.T) {
		solution, err := NewSolver(example, challenge.Part1).Solve()
		require.NoError(t, err)
		assert.Equal(t, 2, solution)
	})

	t.Run("Example Part 2", func(t *testing.T) {
		solution, err := NewSolver(example, challenge.Part2).Solve()
		require.NoError(t, err)
		assert.Equal(t, 4, solution)
	})

	t.Run("Part 1", func(t *testing.T) {
		solution, err := NewSolver(input, challenge.Part1).Solve()
		require.NoError(t, err)
		assert.Equal(t, 588, solution)
	})

	t.Run("Part 2", func(t *testing.T) {
		solution, err := NewSolver(input, challenge.Part2).Solve()
		require.NoError(t, err)
		assert.Equal(t, 911, solution)
	})
}

func TestPair(t *testing.T) {
	testCases := []struct {
		fs, fe, ss, se   int
		covers, overlaps bool
	}{
		{fs: 1, fe: 5, ss: 6, se: 9, covers: false, overlaps: false},
		{fs: 1, fe: 5, ss: 2, se: 4, covers: true, overlaps: true},
		{fs: 1, fe: 5, ss: 5, se: 5, covers: true, overlaps: true},
		{fs: 5, fe: 5, ss: 5, se: 5, covers: true, overlaps: true},
		{fs: 1, fe: 5, ss: 3, se: 4, covers: true, overlaps: true},
		{fs: 1, fe: 5, ss: 3, se: 9, covers: false, overlaps: true},
		{fs: 6, fe: 96, ss: 4, se: 4, covers: false, overlaps: false},
		{fs: 53, fe: 69, ss: 52, se: 69, covers: true, overlaps: true},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			pair := newPair(
				newAssignment(testCase.fs, testCase.fe),
				newAssignment(testCase.ss, testCase.se),
			)
			assert.Equal(t, testCase.covers, pair.hasCover())
			assert.Equal(t, testCase.overlaps, pair.hasOverlap())
		})
	}
}

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
