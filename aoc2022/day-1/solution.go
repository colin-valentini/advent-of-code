package day1

import (
	"fmt"
	"strconv"
	"strings"
)

// Solver solves the Day 1 challenge of Advent of Code 2022.
type Solver struct {
	input string
}

// NewSolver returns a new instance of Solver.
func NewSolver(input string) *Solver {
	return &Solver{input: input}
}

func (s *Solver) Solve() (int, error) {
	elfCalories, err := s.parseInput()
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}
	return elfCalories.maxCalories(), nil
}

func (s *Solver) parseInput() (*elfCalories, error) {
	elfStrs := strings.Split(s.input, "\n\n")
	elves := make([]*elf, 0, len(elfStrs))
	for i, elfStr := range elfStrs {
		elf, err := s.parseInputElf(elfStr)
		if err != nil {
			return nil, fmt.Errorf("parsing input elf %d: %w", i, err)
		}
		elves = append(elves, elf)
	}
	return newElfCalories(elves), nil
}

func (s *Solver) parseInputElf(elfStr string) (*elf, error) {
	strCals := strings.Split(elfStr, "\n")
	calories := make([]uint64, 0, len(strCals))
	for _, calStr := range strCals {
		cal, err := strconv.ParseUint(calStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parsing elf calorie %q: %w", calStr, err)
		}
		calories = append(calories, cal)
	}
	return newElf(calories), nil
}

type elfCalories struct {
	elves []*elf
}

func newElfCalories(elves []*elf) *elfCalories {
	return &elfCalories{elves: elves}
}

func (e *elfCalories) maxCalories() int {
	max := uint64(0)
	for _, elf := range e.elves {
		if elf.totalCalories() > max {
			max = elf.totalCalories()
		}
	}
	return int(max)
}

// String returns the string representation of the elfCalories type.
func (e *elfCalories) String() string {
	strs := make([]string, len(e.elves))
	for i, elf := range e.elves {
		strs[i] = elf.String()
	}
	return strings.Join(strs, "\n\n")
}

type elf struct {
	calories []uint64
}

func newElf(calories []uint64) *elf {
	return &elf{calories: calories}
}

func (e *elf) totalCalories() uint64 {
	total := uint64(0)
	for _, cal := range e.calories {
		total += cal
	}
	return total
}

// String returns the string representation of the elf type.
func (e *elf) String() string {
	strs := make([]string, len(e.calories))
	for i, cal := range e.calories {
		strs[i] = strconv.FormatUint(cal, 10)
	}
	return strings.Join(strs, "\n")
}
