package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	got := newParser(example).parse()
	want := newTestDirectory(t)
	assert.Equal(t, want, got)
}
