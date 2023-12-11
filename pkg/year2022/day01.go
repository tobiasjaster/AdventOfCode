package year2022

import (
	"sort"
	"strconv"
	"strings"
)

type Day01 struct{}

type Elve struct {
	calories int
}

func returnElves(lines []string) []Elve {
	var elves []Elve
	var elve Elve
	for _, line := range lines {
		line = strings.ReplaceAll(line, "\r", "")
		if line != "" {
			calorie, _ := strconv.Atoi(strings.ReplaceAll(line, "\r", ""))
			elve.calories += calorie
		} else {
			elves = append(elves, elve)
			elve = Elve{}
		}
	}
	elves = append(elves, elve)
	return elves
}

func (p Day01) PartA(lines []string) any {
	var elves []Elve = returnElves(lines)
	sort.Slice(elves, func(i, j int) bool { return elves[i].calories > elves[j].calories })
	return elves[0].calories
}

func (p Day01) PartB(lines []string) any {
	var elves []Elve = returnElves(lines)
	sort.Slice(elves, func(i, j int) bool { return elves[i].calories > elves[j].calories })
	return elves[0].calories + elves[1].calories + elves[2].calories
}
