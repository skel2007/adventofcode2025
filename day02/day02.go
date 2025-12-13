package day02

import "fmt"

type Day struct {
	Sum1 uint64
	Sum2 uint64
}

func (d *Day) Init() error {
	sum1, sum2, err := SumInvalidProductIIDs(realInput)
	if err != nil {
		return err
	}

	d.Sum1 = sum1
	d.Sum2 = sum2

	return nil
}

func (d *Day) PartOne() error {
	fmt.Printf("Result: %d\n", d.Sum1)

	return nil
}

func (d *Day) PartTwo() error {
	fmt.Printf("Result: %d\n", d.Sum2)

	return nil
}
