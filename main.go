package main

import (
	"fmt"

	reader "github.com/OllieRees/AdventOfCode/puzzlereader"
)

func main() {
	puzzle_input := reader.NewPuzzleInput(2024, 0, reader.Real)
	b := make([]byte, 20);
	puzzle_input.Read(b);
	fmt.Println(string(b));
}
