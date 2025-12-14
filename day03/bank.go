package day03

import (
	"strconv"
	"strings"
)

type Bank struct {
	batteries string
}

func (b *Bank) Joltage() (uint64, error) {
	f := 0
	for i := 1; i < len(b.batteries)-1; i++ {
		if b.batteries[i] > b.batteries[f] {
			f = i
		}
	}

	s := f + 1
	for i := f + 2; i < len(b.batteries); i++ {
		if b.batteries[i] > b.batteries[s] {
			s = i
		}
	}

	result := string(b.batteries[f]) + string(b.batteries[s])
	return strconv.ParseUint(result, 10, 64)
}

func TotalJoltage(str string) (uint64, error) {
	result := uint64(0)
	for _, line := range strings.Split(str, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		bank := Bank{line}

		joltage, err := bank.Joltage()
		if err != nil {
			return 0, err
		}

		result += joltage
	}

	return result, nil
}
