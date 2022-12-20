package day2

import (
	"fmt"
	"strings"

	"colin-valentini.com/advent-of-code/aoc2022/challenge"
)

// Solver solves the Day 2 challenge of Advent of Code 2022.
type Solver struct {
	input string
	part  challenge.Part
}

// NewSolver returns a new instance of Solver.
func NewSolver(input string, part challenge.Part) *Solver {
	return &Solver{input: input, part: part}
}

func (s *Solver) Solve() (int, error) {
	tournament, err := s.parseInput()
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}
	return tournament.score(), nil
}

func (s *Solver) parseInput() (tournament, error) {
	roundsStr := strings.Split(s.input, "\n")
	rounds := make([]*round, len(roundsStr))
	for i, rstr := range roundsStr {
		playsStr := strings.Split(rstr, " ")
		if len(playsStr) < 2 || len(playsStr[0]) != 1 || len(playsStr[1]) != 1 {
			return nil, fmt.Errorf("invalid input %q on line %d", rstr, i)
		}
		oppShape := s.parseOpponent(playsStr[0])
		shape := s.parsePlay(playsStr[1], oppShape)
		rounds[i] = newRound(shape, oppShape)
	}
	return newTournament(rounds), nil
}

func (s *Solver) parseOpponent(str string) handShape {
	return handShape([]rune(str)[0] - 'A')
}

func (s *Solver) parsePlay(str string, oppShape handShape) handShape {
	char := []rune(str)[0]
	if s.part == challenge.Part1 {
		return handShape(char - 'X')
	}
	switch char {
	case 'X':
		return oppShape.winsAgainst()
	case 'Z':
		return oppShape.losesAgainst()
	default: // 'Y'
		return oppShape.drawsAgainst()
	}
}

type tournament []*round

func newTournament(rounds []*round) tournament {
	return tournament(rounds)
}

func (t tournament) score() int {
	s := 0
	for _, round := range t {
		s += round.score()
	}
	return s
}

type round struct {
	shape    handShape
	oppShape handShape
}

func newRound(shape, oppShape handShape) *round {
	return &round{shape: shape, oppShape: oppShape}
}

func (r *round) score() int {
	return r.shape.value() + r.shape.against(r.oppShape).value()
}
