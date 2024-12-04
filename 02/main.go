package main

import (
	"fmt"
	"math"
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
	// 	fmt.Println(line)
	// }
	fmt.Println(part2(lines))

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

func part1(lines []string) int {
	total := 0
	for _, line := range lines {
		valid := true
		values := strings.Split(line, " ")
		var dir int
		for j := 0; j < len(values)-1; j++ {
			num1, _ := strconv.Atoi(values[j])
			num2, _ := strconv.Atoi(values[j+1])
			if !safety(num1, num2, dir) {
				valid = false
				break
			}
			if dir == 0 {
				dir = (num1 - num2) / int(math.Abs(float64(num1-num2)))
			}
			fmt.Println(dir)
		}
		if valid {
			fmt.Println(line)
			total++
		}

	}
	return total
}

func part2(lines []string) int {
	total := 0
	for i, line := range lines {
		chance := true
		valid := true
		values := strings.Split(line, " ")
		var dir int
		for j := 0; j < len(values)-1; j++ {
			num1, _ := strconv.Atoi(values[j])
			num2, _ := strconv.Atoi(values[j+1])
			if !safety(num1, num2, dir) {
				if chance {
					chance = false
					if j+2 < len(values) {
						num3, _ := strconv.Atoi(values[j+2])
						if safety(num1, num3, dir) {
							j++
							if dir == 0 {
								dir = (num1 - num3) / int(math.Abs(float64(num1-num3)))
							}
						} else if safety(num2, num3, dir) {
							j++
							if dir == 0 {
								dir = (num2 - num3) / int(math.Abs(float64(num2-num3)))
							}
						} else {
							valid = false
							break
						}

					}
				} else {
					valid = false
					break
				}
			}
			if dir == 0 {
				dir = (num1 - num2) / int(math.Abs(float64(num1-num2)))
			}
			//fmt.Println(dir)
		}
		if valid {
			fmt.Println("Linea : ", i, line)
			total++
		}

	}
	return total
}

func safety(num1, num2, dir int) bool {
	if num1 == num2 {
		return false
	}
	if math.Abs(float64(num1-num2)) > 3 || math.Abs(float64(num1-num2)) < 1 {
		return false
	}
	if dir == 0 {
		return true
	} else if dir*(num1-num2) < 0 {
		return false
	}
	return true
}
