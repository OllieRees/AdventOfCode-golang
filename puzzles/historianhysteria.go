package puzzles

import (
	"fmt"
	"iter"
	"sort"
	"strconv"
	"strings"
)

type LocationLists struct {
	Left []int
	Right []int
}

func fromInput(lines []string) LocationLists {
	xs, ys := locationLists(lines)
	return LocationLists {xs, ys}
}

func (lists *LocationLists) TotalDistance() int {
	diff_sum := 0
	sort.Ints(lists.Left); sort.Ints(lists.Right)
	for i := range lists.Left {
		diff_sum += lists.getDistance(lists.Left[i], lists.Right[i])
	}
	return diff_sum
}

func (lists *LocationLists) SimilarityScore() int {
	sim_score := 0
	for _, x := range lists.Left {
		sim_score += x * frequencyCount(lists.Right, x)
	}
	return sim_score
}

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

func (lists *LocationLists) getDistance(x int, y int) int {
	diff := x - y
	if diff < 0 {
		return -1 * diff
	}
	return diff
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

func seq2Array(seq iter.Seq[string]) (xs []string) {
	for x := range seq {
		xs = append(xs, x)
	}
	return xs
}

func HistorianHysteria(lines iter.Seq[string]) {
	lists := fromInput(seq2Array(lines))
	fmt.Printf("Total Distance Between Lists: %d\n", lists.TotalDistance())
	fmt.Printf("Similarity Score: %d\n", lists.SimilarityScore())
}
