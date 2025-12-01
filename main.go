package main

import (
	"fmt"
	"os"

	"adventofcode2025/day01"
)

type Day interface {
	PartOne() error
	PartTwo() error
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main <dayN>")
		os.Exit(1)
	}

	var day Day

	switch os.Args[1] {
	case "1":
		day = day01.Day{}
	default:
		fmt.Printf("Unknown day: %s\n", os.Args[1])
		os.Exit(1)
	}

	if err := day.PartOne(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := day.PartTwo(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
