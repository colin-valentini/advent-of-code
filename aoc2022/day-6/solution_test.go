package day6

import (
	_ "embed"
	"fmt"
	"testing"

	"colin-valentini.com/advent-of-code/aoc2022/challenge"
	"github.com/stretchr/testify/assert"
)

// Note: input is generated from the Advent of Code site and is user specific.
// See https://adventofcode.com/2022/day/6/input.

//go:embed input.txt
var input string

func TestSolver(t *testing.T) {
	examples := []struct {
		in    string
		part1 int
		part2 int
	}{
		{in: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", part1: 7, part2: 19},
		{in: "bvwbjplbgvbhsrlpgdmjqwftvncz", part1: 5, part2: 23},
		{in: "nppdvjthqldpwncqszvftbrmjlhg", part1: 6, part2: 23},
		{in: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", part1: 10, part2: 29},
		{in: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", part1: 11, part2: 26},
	}
	for i, example := range examples {
		t.Run(fmt.Sprintf("Example %d: Part 1", i), func(t *testing.T) {
			got := NewSolver(example.in, challenge.Part1).Solve()
			assert.Equal(t, example.part1, got)
		})
		t.Run(fmt.Sprintf("Example %d: Part 2", i), func(t *testing.T) {
			got := NewSolver(example.in, challenge.Part2).Solve()
			assert.Equal(t, example.part2, got)
		})
	}

	t.Run("Input: Part 1", func(t *testing.T) {
		got := NewSolver(input, challenge.Part1).Solve()
		assert.Equal(t, 1848, got)
	})

	t.Run("Input: Part 2", func(t *testing.T) {
		got := NewSolver(input, challenge.Part2).Solve()
		assert.Equal(t, 2308, got)
	})
}

func BenchmarkSolver(b *testing.B) {
	b.Run("Input", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = NewSolver(input, challenge.Part2).Solve()
		}
	})
}
