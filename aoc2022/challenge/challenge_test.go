package challenge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChallenge(t *testing.T) {
	t.Run("Part Constants", func(t *testing.T) {
		assert.NotEqual(t, Part1, Part2)
	})
}
