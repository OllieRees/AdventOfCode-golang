package main

import (
	"fmt"

	reader "github.com/OllieRees/AdventOfCode/puzzlereader"
)

func main() {
	puzzle_input := reader.NewPuzzleInput(2024, 0, reader.Real)
	for line := range puzzle_input.Lines() {
		fmt.Println(line)
	}
}
