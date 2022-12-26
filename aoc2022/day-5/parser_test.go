package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	parser := newParser()
	crane, moves, err := parser.parse(example)
	require.NoError(t, err)
	assert.Equal(t, newExampleCrane(t), crane)
	assert.Equal(t, []*move{
		newMove(1, 1, 0),
		newMove(3, 0, 2),
		newMove(2, 1, 0),
		newMove(1, 0, 1),
	}, moves)
}
