package main

import (
	"fmt"
	"os"
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
	rows, cols := len(lines), len(lines[0])
	puzzle := make([][]int, rows)
	for i := range puzzle {
		puzzle[i] = make([]int, cols)
	}
	i, j := getStart(lines)
	puzzle[i][j] = 1
	part1(lines, puzzle, i, j, 94)
	fmt.Println("El total de part1 es: ", solve1(puzzle))
	fmt.Println("El total de part2 es: ", solve2(lines, puzzle, i, j, 94))

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

func getStart(lines []string) (int, int) {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == 94 {
				return i, j
			}
		}
	}
	return -1, -1
}

func part1(lines []string, puzzle [][]int, i, j int, dir byte) {
	for i := 0; i < len(puzzle); i++ {
	}
	if next(lines, i, j, dir) == -1 {
		puzzle[i][j] = 1

		return
	} else if next(lines, i, j, dir) == 0 {
		if dir == 94 {
			dir = 62
		} else if dir == 62 {
			dir = 118
		} else if dir == 118 {
			dir = 60
		} else if dir == 60 {
			dir = 94
		}

		part1(lines, puzzle, i, j, dir)
	} else if next(lines, i, j, dir) == 1 {
		di, dj := nextDir(dir)
		puzzle[i][j] = 1
		part1(lines, puzzle, i+di, j+dj, dir)
	}
}

func solve1(puzzle [][]int) int {
	total := 0
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[0]); j++ {
			total += puzzle[i][j]
		}
	}
	return total
}

func next(lines []string, i, j int, dir byte) int {
	di, dj := nextDir(dir)
	if i+di < 0 || i+di >= len(lines) || j+dj < 0 || j+dj >= len(lines[0]) {
		return -1
	} else if lines[i+di][j+dj] == 35 {
		return 0
	}
	return 1
}

func next2(lines []string, puzzle [][]int, i, j int, dir byte) int {
	di, dj := nextDir(dir)
	if i+di < 0 || i+di >= len(lines) || j+dj < 0 || j+dj >= len(lines[0]) {
		return -1
	} else if lines[i+di][j+dj] == 35 || puzzle[i+di][j+dj] == 35 {
		return 0
	}
	return 1
}

func nextDir(dir byte) (int, int) {
	if dir == 94 {
		return -1, 0
	} else if dir == 62 {
		return 0, 1
	} else if dir == 118 {
		return 1, 0
	} else {
		return 0, -1
	}
}

func solve2(lines []string, puzzle [][]int, i, j int, dir byte) int {
	total := 0
	for di := 0; di < len(lines); di++ {
		for dj := 0; dj < len(lines[0]); dj++ {
			puzzle[di][dj] = 35
			part2(lines, puzzle, i, j, dir, &total)
			clean(puzzle)
		}
	}
	return total
}

func clean(puzzle [][]int) {
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[0]); j++ {
			puzzle[i][j] = 0
		}
	}
}
func part2(lines []string, puzzle [][]int, i, j int, dir byte, total *int) {
	for i := 0; i < len(puzzle); i++ {
	}
	if next2(lines, puzzle, i, j, dir) == -1 {
		puzzle[i][j] = int(dir)
		return
	} else if next2(lines, puzzle, i, j, dir) == 0 {
		if dir == 94 {
			dir = 62
		} else if dir == 62 {
			dir = 118
		} else if dir == 118 {
			dir = 60
		} else if dir == 60 {
			dir = 94
		}

		part2(lines, puzzle, i, j, dir, total)
	} else if next2(lines, puzzle, i, j, dir) == 1 {
		di, dj := nextDir(dir)
		if puzzle[i][j] == int(dir) {
			*total++
			return
		}
		puzzle[i][j] = int(dir)
		part2(lines, puzzle, i+di, j+dj, dir, total)
	}
}
