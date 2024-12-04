package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, err := load()
	if err != nil {
		fmt.Println(err)
		return
	}
	var num1 []int
	var num2 []int
	for _, line := range lines {
		nums := strings.Split(line, "   ")
		aux1, _ := strconv.Atoi(nums[0])
		aux2, _ := strconv.Atoi(nums[1])
		num1 = append(num1, aux1)
		num2 = append(num2, aux2)
	}
	total2 := part2(num1, num2)
	total1 := part1(num1, num2)

	fmt.Println(total1)
	fmt.Println(total2)
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

func part1(num1, num2 []int) int {
	sort.Ints(num1)
	sort.Ints(num2)
	total := 0
	for i, number := range num1 {
		total += int(math.Abs(float64(number - num2[i])))
	}
	return total
}

func part2(num1, num2 []int) int {
	total := 0
	for i := 0; i < len(num1); i++ {
		mult := 0
		for j := 0; j < len(num2); j++ {
			if num1[i] == num2[j] {
				mult++
			}
		}
		total += num1[i] * mult
	}
	return total
}
