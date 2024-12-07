package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	values, eq := load()

	fmt.Println(part1(values, eq))

}

func load() ([]int, [][]int) {
	const fileName = "input.txt"
	file, _ := os.ReadFile(fileName)

	output := strings.Split(string(file), "\n")

	var values []int
	eq := make([][]int, len(output))
	for i := 0; i < len(output); i++ {
		val, _ := strconv.Atoi(strings.Split(output[i], ":")[0])
		eqs := strings.Split(strings.TrimSpace(strings.Split(output[i], ":")[1]), " ")
		for _, num := range eqs {
			n, _ := strconv.Atoi(num)
			eq[i] = append(eq[i], n)
		}
		values = append(values, val)
	}
	return values, eq
}
func part1(values []int, eqs [][]int) int {
	total := 0

	for i := 0; i < len(values); i++ {
		//fmt.Println("Inicio", values[i], eqs[i], 1, eqs[i][0])
		if check1(values[i], eqs[i], 1, eqs[i][0]) {
			total += values[i]
		}
	}
	return total
}

func check1(val int, eq []int, pos int, subtotal int) bool {
	if pos == len(eq) {
		//	fmt.Println(pos, len(eq))
		if subtotal == val {
			return true
		} else {
			return false
		}
	}
	if check1(val, eq, pos+1, subtotal+eq[pos]) || check1(val, eq, pos+1, subtotal*eq[pos]) || check1(val, eq, pos+1, concatInt(subtotal, eq[pos])) {
		return true
	}
	return false
}

func concatInt(num1, num2 int) int {
	str1 := strconv.Itoa(num1)
	str2 := strconv.Itoa(num2)
	str := str1 + str2
	num, _ := strconv.Atoi(str)
	return num
}
