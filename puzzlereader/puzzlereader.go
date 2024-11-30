package puzzlereader

import (
	"errors"
	"fmt"
	"os"
)

type Puzzle struct {
	Year int
	Day int
}

type InputType int

const (
	Practice InputType = iota + 1
	Real
)

func (input InputType) filename() (string, error) {
	switch input {
	case Practice:
		return "practice.txt", nil
	case Real:
		return "real.txt", nil
	default:
		return "", errors.New("Invalid InputType")
	}
}

type PuzzleInput struct {
	Puzzle Puzzle
	Type InputType
}

func NewPuzzleInput(year int, day int, input_type InputType) PuzzleInput {
	// puzzle_input := reader.PuzzleInput { Puzzle: reader.Puzzle { Day: 0, Year: 2024 }, Type: reader.Real }
	return PuzzleInput { Puzzle: Puzzle {Year: year, Day: day}, Type: input_type }
}

func (input PuzzleInput) filepath() (string, error) {
	filename, err := input.Type.filename()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("puzzles/%d/%s", input.Puzzle.Day, filename), nil
}

func (input PuzzleInput) Read(b []byte) (n int, err error) {
	filepath, err := input.filepath()
	if err != nil {
		return 0, err
	}

    f, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}

	return f.Read(b)
}
