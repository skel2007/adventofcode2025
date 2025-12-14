package day05

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func (r *Range) IsFresh(id int) bool {
	if id >= r.Start && id <= r.End {
		return true
	}

	return false
}

func ParseRange(line string) (*Range, error) {
	dash := strings.Index(line, "-")
	if dash == -1 {
		return nil, fmt.Errorf("invalid range: %s", line)
	}

	start, err := strconv.Atoi(line[:dash])
	if err != nil {
		return nil, err
	}

	end, err := strconv.Atoi(line[dash+1:])
	if err != nil {
		return nil, err
	}

	return &Range{
		Start: start,
		End:   end,
	}, nil
}

type Ranges struct {
	Ranges []*Range
}

func (ranges *Ranges) IsFresh(id int) bool {
	for _, r := range ranges.Ranges {
		if r.IsFresh(id) {
			return true
		}
	}

	return false
}

func (ranges *Ranges) TotalFreshCount() int {
	type point struct {
		id    int
		start bool
	}

	var ids []point

	for _, r := range ranges.Ranges {
		ids = append(ids, point{r.Start, true})
		ids = append(ids, point{r.End, false})
	}

	sort.Slice(ids, func(i, j int) bool {
		l := ids[i]
		r := ids[j]

		if l.id == r.id {
			return l.start && !r.start
		}

		return l.id < r.id
	})

	result := 0

	openRanges := 0
	start := 0

	for _, id := range ids {
		if id.start {
			if openRanges == 0 {
				start = id.id
			}

			openRanges++
		} else {
			openRanges--

			if openRanges == 0 {
				result += id.id - start + 1
			}
		}
	}

	return result
}

func CountFreshIDs(str string) (int, int, error) {
	isRangesParsing := true
	ranges := &Ranges{}

	var ingredients []int

	for _, line := range strings.Split(str, "\n") {
		if strings.TrimSpace(line) == "" {
			isRangesParsing = false
			continue
		}

		if isRangesParsing {
			r, err := ParseRange(line)
			if err != nil {
				return 0, 0, err
			}

			ranges.Ranges = append(ranges.Ranges, r)
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				return 0, 0, err
			}

			ingredients = append(ingredients, id)
		}
	}

	count := 0
	for _, id := range ingredients {
		if ranges.IsFresh(id) {
			count++
		}
	}

	totalCount := ranges.TotalFreshCount()

	return count, totalCount, nil
}
