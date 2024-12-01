package puzzles

import (
	"fmt"
	"iter"
)

func Test(lines iter.Seq[string]) {
	for line := range lines {
		fmt.Println(line)
	}
}
