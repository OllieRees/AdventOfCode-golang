package puzzles

import (
	"fmt"
	"iter"
	"sort"
	"strconv"
	"strings"
)

func locationLists(lines []string) (xs []int, ys []int) {
	for _, line := range lines {
		fields := strings.Fields(line)
		x, _ := strconv.Atoi(fields[0])
		y, _ := strconv.Atoi(fields[1])
		xs = append(xs, x)
		ys = append(ys, y)
	}
	return xs, ys
}

func seq2Array(seq iter.Seq[string]) (xs []string) {
	for x := range seq {
		xs = append(xs, x)
	}
	return xs
}

func frequencyCount(ys []int, x int) int {
	count := 0
	for _, y := range ys {
		if y == x {
			count++
		}
	}
	return count
}

func HistorianHysteria(lines iter.Seq[string]) {
	xs, ys := locationLists(seq2Array(lines))
	diff_sum := 0
	sim_score := 0
	for _, x := range xs {
		sim_score += x * frequencyCount(ys, x)
	}

	sort.Ints(xs); sort.Ints(ys)
	for i := range xs {
		diff := xs[i] - ys[i]
		if diff < 0 {
			diff = -1 * diff
		}
		diff_sum += diff
	}

	fmt.Println(diff_sum)
	fmt.Println(sim_score)
}
