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

func (m *Manifold) CountSplits() (int, error) {
	tachyon := make(map[int]bool)
	result := 0

	for i, line := range m.diagram {
		if len(tachyon) == 0 {
			for j, cell := range line {
				if cell == CellStart {
					tachyon[j] = true
				}
			}

			continue
		}

		nextTachyon := make(map[int]bool)
		for j := range tachyon {
			cell := line[j]
			switch cell {
			case CellSpace:
				nextTachyon[j] = true
			case CellSplitter:
				result++

				nextTachyon[j-1] = true
				nextTachyon[j+1] = true
			default:
				return 0, fmt.Errorf("invalid cell [%d; %d]: %q", i, j, cell)
			}
		}

		tachyon = nextTachyon
	}

	return result, nil
}

func CountSplits(str string) (int, error) {
	m, err := ParseManifold(str)
	if err != nil {
		return 0, err
	}

	return m.CountSplits()
}
