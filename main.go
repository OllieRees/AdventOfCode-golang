package main

import (
	"fmt"

	reader "github.com/OllieRees/AdventOfCode/puzzlereader"
)

func main() {
	puzzle_input := reader.NewPuzzleInput(2024, 0, reader.Real)
	r, _ := puzzle_input.FileScanner()
	for r.Scan() {
		fmt.Println(r.Text())
	}
}
