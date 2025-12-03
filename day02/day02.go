package day02

import "fmt"

type Day struct {
	SumInvalidProductIIDs uint64
}

func (d *Day) Init() error {
	sumInvalidProductIIDs, err := SumInvalidProductIIDs(realInput)
	if err != nil {
		return err
	}

	d.SumInvalidProductIIDs = sumInvalidProductIIDs

	return nil
}

func (d *Day) PartOne() error {
	fmt.Printf("Result: %d\n", d.SumInvalidProductIIDs)

	return nil
}

func (d *Day) PartTwo() error {
	return nil
}
