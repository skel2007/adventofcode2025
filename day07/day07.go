package day07

import "fmt"

type Day struct {
	CountSplits int
	CountTracks int
}

func (d *Day) Init() error {
	countSplits, countTracks, err := Teleport(realInput)
	if err != nil {
		return err
	}

	d.CountSplits = countSplits
	d.CountTracks = countTracks

	return nil
}

func (d *Day) PartOne() error {
	fmt.Printf("Part One: %d\n", d.CountSplits)

	return nil
}

func (d *Day) PartTwo() error {
	fmt.Printf("Part Two: %d\n", d.CountTracks)

	return nil
}
