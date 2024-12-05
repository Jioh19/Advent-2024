package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Manual struct {
	page  int
	pages []int
}

func main() {
	lines, err := load()
	if err != nil {
		fmt.Println(err)
		return
	}

	first, second := separate(lines)
	// for _, line := range first {
	// 	fmt.Println(line)
	// }
	// fmt.Println("")
	// for _, line := range second {
	// 	fmt.Println(line)
	// }

	manuals := insertVal(first)

	fmt.Println(part2(manuals, second))
	// fmt.Println(total1)
	// fmt.Println(total2)
}

func load() ([]string, error) {
	const fileName = "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	output := strings.Split(string(file), "\n")
	return output, nil
}

func separate(lines []string) ([]string, []string) {
	var first []string
	var second []string
	var change = true
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			change = false
			continue
		}
		if change {
			first = append(first, lines[i])
		} else {
			second = append(second, lines[i])
		}
	}
	return first, second
}

func insertVal(first []string) map[string][]string {
	manuals := make(map[string][]string)
	for _, line := range first {
		values := strings.Split(line, "|")
		manuals[values[0]] = append(manuals[values[0]], values[1])
	}
	return manuals
}

func part1(manuals map[string][]string, update []string) int {
	total := 0
	for _, line := range update {
		//fmt.Println(line)
		data := strings.Split(line, ",")
		if checker(manuals, data) {
			fmt.Println(data[len(data)/2])
			num, _ := strconv.Atoi(data[len(data)/2])
			total += num
		}
	}
	return total
}

func part2(manuals map[string][]string, update []string) int {
	total := 0
	for _, line := range update {
		//fmt.Println(line)
		data := strings.Split(line, ",")
		if !checker(manuals, data) {
			swap(manuals, data)
			fmt.Println(data)
			num, _ := strconv.Atoi(data[len(data)/2])
			total += num
		}
	}
	return total
}

func checker(manuals map[string][]string, line []string) bool {
	for i := len(line) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if contains(manuals[line[i]], line[j]) {
				return false
			}
		}
	}
	return true
}

func swap(manuals map[string][]string, line []string) bool {
	for i := len(line) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if contains(manuals[line[i]], line[j]) {
				line[i], line[j] = line[j], line[i]
			}
		}
	}
	return true
}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
