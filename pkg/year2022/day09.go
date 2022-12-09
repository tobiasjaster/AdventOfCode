// Code generated by aocgen; DO NOT EDIT.
package year2022

import (
	"fmt"
)

type Day09 struct{}

func getInitialChain(length int) [][]int {
	chain := [][]int{}
	for i := 0; i<length; i++ {
		chain = append(chain, []int{0,0})
	}
	return chain
}

func removeDuplicateValues(intSlice [][]int) [][]int {
    keys := make(map[string]bool)
    list := [][]int{}
 
    // If the key(values of the slice) is not equal
    // to the already present value in new slice (list)
    // then we append it. else we jump on another element.
    for _, entry := range intSlice {
		entryString := fmt.Sprintf("%d %d", entry[0], entry[1])
        if _, value := keys[entryString]; !value {
            keys[entryString] = true
            list = append(list, entry)
        }
    }
    return list
}

func returnMove(diff int) int {
	if diff>0 {
		return diff-1
	} else if diff<0 {
		return diff+1
	} 
	return 0
}

func moveHead(direction string, position *[]int) {
	switch direction{
	case "L":
		(*position)[0]--
	case "R":
		(*position)[0]++
	case "U":
		(*position)[1]++
	case "D":
		(*position)[1]--
	}
}

func move(direction string, count int, positions *[][]int, lastElemList *[][]int) {
	length := len((*positions))
	for i := 0; i<count; i++{
		moveHead(direction, &((*positions)[0]))
		for i := 1; i<length; i++ {
			dif_x := (*positions)[i][0]-(*positions)[i-1][0]
			dif_y := (*positions)[i][1]-(*positions)[i-1][1]
			move_x := returnMove(dif_x)
			move_y := returnMove(dif_y)
			if move_x == 0 && move_y == 0 {
				break
			}
			(*positions)[i][0] = (*positions)[i-1][0]+move_x
			(*positions)[i][1] = (*positions)[i-1][1]+move_y
			if i == (length-1) {
				(*lastElemList) = append((*lastElemList), []int{(*positions)[length-1][0], (*positions)[length-1][1]})
			}
		}
	}
}

func (p Day09) PartA(lines []string) any {
	var grid [][]int
	var length int = 2
	positions := getInitialChain(length)
	grid = append(grid, []int{positions[length-1][0], positions[length-1][1]})
	var direction string
	var count int
	for _, line := range lines {
		if line == "" {continue}
		fmt.Sscanf(line, "%s %d", &direction, &count)
		move(direction, count, &positions, &grid)
	}
	filteredGrid := removeDuplicateValues(grid)
	return len(filteredGrid)
}

func (p Day09) PartB(lines []string) any {
	var grid [][]int 
	var length int = 10
	positions := getInitialChain(length)
	grid = append(grid, []int{positions[length-1][0], positions[length-1][1]})
	var direction string
	var count int
	for _, line := range lines {
		if line == "" {continue}
		fmt.Sscanf(line, "%s %d", &direction, &count)
		move(direction, count, &positions, &grid)
	}
	filteredGrid := removeDuplicateValues(grid)
	return len(filteredGrid)
}
