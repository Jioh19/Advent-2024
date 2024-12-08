package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	lines, err := load()
	if err != nil {
		fmt.Println(err)
		return
	}

	nodes := make([][]int, len(lines))
	for i := 0; i < len(lines); i++ {
		node := make([]int, len(lines[0]))
		nodes[i] = node
		fmt.Println(nodes[i])
	}
	antenna := make(map[byte]bool)
	antenna[46] = true
	fmt.Println(part1(nodes, lines, antenna))
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

func part1(nodes [][]int, lines []string, antenna map[byte]bool) int {
	total := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if checkAntenna(antenna, lines[i][j]) {
				continue
			}
			listAntenna := findNext(lines, lines[i][j])
			fmt.Println(listAntenna)
			antenna[lines[i][j]] = true
			writeNodes2(listAntenna, nodes)
		}
	}
	for _, node := range nodes {
		fmt.Println(node)
		for _, val := range node {
			if val == 1 {
				total++
			}
		}
	}
	return total
}

func writeNodes(listAntenna [][2]int, nodes [][]int) {
	for i := 0; i < len(listAntenna)-1; i++ {
		for j := i + 1; j < len(listAntenna); j++ {
			di := listAntenna[i][0] - listAntenna[j][0]
			dj := listAntenna[i][1] - listAntenna[j][1]
			if listAntenna[i][0]+di < 0 || listAntenna[i][0]+di >= len(nodes) || listAntenna[i][1]+dj < 0 || listAntenna[i][1]+dj >= len(nodes[0]) {
			} else {
				nodes[listAntenna[i][0]+di][listAntenna[i][1]+dj] = 1
			}
			if listAntenna[j][0]-di < 0 || listAntenna[j][0]-di >= len(nodes) || listAntenna[j][1]-dj < 0 || listAntenna[j][1]-dj >= len(nodes[0]) {
			} else {
				nodes[listAntenna[j][0]-di][listAntenna[j][1]-dj] = 1
			}

		}
	}
}

func writeNodes2(listAntenna [][2]int, nodes [][]int) {
	fmt.Println("ENtramos")
	max := len(listAntenna)
	for i := 0; i < max-1; i++ {
		for j := i + 1; j < max; j++ {

			di := listAntenna[i][0] - listAntenna[j][0]
			dj := listAntenna[i][1] - listAntenna[j][1]
			x := listAntenna[i][0]
			y := listAntenna[i][1]
			nodes[x][y] = 1
			for x+di >= 0 && x+di < len(nodes) && y+dj >= 0 && y+dj < len(nodes[0]) {
				x += di
				y += dj
				nodes[x][y] = 1
			}
			x = listAntenna[j][0]
			y = listAntenna[j][1]
			nodes[x][y] = 1
			for x-di >= 0 && x-di < len(nodes) && y-dj >= 0 && y-dj < len(nodes[0]) {
				x -= di
				y -= dj
				nodes[x][y] = 1
			}
		}
	}

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findNext(lines []string, val byte) [][2]int {
	listAntenna := [][2]int{}
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == val {
				point := [2]int{i, j}
				listAntenna = append(listAntenna, point)
			}
		}
	}
	return listAntenna
}

func checkAntenna(antenna map[byte]bool, val byte) bool {
	_, exists := antenna[val]
	if exists {
		return true
	} else {
		return false
	}
}

func clear(matrix [][]int) {
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = 0
		}
	}
}
