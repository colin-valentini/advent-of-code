package day7

import (
	_ "embed"
	"testing"

	"colin-valentini.com/advent-of-code/aoc2022/challenge"
	"github.com/stretchr/testify/assert"
)

// Note: input is generated from the Advent of Code site and is user specific.
// See https://adventofcode.com/2022/day/7/input.

//go:embed input.txt
var input string

//go:embed example.txt
var example string

func TestSolver(t *testing.T) {
	t.Run("Example: Part 1", func(t *testing.T) {
		got := NewSolver(example, challenge.Part1).Solve()
		assert.Equal(t, 95_437, got)
	})

	t.Run("Example: Part 2", func(t *testing.T) {
		got := NewSolver(example, challenge.Part2).Solve()
		assert.Equal(t, 24_933_642, got)
	})

	t.Run("Input: Part 1", func(t *testing.T) {
		got := NewSolver(input, challenge.Part1).Solve()
		assert.Equal(t, 1_348_005, got)
	})

	t.Run("Input: Part 2", func(t *testing.T) {
		got := NewSolver(input, challenge.Part2).Solve()
		assert.Equal(t, 12_785_886, got)
	})
}
