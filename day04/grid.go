package day04

import (
	"fmt"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type Grid struct {
	Rolls [][]bool
}

func ParseGrid(str string) (*Grid, error) {
	var rolls [][]bool

	for _, line := range strings.Split(str, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		row := make([]bool, len(line))
		for _, c := range line {
			switch c {
			case '.':
				row = append(row, false)
			case '@':
				row = append(row, true)
			default:
				return nil, fmt.Errorf("grid: invalid character '%c'", c)
			}
		}

		rolls = append(rolls, row)
	}

	return &Grid{
		Rolls: rolls,
	}, nil
}

func (g *Grid) FindAccessibleRolls() []Coordinate {
	var result []Coordinate
	for i := 0; i < len(g.Rolls); i++ {
		for j := 0; j < len(g.Rolls[i]); j++ {
			if !g.Rolls[i][j] {
				continue
			}

			count := 0
			for k := -1; k <= 1; k++ {
				if i+k < 0 || i+k >= len(g.Rolls) {
					continue
				}

				for l := -1; l <= 1; l++ {
					if j+l < 0 || j+l >= len(g.Rolls[i]) {
						continue
					}

					if k == 0 && l == 0 {
						continue
					}

					if g.Rolls[i+k][j+l] {
						count++
					}
				}
			}

			if count < 4 {
				result = append(result, Coordinate{i, j})
			}
		}
	}

	return result
}

func (g *Grid) Remove(coordinates ...Coordinate) {
	for _, coordinate := range coordinates {
		g.Rolls[coordinate.X][coordinate.Y] = false
	}
}

func FindAccessibleRollsCount(str string) (int, int, error) {
	grid, err := ParseGrid(str)
	if err != nil {
		return 0, 0, err
	}

	accessibleRolls := grid.FindAccessibleRolls()
	count := len(accessibleRolls)

	totalCount := count
	for {
		grid.Remove(accessibleRolls...)

		accessibleRolls = grid.FindAccessibleRolls()
		if len(accessibleRolls) == 0 {
			break
		}

		totalCount += len(accessibleRolls)
	}

	return count, totalCount, nil
}
