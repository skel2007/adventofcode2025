package day07

import (
	"fmt"
	"strings"
)

type Cell int

const (
	CellStart Cell = iota
	CellSpace
	CellSplitter
)

func ParseCell(r rune) (Cell, error) {
	switch r {
	case 'S':
		return CellStart, nil
	case '.':
		return CellSpace, nil
	case '^':
		return CellSplitter, nil
	default:
		return 0, fmt.Errorf("invalid cell %q", r)
	}
}

type Manifold struct {
	diagram [][]Cell
}

func ParseManifold(str string) (*Manifold, error) {
	diagram := make([][]Cell, 0)

	for _, line := range strings.Split(str, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		level := make([]Cell, 0)

		for _, r := range line {
			cell, err := ParseCell(r)
			if err != nil {
				return nil, err
			}

			level = append(level, cell)
		}

		diagram = append(diagram, level)
	}

	return &Manifold{
		diagram: diagram,
	}, nil
}

func (m *Manifold) Next(i int, j int) (map[int]int, int, error) {
	cell := m.diagram[i][j]

	next := make(map[int]int)
	splits := 0

	switch cell {
	case CellSpace:
		next[j]++
	case CellSplitter:
		splits++

		next[j-1]++
		next[j+1]++
	default:
		return nil, 0, fmt.Errorf("invalid cell [%d; %d]: %q", i, j, cell)
	}

	return next, splits, nil
}

func (m *Manifold) Teleport() (int, int, error) {
	tachyon := make(map[int]int)

	countSplits := 0

	for i, line := range m.diagram {
		if len(tachyon) == 0 {
			for j, cell := range line {
				if cell == CellStart {
					tachyon[j] = 1
				}
			}

			continue
		}

		nextTachyon := make(map[int]int)
		for j, t := range tachyon {
			cell := m.diagram[i][j]

			switch cell {
			case CellSpace:
				nextTachyon[j] += t
			case CellSplitter:
				countSplits++

				nextTachyon[j-1] += t
				nextTachyon[j+1] += t
			default:
				return 0, 0, fmt.Errorf("invalid cell [%d; %d]: %q", i, j, cell)
			}
		}

		tachyon = nextTachyon
	}

	countTracks := 0
	for _, t := range tachyon {
		countTracks += t
	}

	return countSplits, countTracks, nil
}

func Teleport(str string) (int, int, error) {
	m, err := ParseManifold(str)
	if err != nil {
		return 0, 0, err
	}

	return m.Teleport()
}
