package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := load()
	if err != nil {
		fmt.Println(err)
		return
	}

	// for _, line := range lines {
	//fmt.Println(lines)
	// }
	fmt.Println(part2(lines))
	// fmt.Println(total1)
	// fmt.Println(total2)
}

func load() (string, error) {
	const fileName = "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	output := file
	return string(output), nil
}

func part1(lines string) int {
	total := 0
	inputs := strings.Split(lines, "mul(")
	for i := 0; i < len(inputs); i++ {
		total += calculate(inputs[i])
	}
	return total
}

func part2(lines string) int {
	inputs := strings.Split(lines, "do")
	total := part1(inputs[0])
	inputs = inputs[1:]
	for _, input := range inputs {

		if input[:2] == "()" {
			total += part1(input)
		}
	}

	return total
}

func calculate(line string) int {
	input := strings.Split(line, ")")[0]
	values := strings.Split(input, ",")
	if len(values) != 2 {
		return 0
	}
	num1, err1 := strconv.Atoi(values[0])
	num2, err2 := strconv.Atoi(values[1])
	if err1 != nil || err2 != nil {
		return 0
	}
	return num1 * num2
}
