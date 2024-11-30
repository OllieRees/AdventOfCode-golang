package main

import (
	"fmt"

	reader "github.com/OllieRees/AdventOfCode/puzzlereader"
)

func main() {
	puzzle := reader.NewPuzzle()
	puzzle_input := reader.NewPuzzleInput(puzzle.Year, puzzle.Day, reader.Real)
	for line := range puzzle_input.Lines() {
		fmt.Println(line)
	}
}
