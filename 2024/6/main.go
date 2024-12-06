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

type CoordPlusDirection struct {
	X         int
	Y         int
	Direction int
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
	totalSecond := 0
	for currentX < len(grid) && currentX >= 0 && currentY < len(grid[0]) && currentY >= 0 {
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

	for obstacleX := range grid {
		for obstacleY := range grid[0] {
			currentX, currentY := initX, initY
			direction := 0
			seen := make(map[CoordPlusDirection]bool)
			for true {
				coordPlus := CoordPlusDirection{
					currentX,
					currentY,
					direction,
				}
				if seen[coordPlus] {
					totalSecond += 1
					break
				}
				seen[coordPlus] = true
				newDirection := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}[direction]
				newX, newY := currentX+newDirection[0], currentY+newDirection[1]
				if newX < 0 || newX >= len(grid) || newY < 0 || newY >= len(grid) {
					break
				}
				if grid[newX][newY] == '#' || newX == obstacleX && newY == obstacleY {
					direction = (direction + 1) % 4
				} else {
					currentX, currentY = newX, newY
				}
			}
		}
	}
	for _, line := range grid {
		for _, rn := range line {
			if rn == 'X' {
				total++
			}
		}
	}
	fmt.Println(total)
	fmt.Println(totalSecond)
}
