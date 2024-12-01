package puzzlereader

import (
	"bufio"
	"errors"
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"


	puzzles "github.com/OllieRees/AdventOfCode/puzzles"
)

type Puzzle struct {
	Year int
	Day int
}

func NewPuzzle() Puzzle {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the day: ")
	day_str, err := r.ReadString('\n'); if err != nil {
		panic("Can't read day")
	}
	day, err := strconv.Atoi(strings.TrimSuffix(day_str, "\n")); if err != nil {
		panic("Can't convert day to int")
	}
	return Puzzle {2024, day}
}

func (puzzle Puzzle) puzzleRoutine() func(iter.Seq[string]) {
	switch puzzle.Year {
	case 2024:
		switch puzzle.Day {
		case 0:
			return puzzles.Test
		default:
			panic(fmt.Sprintf("Can't find puzzle for year %s and day %s", puzzle.Year, puzzle.Day))
		}
	default:
		panic(fmt.Sprintf("Can't find puzzle for year %s", puzzle.Year))
	}
}

func (puzzle Puzzle) Run(input_type InputType) {
	puzzle_input := NewPuzzleInput(puzzle.Year, puzzle.Day, input_type)
	fmt.Println(puzzle_input)
	puzzle.puzzleRoutine()(puzzle_input.Lines())
}

type InputType int

const (
	Practice InputType = iota + 1
	Real
)

func (input InputType) String() string {
	switch input {
	case Practice:
		return "Practice"
	case Real:
		return "Real"
	default:
		return ""
	}
}

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
	File *os.File `default:"nil"`
}

func NewPuzzleInput(year int, day int, input_type InputType) PuzzleInput {
	return PuzzleInput { Puzzle: Puzzle {Year: year, Day: day}, Type: input_type }
}

func (input PuzzleInput) String() string {
	return fmt.Sprintf("%s Puzzle for Year %d and Day %d", input.Type, input.Puzzle.Year, input.Puzzle.Day)
}

func (input PuzzleInput) filepath() (string, error) {
	filename, err := input.Type.filename()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("input/%d/%s", input.Puzzle.Day, filename), nil
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
	input.File = f
	return f, err
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
				input.File.Close()
				return
			}
		}
	}
}
