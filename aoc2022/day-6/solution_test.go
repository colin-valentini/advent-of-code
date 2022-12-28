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
		in   string
		want int
	}{
		{in: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want: 7},
		{in: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 5},
		{in: "nppdvjthqldpwncqszvftbrmjlhg", want: 6},
		{in: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 10},
		{in: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 11},
	}
	for i, example := range examples {
		t.Run(fmt.Sprintf("Example %d", i), func(t *testing.T) {
			got := newSolver(example.in, challenge.Part1).Solve()
			assert.Equal(t, example.want, got)
		})
	}

	t.Run("Input", func(t *testing.T) {
		got := newSolver(input, challenge.Part1).Solve()
		assert.Equal(t, 1848, got)
	})
}
