package day3

import (
	_ "embed"
	"testing"

	"colin-valentini.com/advent-of-code/aoc2022/challenge"
	"github.com/stretchr/testify/assert"
)

// Note: input is generated from the Advent of Code site and is user specific.
// See https://adventofcode.com/2022/day/3/input.

//go:embed input.txt
var input string

//go:embed example.txt
var example string

func TestSolver(t *testing.T) {
	t.Run("Example Part 1", func(t *testing.T) {
		solution := NewSolver(example, challenge.Part1).Solve()
		assert.Equal(t, 157, solution)
	})

	t.Run("Example Part 2", func(t *testing.T) {
		solution := NewSolver(example, challenge.Part2).Solve()
		assert.Equal(t, 70, solution)
	})

	t.Run("Part 1", func(t *testing.T) {
		solution := NewSolver(input, challenge.Part1).Solve()
		assert.Equal(t, 7_746, solution)
	})

	t.Run("Part 2", func(t *testing.T) {
		solution := NewSolver(input, challenge.Part2).Solve()
		assert.Equal(t, 2_604, solution)
	})
}

func BenchmarkSolver(b *testing.B) {
	b.Run("Part 1", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = NewSolver(input, challenge.Part1).Solve()
		}
	})
	b.Run("Part 2", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = NewSolver(input, challenge.Part2).Solve()
		}
	})
}
