package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type ProductID uint64

func (id ProductID) CeilSeq() (uint64, error) {
	s := strconv.FormatUint(uint64(id), 10)
	if len(s)%2 == 1 {
		var result uint64 = 1
		for i := 0; i < len(s)/2; i++ {
			result *= 10
		}

		return result, nil
	}

	mod := uint64(1)
	for i := 0; i < len(s)/2; i++ {
		mod *= 10
	}

	first := uint64(id) / mod
	second := uint64(id) % mod

	if first < second {
		return first + 1, nil
	}

	return first, nil
}

func (id ProductID) FloorSeq() (uint64, error) {
	s := strconv.FormatUint(uint64(id), 10)
	if len(s)%2 == 1 {
		var result uint64 = 1
		for i := 0; i < len(s)/2; i++ {
			result *= 10
		}

		return result - 1, nil
	}

	mod := uint64(1)
	for i := 0; i < len(s)/2; i++ {
		mod *= 10
	}

	first := uint64(id) / mod
	second := uint64(id) % mod

	if first <= second {
		return first, nil
	}

	return first - 1, nil
}

type Range struct {
	Start ProductID
	End   ProductID
}

func ParseRange(str string) (*Range, error) {
	dash := strings.Index(str, "-")
	if dash == -1 {
		return nil, fmt.Errorf("invalid range: '%s'", str)
	}

	start, err := strconv.Atoi(str[:dash])
	if err != nil {
		return nil, err
	}

	end, err := strconv.Atoi(str[dash+1:])
	if err != nil {
		return nil, err
	}

	return &Range{
		Start: ProductID(start),
		End:   ProductID(end),
	}, nil
}

func (r *Range) FindInvalidIDs() ([]ProductID, error) {
	startSeq, err := r.Start.CeilSeq()
	if err != nil {
		return nil, err
	}

	endSeq, err := r.End.FloorSeq()
	if err != nil {
		return nil, err
	}

	var result []ProductID
	for seq := startSeq; seq <= endSeq; seq++ {
		invalidID := seqToInvalidID(seq)
		result = append(result, invalidID)
	}

	return result, nil
}

func seqToInvalidID(seq uint64) ProductID {
	s := strconv.FormatUint(seq, 10)

	result := seq
	for i := 0; i < len(s); i++ {
		result *= 10
	}

	result += seq
	return ProductID(result)
}

func SumInvalidProductIIDs(str string) (uint64, error) {
	var result uint64
	for _, line := range strings.Split(str, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		for _, s := range strings.Split(line, ",") {
			r, err := ParseRange(s)
			if err != nil {
				return 0, err
			}

			invalidIDs, err := r.FindInvalidIDs()
			if err != nil {
				return 0, err
			}

			for _, id := range invalidIDs {
				result += uint64(id)
			}
		}
	}

	return result, nil
}
