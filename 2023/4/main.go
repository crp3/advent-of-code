package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getLinePoints(line string, copiesMap map[int]int) int {
	splitLine := strings.Split(line, "|")
	firstPart := strings.Split(splitLine[0], ":")
	winningCardsString := firstPart[1]
	winningCards := getWinningCardsSet(strings.Split(strings.Trim(winningCardsString, " "), " "))
	handCards := strings.Split(strings.Trim(splitLine[1], " "), " ")
	splitGameId := strings.Split(strings.Trim(firstPart[0], " "), " ")
	gameID, _ := strconv.Atoi(splitGameId[len(splitGameId)-1])
	copiesMap[gameID] += 1
	winningCount := 0
	for _, card := range handCards {
		if winningCards[card] {
			winningCount += 1
		}
	}
	minPossible := 1
	if copiesMap[gameID] > minPossible {
		minPossible = copiesMap[gameID]
	}
	for count := 0; count < winningCount; count += 1 {
		copiesMap[gameID+count+1] += minPossible
	}

	if winningCount == 0 {
		return 0
	}
	result := 1
	for count := winningCount - 1; count > 0; count -= 1 {
		result *= 2
	}
	return result
}

func getWinningCardsSet(winningCards []string) map[string]bool {
	set := make(map[string]bool)
	for _, winningCard := range winningCards {
		if winningCard != " " && winningCard != "" {
			set[winningCard] = true
		}
	}
	return set
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	points := 0
	copies := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		points += getLinePoints(line, copies)
	}
	total := 0
	for _, val := range copies {
		total += val
	}

	fmt.Println(copies)
	fmt.Println("total copies: ", total)
	fmt.Println("points: ", points)
}
