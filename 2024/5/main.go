package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewScanner(file)
	result := 0
	secondPart := 0
	graph := make(map[int]map[int]bool)
	incorrectlyOrdered := make([][]int, 0)
	for reader.Scan() {
		line := reader.Text()
		pairSplit := strings.Split(line, "|")
		if len(pairSplit) > 1 {
			x, _ := strconv.Atoi(pairSplit[0])
			y, _ := strconv.Atoi(pairSplit[1])
			if graph[y] == nil {
				graph[y] = make(map[int]bool, 0)
			}
			graph[y][x] = true
		}
		queueSplit := strings.Split(line, ",")
		if len(queueSplit) > 1 {
			queue := make([]int, 0)
			for _, num := range queueSplit {
				convert, _ := strconv.Atoi(num)
				queue = append(queue, convert)
			}
			ok := true
			for i, x := range queue {
				for j, y := range queue {
					if i < j && graph[x][y] {
						ok = false
					}
				}
			}
			if ok {
				result += queue[len(queue)/2]
			} else {
				incorrectlyOrdered = append(incorrectlyOrdered, queue)
			}
		}
	}
	for _, queue := range incorrectlyOrdered {
		slices.SortFunc(queue, func(x, y int) int {
			if graph[x][y] {
				return -1
			}
			return 1
		})
		secondPart += queue[len(queue)/2]
	}

	fmt.Printf("response: %d\n", result)
	fmt.Printf("secondPart: %d\n", secondPart)

}
