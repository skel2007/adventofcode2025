package day02

import (
	"strconv"
	"strings"
)

type ProductID string

func (id ProductID) CeilSeq(n int) (uint64, error) {
	if len(id)%n == 0 {
		first, err := strconv.ParseUint(string(id[:len(id)/n]), 10, 64)
		if err != nil {
			return 0, err
		}

		prev := first
		for i := 1; i < n; i++ {
			start := i * (len(id) / n)
			end := (i + 1) * (len(id) / n)

			next, err := strconv.ParseUint(string(id[start:end]), 10, 64)
			if err != nil {
				return 0, err
			}

			if prev < next {
				return first + 1, nil
			}

			if prev > next {
				return first, nil
			}

			prev = next
		}

		return first, nil
	}

	result := "1" + strings.Repeat("0", len(id)/n)
	return strconv.ParseUint(result, 10, 64)
}

func (id ProductID) FloorSeq(n int) (uint64, error) {
	if len(id) < n {
		return 0, nil
	}

	if len(id)%n == 0 {
		first, err := strconv.ParseUint(string(id[:len(id)/n]), 10, 64)
		if err != nil {
			return 0, err
		}

		prev := first
		for i := 1; i < n; i++ {
			start := i * (len(id) / n)
			end := (i + 1) * (len(id) / n)

			next, err := strconv.ParseUint(string(id[start:end]), 10, 64)
			if err != nil {
				return 0, err
			}

			if prev < next {
				return first, nil
			}

			if prev > next {
				return first - 1, nil
			}

			prev = next
		}

		return first, nil
	}

	result := strings.Repeat("9", len(id)/n)
	return strconv.ParseUint(result, 10, 64)
}

func SumInvalidProductIIDs(str string) (uint64, uint64, error) {
	var sumSimple, sumAll uint64
	var uniqueIDs = make(map[uint64]struct{})

	for _, line := range strings.Split(str, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		for _, s := range strings.Split(line, ",") {
			r, err := ParseRange(s)
			if err != nil {
				return 0, 0, err
			}

			invalidIDs, err := r.FindInvalidIDs()
			if err != nil {
				return 0, 0, err
			}

			for _, id := range invalidIDs {
				value, err := id.Value()
				if err != nil {
					return 0, 0, err
				}

				if _, ok := uniqueIDs[value]; ok {
					continue
				}

				if id.n == 2 {
					sumSimple += value
				}

				sumAll += value
				uniqueIDs[value] = struct{}{}
			}
		}
	}

	return sumSimple, sumAll, nil
}
