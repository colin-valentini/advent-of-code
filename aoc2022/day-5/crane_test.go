package day5

import (
	"testing"

	"colin-valentini.com/advent-of-code/aoc2022/challenge"
	"github.com/stretchr/testify/require"
)

func TestCrane(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		part := challenge.Part1
		//     [D]
		// [N] [C]
		// [Z] [M] [P]
		//  1   2   3
		crane := newExampleCrane(t)
		require.Equal(t, 3, crane.numStacks())

		// Check that the top of each crate matches expected.
		requireTops(t, crane, []crate{'N', 'D', 'P'})

		// [C]
		// [D]
		// [N]
		// [Z] [M] [P]
		//  1   2   3
		// Move two crates from the second stack to the first.
		require.NoError(t, crane.apply(newMove(2, 1, 0), part))
		requireTops(t, crane, []crate{'C', 'M', 'P'})

		//         [Z]
		//         [N]
		//         [D]
		//         [C]
		//     [M] [P]
		//  1   2   3
		// Move four crates from the first to the third stack.
		require.NoError(t, crane.apply(newMove(4, 0, 2), part))
		requireTops(t, crane, []crate{nilCrate, 'M', 'Z'})

		// Fail to move a crate from the first (empty) to the second stacks.
		err := crane.apply(newMove(1, 0, 1), part)
		require.ErrorIs(t, errInvalidMove, err)
	})

	t.Run("Part 2", func(t *testing.T) {
		part := challenge.Part2
		//     [D]
		// [N] [C]
		// [Z] [M] [P]
		//  1   2   3
		crane := newExampleCrane(t)
		require.Equal(t, 3, crane.numStacks())

		// Check that the top of each crate matches expected.
		requireTops(t, crane, []crate{'N', 'D', 'P'})

		// [D]
		// [C]
		// [N]
		// [Z] [M] [P]
		//  1   2   3
		// Move two crates from the second stack to the first.
		require.NoError(t, crane.apply(newMove(2, 1, 0), part))
		requireTops(t, crane, []crate{'D', 'M', 'P'})

		//         [D]
		//         [C]
		//         [N]
		//         [Z]
		//     [M] [P]
		//  1   2   3
		// Move four crates from the first to the third stack.
		require.NoError(t, crane.apply(newMove(4, 0, 2), part))
		requireTops(t, crane, []crate{nilCrate, 'M', 'D'})

		// Fail to move a crate from the first (empty) to the second stacks.
		err := crane.apply(newMove(1, 0, 1), part)
		require.ErrorIs(t, errInvalidMove, err)
	})
}

func requireTops(t *testing.T, crane *crane, tops []crate) {
	t.Helper()
	for i, topCrate := range tops {
		require.Equal(t, topCrate, crane.top(i))
	}
}

func newExampleCrane(t *testing.T) *crane {
	t.Helper()
	return newCrane([]*stack{
		newStack([]crate{'Z', 'N'}),
		newStack([]crate{'M', 'C', 'D'}),
		newStack([]crate{'P'}),
	})
}
