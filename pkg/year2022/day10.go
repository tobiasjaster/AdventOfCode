// Code generated by aocgen; DO NOT EDIT.
package year2022

import (
	"fmt"
	"strconv"
	"strings"
)

type Day10 struct{}

func (p Day10) PartA(lines []string) any {
	var sum int
	checkCycle := []int{20,60,100,140,180,220}
	var currentCycleRegister []int
	register := 1
	currentCycleRegister = append(currentCycleRegister, register)
	for _, line := range lines{
		if line=="" {continue}
		if strings.HasPrefix(line, "noop") {
			currentCycleRegister = append(currentCycleRegister, register)
		} else if strings.HasPrefix(line, "addx"){
			currentCycleRegister = append(currentCycleRegister, register)
			valueString := strings.Split(line, " ")
			value, _ := strconv.Atoi(valueString[1])
			register += value
			currentCycleRegister = append(currentCycleRegister, register)
		}else{
			fmt.Println("fehler")
		}
	}
	for _, value := range checkCycle{
		sum += currentCycleRegister[value-1]*value
	}
	return sum
}

func (p Day10) PartB(lines []string) any {
	display := make([][]string, 6)
	display[0] = make([]string, 40)
	display[1] = make([]string, 40)
	display[2] = make([]string, 40)
	display[3] = make([]string, 40)
	display[4] = make([]string, 40)
	display[5] = make([]string, 40)
	var currentCycleRegister []int
	register := 1
	currentCycleRegister = append(currentCycleRegister, register)
	for _, line := range lines{
		if line=="" {continue}
		if strings.HasPrefix(line, "noop") {
			currentCycleRegister = append(currentCycleRegister, register)
		} else if strings.HasPrefix(line, "addx"){
			currentCycleRegister = append(currentCycleRegister, register)
			valueString := strings.Split(line, " ")
			value, _ := strconv.Atoi(valueString[1])
			register += value
			currentCycleRegister = append(currentCycleRegister, register)
		}else{
			fmt.Println("fehler")
		}
	}
	for i, value := range currentCycleRegister[:240]{
		var print string
		if i%40+1 == value || i%40+1 == value+1 || i%40+1 == value +2 {
			print = "#"
		}else{
			print = "."
		}
		display[int(i/40)][i%40] = print
	}
	for _, value := range display {
		fmt.Println(value)
	}
	return "PLPAFBCL"
}
