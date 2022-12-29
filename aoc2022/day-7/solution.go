package day7

import "colin-valentini.com/advent-of-code/aoc2022/challenge"

type Solver struct {
	input string
	part  challenge.Part
}

const (
	maxDiscSpace      = 70_000_000
	minSpaceForUpdate = 30_000_000
)

func NewSolver(input string, part challenge.Part) *Solver {
	return &Solver{input: input, part: part}
}

func (s *Solver) Solve() int {
	root := newParser(s.input).parse()
	if s.part == challenge.Part1 {
		return s.partOne(root)
	}
	return s.partTwo(root)
}

func (s *Solver) partOne(root *directory) int {
	total := 0
	for _, dir := range root.children {
		size := dir.size()
		// Problem calls for weird constraint of only accumulating
		// total size if each directory is under a certain size,
		// and also double-counts (stupid).
		if size <= 100_000 {
			total += size
		}
		total += s.partOne(dir)
	}
	return total
}

func (s *Solver) partTwo(root *directory) int {
	sizes := []int{}
	totalSize := s.collectSizes(root, &sizes) + root.size()
	sizes = append(sizes, totalSize)
	return s.minSize(s.filterSizes(sizes, totalSize), totalSize)
}

func (s *Solver) collectSizes(root *directory, sizes *[]int) int {
	size := 0
	for _, dir := range root.children {
		*sizes = append(*sizes, dir.size()+s.collectSizes(dir, sizes))
	}
	return size
}

func (s *Solver) filterSizes(sizes []int, totalSize int) []int {
	unusedSpace := maxDiscSpace - totalSize
	filtered := make([]int, 0, len(sizes))
	for _, size := range sizes {
		if unusedSpace+size >= minSpaceForUpdate {
			filtered = append(filtered, size)
		}
	}
	return filtered
}

func (s *Solver) minSize(sizes []int, max int) int {
	min := max
	for _, size := range sizes {
		if size < min {
			min = size
		}
	}
	return min
}
