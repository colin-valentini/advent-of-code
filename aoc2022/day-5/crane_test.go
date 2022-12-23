package day5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCrane(t *testing.T) {
	// [D]
	// [N] [C]
	// [Z] [M] [P]
	//  1   2   3
	crane := newCrane([]*stack{
		newStack([]crate{'Z', 'N', 'D'}),
		newStack([]crate{'M', 'C'}),
		newStack([]crate{'P'}),
	})
	require.Equal(t, 3, crane.numStacks())

	// Check that the top of each crate matches expected.
	requireTops(t, crane, []crate{'D', 'C', 'P'})

	//     [D]
	// [N] [C]
	// [Z] [M] [P]
	//  1   2   3
	// Move one crate from the first stack to the second.
	require.NoError(t, crane.move(instruction{num: 1, from: 0, to: 1}))
	requireTops(t, crane, []crate{'N', 'D', 'P'})

	// [C]
	// [D]
	// [N]
	// [Z] [M] [P]
	//  1   2   3
	// Move two crates from the second stack to the first.
	require.NoError(t, crane.move(instruction{num: 2, from: 1, to: 0}))
	requireTops(t, crane, []crate{'C', 'M', 'P'})

	//         [Z]
	//         [N]
	//         [D]
	//         [C]
	//     [M] [P]
	//  1   2   3
	// Move four crates from the first to the third stack.
	require.NoError(t, crane.move(instruction{num: 4, from: 0, to: 2}))
	requireTops(t, crane, []crate{nilCrate, 'M', 'Z'})

	// Fail to move a crate from the first (empty) to the second stacks.
	err := crane.move(instruction{num: 1, from: 0, to: 1})
	require.ErrorIs(t, errInvalidMove, err)
}

func requireTops(t *testing.T, crane *crane, tops []crate) {
	t.Helper()
	for i, topCrate := range tops {
		require.Equal(t, topCrate, crane.top(i))
	}
}
