package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validateSubsets(subSets string) (bool, int) {
	maxCount := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	minimumColor := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	valid := true
	sets := strings.Split(subSets, "; ")
	for _, set := range sets {
		die := strings.Split(set, ", ")
		for _, cubes := range die {
			countColor := strings.Split(cubes, " ")
			countStr, color := countColor[0], countColor[1]
			count, _ := strconv.Atoi(countStr)
			if count > maxCount[color] {
				valid = false
			}
			if count > minimumColor[color] {
				minimumColor[color] = count
			}
		}
	}
	powerColor := 1
	for _, val := range minimumColor {
		powerColor *= val
	}
	return valid, powerColor
}

func validateGame(readLine string) (int, int) {
	split := strings.Split(readLine, ":")
	gameIDStr, subSets := split[0], split[1]
	id, _ := strconv.Atoi(strings.Split(gameIDStr, " ")[1])
	valid, power := validateSubsets(strings.Trim(subSets, " "))
	if valid {
		return id, power
	}
	return 0, power
}

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewScanner(file)
	var gameIDResult int
	var powerResult int

	for reader.Scan() {
		id, power := validateGame(reader.Text())
		gameIDResult += id
		powerResult += power
	}
	fmt.Println("gameID sums: ", gameIDResult)
	fmt.Println("power sums: ", powerResult)
}
