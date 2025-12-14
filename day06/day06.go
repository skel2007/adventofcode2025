package day06

import "fmt"

type Day struct {
	GrandTotal int
}

func (d *Day) Init() error {
	grandTotal, err := GrandTotal(realInput)
	if err != nil {
		return err
	}

	d.GrandTotal = grandTotal

	return nil
}

func (d *Day) PartOne() error {
	fmt.Printf("Part One: %d\n", d.GrandTotal)

	return nil
}

func (d *Day) PartTwo() error {
	fmt.Printf("Part Two: %d\n", d.GrandTotal)

	return nil
}
