package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func convertNumList(nums string) []int {
	numList := make([]int, 0)
	splitString := strings.Split(nums, " ")
	for _, num := range splitString {
		conv, _ := strconv.Atoi(num)
		numList = append(numList, conv)
	}
	return numList
}

func isValid(target int, numList []int) bool {
	if len(numList) == 1 {
		return target == numList[0]
	}
	op1 := isValid(target, slices.Concat([]int{numList[0] + numList[1]}, numList[2:]))
	op2 := isValid(target, slices.Concat([]int{numList[0] * numList[1]}, numList[2:]))
	concatInt, _ := strconv.Atoi(strconv.Itoa(numList[0]) + strconv.Itoa(numList[1]))
	op3 := isValid(target, slices.Concat([]int{concatInt}, numList[2:]))
	return op1 || op2 || op3
}

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewScanner(file)
	total := 0
	for reader.Scan() {
		line := reader.Text()
		splitLine := strings.Split(line, ": ")
		target, _ := strconv.Atoi(splitLine[0])
		numList := convertNumList(splitLine[1])
		isValid := isValid(target, numList)
		if isValid {
			total += target
		}
		fmt.Println(target, numList, isValid)
	}
	fmt.Println(total)
}
