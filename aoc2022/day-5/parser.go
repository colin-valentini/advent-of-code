package day5

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type parser struct {
	cp *craneParser
	mp *moveParser
}

func newParser() *parser {
	return &parser{cp: newCraneParser(), mp: newMoveParser()}
}

func (p *parser) parse(str string) (*crane, []*move, error) {
	lines := strings.Split(str, "\n")
	var mid int
	for i, line := range lines {
		if line == "" {
			mid = i
			break
		}
	}
	crane, err := p.parseCrane(lines[:mid])
	if err != nil {
		return nil, nil, fmt.Errorf("parsing crane: %w", err)
	}
	moves, err := p.parseMoves(lines[mid+1:])
	if err != nil {
		return nil, nil, fmt.Errorf("parsing moves: %w", err)
	}
	return crane, moves, nil
}

func (p *parser) parseCrane(lines []string) (*crane, error) {
	crane, err := p.cp.parse(lines)
	if err != nil {
		return nil, fmt.Errorf("parsing crane: %w", err)
	}
	return crane, nil
}

type craneParser struct {
	re *regexp.Regexp
}

func newCraneParser() *craneParser {
	re := regexp.MustCompile(`\d`)
	return &craneParser{re: re}
}

func (p *craneParser) parse(lines []string) (*crane, error) {
	if len(lines) < 1 {
		return nil, errors.New("not enough lines to parse crane")
	}
	positions, err := p.parseStackPositions(lines[len(lines)-1])
	if err != nil {
		return nil, fmt.Errorf("parsing stack positions: %w", err)
	}
	stacks := make([]*stack, len(positions))
	for i := range positions {
		stacks[i] = newStack([]crate{})
	}
	for i := len(lines) - 2; i >= 0; i-- {
		line := lines[i]
		for i, pos := range positions {
			char := line[pos]
			if 'A' <= char && char <= 'Z' {
				stacks[i].push(crate(char))
			}
		}
	}
	return newCrane(stacks), nil
}

func (p *craneParser) parseStackPositions(line string) ([]int, error) {
	positions := make([]int, 0, 9) // Problem can only have 9 stacks
	matches := p.re.FindAllStringSubmatchIndex(line, -1)
	for _, submatch := range matches {
		if len(submatch) != 2 {
			return nil, fmt.Errorf("unexpected result %v from regex", submatch)
		}
		index := submatch[0]
		positions = append(positions, index)
	}
	return positions, nil
}

func (p *parser) parseMoves(lines []string) ([]*move, error) {
	moves, err := p.mp.parse(lines)
	if err != nil {
		return nil, fmt.Errorf("parsing moves: %w", err)
	}
	return moves, nil
}

type moveParser struct {
	re *regexp.Regexp
}

func newMoveParser() *moveParser {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	return &moveParser{re: re}
}

func (p *moveParser) parse(lines []string) ([]*move, error) {
	moves := make([]*move, len(lines))
	for i, line := range lines {
		move, err := p.parseMove(line)
		if err != nil {
			return nil, fmt.Errorf("parsing move: %w", err)
		}
		moves[i] = move
	}
	return moves, nil
}

func (p *moveParser) parseMove(line string) (*move, error) {
	matches := p.re.FindAllStringSubmatch(line, -1)
	if len(matches) != 1 {
		return nil, fmt.Errorf("could not parse line %q as a move", line)
	}
	submatches := matches[0]
	if len(submatches) != 4 {
		return nil, fmt.Errorf("line %q is missing move information", line)
	}
	num, err := strconv.Atoi(submatches[1])
	if err != nil {
		return nil, fmt.Errorf("invalid number of crates to move in %q", line)
	}
	from, err := strconv.Atoi(submatches[2])
	if err != nil {
		return nil, fmt.Errorf("invalid crate to move from %q", line)
	}
	to, err := strconv.Atoi(submatches[3])
	if err != nil {
		return nil, fmt.Errorf("invalid crate to move to %q", line)
	}
	return newMove(num, from-1, to-1), nil
}
