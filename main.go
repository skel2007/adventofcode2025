package main

import (
	"fmt"
	"os"

	"adventofcode2025/day01"
	"adventofcode2025/day02"
	"adventofcode2025/day03"
	"adventofcode2025/day04"
)

type Day interface {
	Init() error
	PartOne() error
	PartTwo() error
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main <dayN>")
		os.Exit(1)
	}

	dayN := os.Args[1]

	var day Day

	switch dayN {
	case "1":
		day = &day01.Day{}
	case "2":
		day = &day02.Day{}
	case "3":
		day = &day03.Day{}
	case "4":
		day = &day04.Day{}
	default:
		fmt.Printf("Unknown day: %s\n", dayN)
		os.Exit(1)
	}

	fmt.Printf("=== Day %s, Init =======\n", dayN)
	if err := day.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("=== Day %s, Part One ===\n", dayN)
	if err := day.PartOne(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("=== Day %s, Part Two ===\n", dayN)
	if err := day.PartTwo(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
