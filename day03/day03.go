package day03

import "fmt"

type Day struct {
	Joltage uint64
}

func (d *Day) Init() error {
	joltage, err := TotalJoltage(realInput)
	if err != nil {
		return err
	}

	d.Joltage = joltage

	return nil
}

func (d *Day) PartOne() error {
	fmt.Printf("Result: %d\n", d.Joltage)

	return nil
}

func (d *Day) PartTwo() error {
	return nil
}
