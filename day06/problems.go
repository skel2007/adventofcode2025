package day06

import (
	"strconv"
	"strings"
)

type Problems struct {
	Numbers   [][]int
	Operators []bool // true for multiply, false for add
}

func ParseProblems(str string) (*Problems, error) {
	problems := &Problems{}

	for _, line := range strings.Split(str, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		var numbers []int
		for _, cell := range strings.Split(line, " ") {
			if cell == "" {
				continue
			}

			switch cell[0] {
			case '+':
				problems.Operators = append(problems.Operators, false)
			case '*':
				problems.Operators = append(problems.Operators, true)
			default:
				num, err := strconv.Atoi(cell)
				if err != nil {
					return nil, err
				}

				numbers = append(numbers, num)
			}
		}

		if len(numbers) != 0 {
			problems.Numbers = append(problems.Numbers, numbers)
		}
	}

	return problems, nil
}

func (problems *Problems) GrandTotal() int {
	result := 0

	for i, operator := range problems.Operators {
		answer := 0
		if operator {
			answer = 1
		}

		for _, numbers := range problems.Numbers {
			if operator {
				answer *= numbers[i]
			} else {
				answer += numbers[i]
			}
		}

		result += answer
	}

	return result
}

func GrandTotal(str string) (int, error) {
	problems, err := ParseProblems(str)
	if err != nil {
		return 0, err
	}

	return problems.GrandTotal(), nil
}
