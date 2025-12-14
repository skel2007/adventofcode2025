package day03

import "fmt"

type Day struct {
	Joltage2  uint64
	Joltage12 uint64
}

func (d *Day) Init() error {
	joltage2, joltage12, err := TotalJoltage(realInput)
	if err != nil {
		return err
	}

	d.Joltage2 = joltage2
	d.Joltage12 = joltage12

	return nil
}

func (d *Day) PartOne() error {
	fmt.Printf("Result: %d\n", d.Joltage2)

	return nil
}

func (d *Day) PartTwo() error {
	fmt.Printf("Result: %d\n", d.Joltage12)

	return nil
}
