package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRockPaperScissors(t *testing.T) {
	t.Run("Against", func(t *testing.T) {
		assert.Equal(t, win, rock.against(scissors))
		assert.Equal(t, loss, rock.against(paper))
		assert.Equal(t, draw, rock.against(rock))
		assert.Equal(t, loss, paper.against(scissors))
		assert.Equal(t, draw, paper.against(paper))
		assert.Equal(t, win, paper.against(rock))
		assert.Equal(t, draw, scissors.against(scissors))
		assert.Equal(t, win, scissors.against(paper))
		assert.Equal(t, loss, scissors.against(rock))
	})

	t.Run("Value", func(t *testing.T) {
		assert.Equal(t, 1, rock.value())
		assert.Equal(t, 2, paper.value())
		assert.Equal(t, 3, scissors.value())
	})

	t.Run("Wins Against", func(t *testing.T) {
		assert.Equal(t, scissors, rock.winsAgainst())
		assert.Equal(t, rock, paper.winsAgainst())
		assert.Equal(t, paper, scissors.winsAgainst())
	})

	t.Run("Loses Against", func(t *testing.T) {
		assert.Equal(t, paper, rock.losesAgainst())
		assert.Equal(t, scissors, paper.losesAgainst())
		assert.Equal(t, rock, scissors.losesAgainst())
	})

	t.Run("Draws Against", func(t *testing.T) {
		assert.Equal(t, rock, rock.drawsAgainst())
		assert.Equal(t, paper, paper.drawsAgainst())
		assert.Equal(t, scissors, scissors.drawsAgainst())
	})
}
