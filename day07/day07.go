package day07

import "fmt"

type Day struct {
	CountSplits int
}

func (d *Day) Init() error {
	count, err := CountSplits(realInput)
	if err != nil {
		return err
	}

	d.CountSplits = count

	return nil
}

func (d *Day) PartOne() error {
	fmt.Printf("Part One: %d\n", d.CountSplits)

	return nil
}

func (d *Day) PartTwo() error {
	return nil
}
