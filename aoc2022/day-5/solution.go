package day5

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"colin-valentini.com/advent-of-code/aoc2022/challenge"
)

// Solver solves the Day 4 challenge of Advent of Code 2022.
type Solver struct {
	input string
	part  challenge.Part
	crane *crane
	moves []instruction
}

// NewSolver returns a new instance of Solver.
func NewSolver(input string, part challenge.Part) *Solver {
	return &Solver{input: input, part: part}
}

func (s *Solver) Solve() (string, error) {
	err := s.parseInput()
	if err != nil {
		return "", fmt.Errorf("parsing input: %w", err)
	}
	for i, m := range s.moves {
		if err := s.crane.move(m); err != nil {
			return "", fmt.Errorf("applying move %d", i)
		}
	}
	var builder strings.Builder
	for i := 0; i < s.crane.numStacks(); i++ {
		builder.WriteRune(rune(s.crane.top(i)))
	}
	return builder.String(), nil
}

func (s *Solver) parseInput() error {
	lines := strings.Split(s.input, "\n")
	var moves []instruction
	var stacks []*stack
	for i, line := range lines {
		isStackPositions, err := regexp.MatchString(`^(\s*\d*)*\s*$`, line)
		if err != nil {
			return fmt.Errorf("regex match check for stack indexes: %w", err)
		}
		if line != "" && isStackPositions {
			stackPositions := s.parseStackPositions(line)
			stacks = s.parseStacks(lines[0:i], stackPositions)
		}
		if line == "" {
			moves = s.parseMoves(lines[i+1:])
			break
		}
	}
	s.crane = newCrane(stacks)
	s.moves = moves
	return nil
}

func (s *Solver) parseStackPositions(line string) []int {
	positions := []int{}
	for i, char := range line {
		// Assumes we won't have more than 9 stacks :)
		// And that the stacks are numbered in order ;)
		if '1' <= char && char <= '9' {
			positions = append(positions, i)
		}
	}
	return positions
}

func (s *Solver) parseMoves(lines []string) []instruction {
	moves := make([]instruction, len(lines))
	pattern := `move (\d+) from (\d+) to (\d+)`
	re := regexp.MustCompile(pattern)
	for i, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		if len(matches) == 0 {
			// Return error instead
			panic(fmt.Sprintf("unexpected move line %q", line))
		}
		match := matches[0]
		num := mustParseInt(match[1])
		from := mustParseInt(match[2]) - 1
		to := mustParseInt(match[3]) - 1
		moves[i] = newInstruction(num, from, to)
	}
	return moves
}

func (s *Solver) parseStacks(lines []string, stackPositions []int) []*stack {
	stacks := make([]*stack, len(stackPositions))
	// Iterate backwards up the lines so that the stack is
	// correctly populated bottoms up.
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		for j, pos := range stackPositions {
			char := line[pos]
			if 'A' <= char && char <= 'Z' {
				if stacks[j] == nil {
					stacks[j] = newStack([]crate{})
				}
				stacks[j].push(crate(char))
			}
		}
	}
	return stacks
}

func mustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("could not parse %q to int", s))
	}
	return i
}
