package day04

import "fmt"

type Day struct {
	Count      int
	TotalCount int
}

func (d *Day) Init() error {
	count, totalCount, err := FindAccessibleRollsCount(realInput)
	if err != nil {
		return err
	}

	d.Count = count
	d.TotalCount = totalCount

	return nil
}

func (d *Day) PartOne() error {
	fmt.Printf("Result: %d\n", d.Count)

	return nil
}

func (d *Day) PartTwo() error {
	fmt.Printf("Result: %d\n", d.TotalCount)

	return nil
}
