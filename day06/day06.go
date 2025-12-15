package day06

import "fmt"

type Day struct {
	GrandTotal           int
	GrandTotalCephalopod int
}

func (d *Day) Init() error {
	grandTotal, grandTotalCephalopod, err := GrandTotal(realInput)
	if err != nil {
		return err
	}

	d.GrandTotal = grandTotal
	d.GrandTotalCephalopod = grandTotalCephalopod

	return nil
}

func (d *Day) PartOne() error {
	fmt.Printf("Part One: %d\n", d.GrandTotal)

	return nil
}

func (d *Day) PartTwo() error {
	fmt.Printf("Part Two: %d\n", d.GrandTotalCephalopod)

	return nil
}
