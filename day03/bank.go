package day03

import (
	"strconv"
	"strings"
)

type Bank struct {
	batteries string
}

func (b *Bank) Joltage(n int) (uint64, error) {
	result := make([]int, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			result[i] = result[i-1] + 1
		}

		for j := result[i] + 1; j < len(b.batteries)-n+i+1; j++ {
			if b.batteries[j] > b.batteries[result[i]] {
				result[i] = j
			}
		}
	}

	joltage := ""
	for _, i := range result {
		u := string(b.batteries[i])
		joltage += u
	}

	return strconv.ParseUint(joltage, 10, 64)
}

func TotalJoltage(str string) (uint64, uint64, error) {
	result2 := uint64(0)
	result12 := uint64(0)

	for _, line := range strings.Split(str, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		bank := Bank{line}

		joltage2, err := bank.Joltage(2)
		if err != nil {
			return 0, 0, err
		}

		joltage12, err := bank.Joltage(12)
		if err != nil {
			return 0, 0, err
		}

		result2 += joltage2
		result12 += joltage12
	}

	return result2, result12, nil
}
