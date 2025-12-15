package day06

import (
	"fmt"
	"strconv"
	"strings"
)

type Problems struct {
	Numbers   [][]string
	Operators []bool // true for multiply, false for add
}

func ParseProblems(str string) (*Problems, error) {
	var lines []string
	for _, line := range strings.Split(str, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		lines = append(lines, line)
	}

	if len(lines) <= 1 {
		return nil, fmt.Errorf("invalid input")
	}

	columns := []int{0}
	for i := range lines[0] {
		next := true
		for _, line := range lines {
			if line[i] != ' ' {
				next = false
				break
			}
		}

		if next {
			columns = append(columns, i+1)
		}
	}

	columns = append(columns, len(lines[0]))

	numbers := lines[:len(lines)-1]
	operators := lines[len(lines)-1]

	problems := &Problems{
		Numbers: make([][]string, len(numbers)),
	}

	for i := 0; i < len(columns)-1; i++ {
		start := columns[i]
		end := columns[i+1]

		for i, line := range numbers {
			problems.Numbers[i] = append(problems.Numbers[i], line[start:end])
		}

		operator := strings.TrimSpace(operators[start:end])
		switch operator {
		case "*":
			problems.Operators = append(problems.Operators, true)
		case "+":
			problems.Operators = append(problems.Operators, false)
		default:
			return nil, fmt.Errorf("invalid operator: %s", operator)
		}
	}

	return problems, nil
}

func (problems *Problems) GrandTotal() (int, error) {
	result := 0

	for i, operator := range problems.Operators {
		answer := 0
		if operator {
			answer = 1
		}

		for _, numbers := range problems.Numbers {
			number, err := strconv.Atoi(strings.TrimSpace(numbers[i]))
			if err != nil {
				return 0, err
			}

			if operator {
				answer *= number
			} else {
				answer += number
			}
		}

		result += answer
	}

	return result, nil
}

func (problems *Problems) GrandTotalCephalopod() (int, error) {
	result := 0

	for i, operator := range problems.Operators {
		answer := 0
		if operator {
			answer = 1
		}

		for j := range problems.Numbers[0][i] {
			str := ""
			for _, numbers := range problems.Numbers {
				str += string(numbers[i][j])
			}

			if strings.TrimSpace(str) == "" {
				continue
			}

			number, err := strconv.Atoi(strings.TrimSpace(str))
			if err != nil {
				return 0, err
			}

			if operator {
				answer *= number
			} else {
				answer += number
			}
		}

		result += answer
	}

	return result, nil
}

func GrandTotal(str string) (int, int, error) {
	problems, err := ParseProblems(str)
	if err != nil {
		return 0, 0, err
	}

	total, err := problems.GrandTotal()
	if err != nil {
		return 0, 0, err
	}

	totalCephalopod, err := problems.GrandTotalCephalopod()
	if err != nil {
		return 0, 0, err
	}

	return total, totalCephalopod, err
}
