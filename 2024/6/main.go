package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
	fmt.Println("")
}

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewScanner(file)
	grid := make([][]rune, 0)
	initX, initY := 0, 0
	gridIndex := 0
	for reader.Scan() {
		line := reader.Text()
		grid = append(grid, []rune(line))
		li := strings.LastIndex(line, "^")
		if li != -1 {
			initY = li
			initX = gridIndex
		}
		gridIndex += 1
	}
	currentX, currentY := initX, initY
	total := 0
	for currentX < len(grid) && currentX >= 0 && currentY < len(grid[0]) && currentY >= 0 {
		printGrid(grid)
		switch grid[currentX][currentY] {
		case '^':
			if currentX == 0 {
				grid[currentX][currentY] = 'X'
				currentX--
				continue
			}
			if grid[currentX-1][currentY] == '#' {
				grid[currentX][currentY] = '>'
				continue
			}
			if grid[currentX-1][currentY] == '.' || grid[currentX-1][currentY] == 'X' {
				grid[currentX][currentY] = 'X'
				grid[currentX-1][currentY] = '^'
				currentX--
			}
		case '>':
			if currentY == len(grid[0])-1 {
				grid[currentX][currentY] = 'X'
				currentY++
				continue
			}
			if grid[currentX][currentY+1] == '#' {
				grid[currentX][currentY] = 'v'
				continue
			}
			if grid[currentX][currentY+1] == '.' || grid[currentX][currentY+1] == 'X' {
				grid[currentX][currentY] = 'X'
				grid[currentX][currentY+1] = '>'
				currentY++
			}
		case 'v':
			if currentX == len(grid)-1 {
				grid[currentX][currentY] = 'X'
				currentX++
				continue
			}
			if grid[currentX+1][currentY] == '#' {
				grid[currentX][currentY] = '<'
				continue
			}
			if grid[currentX+1][currentY] == '.' || grid[currentX+1][currentY] == 'X' {
				grid[currentX][currentY] = 'X'
				grid[currentX+1][currentY] = 'v'
				currentX++
			}
		case '<':
			if currentY == 0 {
				grid[currentX][currentY] = 'X'
				currentY--
				continue
			}
			if grid[currentX][currentY-1] == '#' {
				grid[currentX][currentY] = '^'
				continue
			}
			if grid[currentX][currentY-1] == '.' || grid[currentX][currentY-1] == 'X' {
				grid[currentX][currentY] = 'X'
				grid[currentX][currentY-1] = '<'
				currentY--
			}
		}
	}

	printGrid(grid)
	for _, line := range grid {
		for _, rn := range line {
			if rn == 'X' {
				total++
			}
		}
	}
	fmt.Println(total)
}
