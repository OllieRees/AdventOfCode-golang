package puzzlereader

import (
	"bufio"
	"errors"
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct {
	Year int
	Day int
}

func NewPuzzle() Puzzle {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the day: ")
	day_str, err := r.ReadString('\n'); if err != nil {
		panic("Can't read day")
	}
	day, err := strconv.Atoi(strings.TrimSuffix(day_str, "\n")); if err != nil {
		panic("Can't convert day to int")
	}
	return Puzzle {2024, day}
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
	return PuzzleInput { Puzzle: Puzzle {Year: year, Day: day}, Type: input_type }
}

func (input PuzzleInput) filepath() (string, error) {
	filename, err := input.Type.filename()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("puzzles/%d/%s", input.Puzzle.Day, filename), nil
}

func (input PuzzleInput) file() (file *os.File, err error) {
	filepath, err := input.filepath()
	if err != nil {
		return nil, err
	}

    f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (input PuzzleInput) fileScanner() (r *bufio.Scanner, err error) {
	f, err := input.file()
	if err != nil {
		return nil, err
	}
	return bufio.NewScanner(f), nil
}

func (input PuzzleInput) Lines() iter.Seq[string] {
	s, _ := input.fileScanner()
	return func(yield func(string) bool) {
		for s.Scan() {
			if !yield(s.Text()) {
				return
			}
		}
	}
}
