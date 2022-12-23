package day5

import (
	"errors"
)

var errInvalidMove = errors.New("invalid instruction")

const nilCrate = crate(0)

type crane struct {
	stacks []*stack
}

func newCrane(stacks []*stack) *crane {
	return &crane{stacks: stacks}
}

func (c *crane) move(inst instruction) error {
	for i := 0; i < inst.num; i++ {
		crate, ok := c.stacks[inst.from].pop()
		if !ok {
			return errInvalidMove
		}
		c.stacks[inst.to].push(crate)
	}
	return nil
}

func (c *crane) numStacks() int {
	return len(c.stacks)
}

func (c *crane) top(i int) crate {
	if len(c.stacks) <= i {
		return nilCrate
	}
	crate, _ := c.stacks[i].peek()
	return crate
}

// stack represents a stack of crates.
// The last element in the stack is the top.
type stack struct {
	crates []crate
}

func newStack(crates []crate) *stack {
	return &stack{crates: crates}
}

func (s *stack) push(c crate) {
	s.crates = append(s.crates, c)
}

func (s *stack) pop() (crate, bool) {
	if len(s.crates) == 0 {
		return nilCrate, false
	}
	i := len(s.crates) - 1
	c := s.crates[i]
	s.crates = s.crates[:i]
	return c, true
}

func (s *stack) peek() (crate, bool) {
	if len(s.crates) == 0 {
		return nilCrate, false
	}
	return s.crates[len(s.crates)-1], true
}

type crate rune

type instruction struct {
	num  int
	from int
	to   int
}

func newInstruction(num, from, to int) instruction {
	return instruction{num: num, from: from, to: to}
}
