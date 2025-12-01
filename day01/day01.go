package day01

import "fmt"

type Day struct {
	PartOnePassword uint32
	PartTwoPassword uint32
}

func (d *Day) Init() error {
	partOnePassword, partTwoPassword, err := CalculateRealPasswords()
	if err != nil {
		return err
	}

	d.PartOnePassword = partOnePassword
	d.PartTwoPassword = partTwoPassword

	return nil
}

func (d *Day) PartOne() error {
	fmt.Printf("Password: %d\n", d.PartOnePassword)

	return nil
}

func (d *Day) PartTwo() error {
	fmt.Printf("Password: %d\n", d.PartTwoPassword)

	return nil
}
