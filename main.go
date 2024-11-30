package main

import (
	reader "github.com/OllieRees/AdventOfCode/puzzlereader"
)

func main() {
	puzzle := reader.NewPuzzle()
	puzzle.Run(reader.Practice)
	puzzle.Run(reader.Real)
}
