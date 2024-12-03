package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseLine(line string) []int {
	numList := make([]int, 0)
	index := 0
	for index < len(line) {
		numBuffer := make([]byte, 0)
		localIndex := index
		for localIndex < len(line) && line[localIndex] != ' ' {
			numBuffer = append(numBuffer, line[localIndex])
			localIndex++
		}
		val, _ := strconv.Atoi(string(numBuffer))
		numList = append(numList, val)
		index = localIndex + 1
	}
	return numList
}
func isAscending(numList []int) bool {
	if numList[0] < numList[len(numList)-1] {
		return true
	}
	return false
}

func isSafe(numList []int) bool {
	isAscending := isAscending(numList)
	for index := 0; index < len(numList)-1; index++ {
		if isAscending {
			if numList[index+1]-numList[index] > 3 || numList[index+1]-numList[index] < 1 {
				return false
			}
		} else {
			if numList[index]-numList[index+1] > 3 || numList[index]-numList[index+1] < 1 {
				return false
			}
		}
	}
	return true
}

func createNewNumList(numList []int, removeIndex int) []int {
	newNumList := make([]int, 0)
	for index := range numList {
		if index != removeIndex {
			newNumList = append(newNumList, numList[index])
		}
	}
	return newNumList
}

func isSafeRampened(numList []int) bool {
	isAscending := isAscending(numList)
	valid := false
	fmt.Println(numList)
	for removeIndex := range numList {
		newNumList := createNewNumList(numList, removeIndex)
		localValid := true
		fmt.Println(newNumList)
		for index := 0; index < len(newNumList)-1; index++ {
			if isAscending {
				diff := newNumList[index+1] - newNumList[index]
				localValid = localValid && (diff <= 3 && diff >= 1)
			} else {
				diff := newNumList[index] - newNumList[index+1]
				localValid = localValid && (diff <= 3 && diff >= 1)
			}
		}
		valid = valid || localValid
		if valid {
			fmt.Println(valid)
			return true
		}
	}
	fmt.Println(valid)
	return false
}

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewScanner(file)
	safeCount := 0
	safeCountRampened := 0
	for reader.Scan() {
		line := reader.Text()
		numList := parseLine(line)
		if isSafe(numList) {
			safeCount += 1
		}
		if isSafeRampened(numList) {
			safeCountRampened += 1
		}
	}
	fmt.Printf("%d safe lines\n", safeCount)
	fmt.Printf("%d rampened safe lines\n", safeCountRampened)
}
