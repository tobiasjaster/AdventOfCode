// Code generated by aocgen; DO NOT EDIT.
package year2023

import (
	"slices"
	"strconv"
	"strings"
	"unicode"
	"sort"
)

type Day01 struct{}

func (p Day01) PartA(lines []string) any {
	calibrations := []string{}
	for _, line := range lines {
		byteline := []byte(line)
		value := ""
		for _, char := range byteline {
			if unicode.IsDigit(rune(char)) {
				value = value + string(char)
				break
			}
		}
		slices.Reverse(byteline)
		for _, char := range byteline {
			if unicode.IsDigit(rune(char)) {
				value = value + string(char)
				break
			}
		}
		calibrations = append(calibrations, value)
	}
	number := 0
	for _, calibration := range calibrations {
		i, _ := strconv.Atoi(calibration)
		number = number + i
	}
	return number
}

var (
	numberMap = map[string]string {
		"one":"o1e",
		"two":"t2o",
		"three":"t3e",
		"four":"f4r",
		"five":"f5e",
		"six":"s6x",
		"seven":"s7n",
		"eight":"e8t",
		"nine": "n9e",
	}
)

func (p Day01) PartB(lines []string) any {
	calibrations := []string{}
	for _, line := range lines {
		numbers :=map[int]string{}
		for str, number := range numberMap {
			line = strings.Replace(line,str,number,-1)
		}
		for idx, char := range []byte(line) {
			if unicode.IsDigit(rune(char)) {
				numbers[idx] = string(char)
			}
		}
		keys := []int{}
		reverseKeys := []int{}
		for k := range numbers {
			keys = append(keys, k)
			reverseKeys = append(reverseKeys, k)
		}
		sort.Ints(keys)
		sort.Ints(reverseKeys)
		slices.Reverse(reverseKeys)
		value := numbers[keys[0]]+numbers[reverseKeys[0]]
		calibrations = append(calibrations, value)
	}
	number := 0
	for _, calibration := range calibrations {
		i, _ := strconv.Atoi(calibration)
		number = number + i
	}
	return number
}
