package day02

import (
	"fmt"
	"strings"
)

type Range struct {
	Start ProductID
	End   ProductID
}

func ParseRange(str string) (*Range, error) {
	dash := strings.Index(str, "-")
	if dash == -1 {
		return nil, fmt.Errorf("invalid range: '%s'", str)
	}

	return &Range{
		Start: ProductID(str[:dash]),
		End:   ProductID(str[dash+1:]),
	}, nil
}

func (r *Range) String() string {
	return fmt.Sprintf("%s-%s", r.Start, r.End)
}

func (r *Range) FindInvalidIDs() ([]InvalidID, error) {
	var result []InvalidID

	for n := 2; n <= len(r.End); n++ {
		invalidIDs, err := r.findInvalidIDs(n)
		if err != nil {
			return nil, err
		}

		result = append(result, invalidIDs...)
	}

	return result, nil
}

func (r *Range) findInvalidIDs(n int) ([]InvalidID, error) {
	startSeq, err := r.Start.CeilSeq(n)
	if err != nil {
		return nil, err
	}

	endSeq, err := r.End.FloorSeq(n)
	if err != nil {
		return nil, err
	}

	var result []InvalidID
	for seq := startSeq; seq <= endSeq; seq++ {
		invalidID := NewInvalidID(seq, n)
		result = append(result, invalidID)
	}

	return result, nil
}
