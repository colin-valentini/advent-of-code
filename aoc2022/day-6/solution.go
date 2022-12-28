package day6

import "colin-valentini.com/advent-of-code/aoc2022/challenge"

type Solver struct {
	input string
	part  challenge.Part
}

func NewSolver(input string, part challenge.Part) *Solver {
	return &Solver{input: input, part: part}
}

func (s *Solver) Solve() int {
	// TODO: Do something smarter. Might want to jump ahead
	// sufficiently far when we know there is going to be a
	// duplicate character in the next window.
	bytes := []byte(s.input)
	winSize := s.windowSize()
	for i := winSize; i <= len(bytes); i++ {
		if newWindow(bytes[i-winSize : i]).isMarker() {
			return i
		}
	}
	return -1
}

func (s *Solver) windowSize() int {
	if s.part == challenge.Part1 {
		return 4
	}
	return 14
}

type window struct {
	bytes []byte
}

func newWindow(bytes []byte) *window {
	return &window{bytes: bytes}
}

func (w *window) isMarker() bool {
	set := make(map[byte]nothing)
	for _, char := range w.bytes {
		if _, ok := set[char]; ok {
			return false
		}
		set[char] = nothing{}
	}
	return true
}

type nothing struct{}
