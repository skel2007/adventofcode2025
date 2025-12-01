package day01

import "fmt"

type Day struct{}

func (d Day) PartOne() error {
	password, err := CalculateRealPassword()
	if err != nil {
		return err
	}

	fmt.Printf("Password: %d\n", password)

	return nil
}

func (d Day) PartTwo() error {
	return nil
}
