package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	line, err := load()
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(line)
	results1 := part1(line)

	fmt.Println(part2(results1))
}

func part1(line []byte) []string {
	result := []string{}
	for i, char := range line {
		num, _ := strconv.Atoi(string(char))
		if i%2 != 0 {
			for j := 0; j < num; j++ {
				result = append(result, ".")
			}
		} else {
			for j := 0; j < num; j++ {
				result = append(result, strconv.Itoa(i/2))
			}
		}
	}
	return result
}

func checksum(lines []string) int {
	total := 0
	j := 0
	continua := true
	for i := 0; continua; i++ {
		for lines[len(lines)-1-i-j] == "." {
			j++
			if i >= len(lines)-j-i {
				break
			}
		}
		fmt.Println(lines)
		for k, char := range lines {
			if char == "." {
				if len(lines)-1-i-j < k {
					continua = false
				} else {
					lines[k], lines[len(lines)-1-i-j] = lines[len(lines)-1-i-j], lines[k]
				}
				break
			}
		}
	}
	for i, val := range lines {
		num, _ := strconv.Atoi(string(val))
		total += num * i
	}
	return total
}

func part2(lines []string) int {
	total := 0

	for i := len(lines) - 1; i > 0; i-- {
		if lines[i] != "." {
			size := 0
			aux := i
			for j := i - 1; j > 0; j-- {
				if lines[j] != lines[i] {
					size = i - j
					i -= size
					i++
					break
				}
			}
			//	fmt.Println("size", size, lines[i], i, aux)
			space := 0
			for k := 0; k < aux-size; k++ {
				if lines[k] == "." {
					space++
				} else {
					space = 0
				}

				if space == size && size > 0 {
					//fmt.Println(size-1, k)
					for l := size - 1; l >= 0; l-- {
						lines[k-l], lines[aux-l] = lines[aux-l], lines[k-l]
					}
					break
				}
			}

		}
		//fmt.Println(i, lines)
	}

	for i, val := range lines {
		num, _ := strconv.Atoi(string(val))
		//	fmt.Println(i, val, num)
		total += num * i
	}
	//fmt.Println(lines)
	return total
}

func load() ([]byte, error) {
	const fileName = "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}
