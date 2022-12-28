package day6

import "colin-valentini.com/advent-of-code/aoc2022/challenge"

type Solver struct {
	input string
	part  challenge.Part
}

func newSolver(input string, part challenge.Part) *Solver {
	return &Solver{input: input, part: part}
}

func (s *Solver) Solve() int {
	bytes := []byte(s.input)
	for i := 4; i <= len(bytes); i++ {
		if newWindow(bytes[i-4 : i]).isMarker() {
			return i
		}
	}
	return -1
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
