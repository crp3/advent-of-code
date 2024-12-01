package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func difference(listOne, listTwo []int) int {
	diff := 0
	for index := range listOne {
		if listOne[index] > listTwo[index] {
			diff += listOne[index] - listTwo[index]
		} else {
			diff += listTwo[index] - listOne[index]
		}
	}
	return diff
}

func parseFirstNumber(str string) int {
	numBuffer := make([]byte, 0)
	for index := 0; str[index] != ' '; index++ {
		numBuffer = append(numBuffer, str[index])
	}
	val, _ := strconv.Atoi(string(numBuffer))
	return int(val)
}

func parseLastNumber(str string) int {
	numBuffer := make([]byte, 0)
	index := 0
	for local := index; str[local] != ' '; local++ {
		index = local
	}
	index++
	for local := index; str[local] == ' '; local++ {
		index = local
	}
	index++

	for local := index; local < len(str); local++ {
		numBuffer = append(numBuffer, str[local])
	}

	val, _ := strconv.Atoi(string(numBuffer))
	return int(val)
}

func similarityScore(listOne, listTwo []int) int {
	counts := make(map[int]int)
	for _, num := range listTwo {
		counts[num]++
	}

	score := 0
	for _, num := range listOne {
		score += num * counts[num]
	}
	return score
}

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewScanner(file)
	listOne := make([]int, 0)
	listTwo := make([]int, 0)
	for reader.Scan() {
		line := reader.Text()
		listOne = append(listOne, parseFirstNumber(line))
		listTwo = append(listTwo, parseLastNumber(line))
	}
	sort.Slice(listOne, func(i, j int) bool { return listOne[i] < listOne[j] })
	sort.Slice(listTwo, func(i, j int) bool { return listTwo[i] < listTwo[j] })
	fmt.Println(difference(listOne, listTwo))
	fmt.Println(similarityScore(listOne, listTwo))
}
