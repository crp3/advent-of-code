package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewScanner(file)
	occurrences := 0
	xOccurrences := 0
	grid := make([]string, 0)
	for reader.Scan() {
		line := reader.Text()
		grid = append(grid, line)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'X' {
				//vertical
				if i < len(grid)-3 {
					if grid[i+1][j] == 'M' && grid[i+2][j] == 'A' && grid[i+3][j] == 'S' {
						occurrences += 1
					}
				}
				if i >= 3 {
					if grid[i-1][j] == 'M' && grid[i-2][j] == 'A' && grid[i-3][j] == 'S' {
						occurrences += 1
					}
				}
				//horizontal
				if j < len(grid[i])-3 {
					if grid[i][j+1] == 'M' && grid[i][j+2] == 'A' && grid[i][j+3] == 'S' {
						occurrences += 1
					}
				}
				if j >= 3 {
					if grid[i][j-1] == 'M' && grid[i][j-2] == 'A' && grid[i][j-3] == 'S' {
						occurrences += 1
					}
				}
				//diagonal
				if j >= 3 && i >= 3 {
					if grid[i-1][j-1] == 'M' && grid[i-2][j-2] == 'A' && grid[i-3][j-3] == 'S' {
						occurrences += 1
					}
				}
				if j >= 3 && i < len(grid)-3 {
					if grid[i+1][j-1] == 'M' && grid[i+2][j-2] == 'A' && grid[i+3][j-3] == 'S' {
						occurrences += 1
					}
				}
				if j < len(grid[i])-3 && i < len(grid)-3 {
					if grid[i+1][j+1] == 'M' && grid[i+2][j+2] == 'A' && grid[i+3][j+3] == 'S' {
						occurrences += 1
					}
				}
				if j < len(grid[i])-3 && i >= 3 {
					if grid[i-1][j+1] == 'M' && grid[i-2][j+2] == 'A' && grid[i-3][j+3] == 'S' {
						occurrences += 1
					}
				}
			}
			if grid[i][j] == 'A' {
				if i > 0 && i < len(grid)-1 && j > 0 && j < len(grid)-1 {
					if grid[i-1][j-1] == 'M' && grid[i+1][j+1] == 'S' {
						if grid[i-1][j+1] == 'M' && grid[i+1][j-1] == 'S' {
							xOccurrences += 1
						}
						if grid[i-1][j+1] == 'S' && grid[i+1][j-1] == 'M' {
							xOccurrences += 1
						}
					}
					if grid[i-1][j-1] == 'S' && grid[i+1][j+1] == 'M' {
						if grid[i-1][j+1] == 'M' && grid[i+1][j-1] == 'S' {
							xOccurrences += 1
						}
						if grid[i-1][j+1] == 'S' && grid[i+1][j-1] == 'M' {
							xOccurrences += 1
						}
					}
				}
			}
		}
		fmt.Println("")
	}
	fmt.Println(occurrences)
	fmt.Println(xOccurrences)
}
