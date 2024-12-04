package main

import (
	"fmt"
	"os"
	"strings"
)

type CharMatrix [][]rune

func main() {
	lines, err := load()
	if err != nil {
		fmt.Println(err)
		return
	}
	puzzle := FromString(lines)

	fmt.Println(makeDir(puzzle))
}

func makeDir(puzzle CharMatrix) int {
	result := 0
	for i := 1; i < len(puzzle)-1; i++ {
		for j := 1; j < len(puzzle[0])-1; j++ {
			if puzzle[i][j] == 65 {
				if checkMas1(puzzle, i, j) && checkMas2(puzzle, i, j) {
					result++
				}
			}
		}
	}
	return result
}

func checkMas1(puzzle CharMatrix, i, j int) bool {
	if (puzzle[i-1][j-1] == 77 && puzzle[i+1][j+1] == 83) || (puzzle[i-1][j-1] == 83 && puzzle[i+1][j+1] == 77) {
		return true
	}
	return false
}

func checkMas2(puzzle CharMatrix, i, j int) bool {
	if (puzzle[i-1][j+1] == 77 && puzzle[i+1][j-1] == 83) || (puzzle[i-1][j+1] == 83 && puzzle[i+1][j-1] == 77) {
		return true
	}
	return false
}

func NewCharMatrix(rows, cols int) CharMatrix {
	matrix := make(CharMatrix, rows)
	for i := range matrix {
		matrix[i] = make([]rune, cols)
	}
	return matrix
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

func FromString(input string) CharMatrix {
	lines := strings.Split(input, "\n")
	matrix := NewCharMatrix(len(lines), len(lines[0]))

	for i, line := range lines {
		for j, char := range line {
			matrix[i][j] = char
		}
	}

	return matrix
}
