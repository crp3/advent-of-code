package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type indexPair struct {
	row int
	col int
}

type engineNumber struct {
	indices []indexPair
	number  int
}

type symbol struct {
	exists bool
	symbol rune
}

var (
	neighbors = []indexPair{
		{0, 1},
		{1, 0},
		{1, 1},
		{0, -1},
		{-1, 0},
		{-1, -1},
		{-1, 1},
		{1, -1},
	}
)

func updateSymbolSet(symbolSet map[indexPair]symbol, line string, rowNum int) {
	for colNum, char := range line {
		if char != '.' && !unicode.IsDigit(char) {
			symbolSet[indexPair{rowNum, colNum}] = symbol{true, char}
		}
	}
}

func updateEngineNumbers(engineNumbers *[]engineNumber, line string, rowNum int) {
	numberBuffer := make([]rune, 0)
	for colNum := range line {
		char := line[colNum]
		if unicode.IsDigit(rune(char)) {
			numberBuffer = append(numberBuffer, rune(char))
		} else {
			if len(numberBuffer) > 0 {
				intg, _ := strconv.Atoi(string(numberBuffer))
				indices := make([]indexPair, 0)
				for colOffset := range numberBuffer {
					indices = append(indices, indexPair{
						row: rowNum,
						col: colNum - (colOffset + 1),
					})
				}
				*engineNumbers = append(*engineNumbers, engineNumber{
					number:  intg,
					indices: indices,
				})
				numberBuffer = make([]rune, 0)
			}
		}
	}
	if len(numberBuffer) > 0 {
		intg, _ := strconv.Atoi(string(numberBuffer))
		indices := make([]indexPair, 0)
		for colOffset := range numberBuffer {
			indices = append(indices, indexPair{
				row: rowNum,
				col: len(line) - (colOffset + 1),
			})
		}
		*engineNumbers = append(*engineNumbers, engineNumber{
			number:  intg,
			indices: indices,
		})
	}
}

func sumNumbers(engineNumbers []engineNumber, symbolSet map[indexPair]symbol) int {
	total := 0
	for _, number := range engineNumbers {
		if hasSymbolNeighbors(number, symbolSet) {
			total += number.number
		}
	}

	return total
}

func hasSymbolNeighbors(engineNumber engineNumber, symbolSet map[indexPair]symbol) bool {
	for _, neighbor := range neighbors {
		for _, index := range engineNumber.indices {
			newRow := neighbor.row + index.row
			newCol := neighbor.col + index.col
			if symbolSet[indexPair{newRow, newCol}].exists {
				return true
			}
		}
	}

	return false
}

func getGearRatios(numbers []engineNumber, symbolSet map[indexPair]symbol) int64 {
	symbolMap := make(map[indexPair]map[string]bool)
	var result int64

	for _, num := range numbers {
		for _, index := range num.indices {
			for _, neighbor := range neighbors {
				newRow := neighbor.row + index.row
				newCol := neighbor.col + index.col
				key := indexPair{newRow, newCol}
				if symbolSet[key].exists && symbolSet[key].symbol == '*' {
					if symbolMap[key] == nil {
						symbolMap[key] = make(map[string]bool)
					}
					symbolMap[key][fmt.Sprintf("%d,%d:%d", num.indices[0].row, num.indices[0].col, num.number)] = true
				}
			}
		}
	}
	for _, map_ := range symbolMap {
		var localResult int64 = 1
		if len(map_) == 2 {
			for val := range map_ {
				intVal, _ := strconv.Atoi(strings.Split(val, ":")[1])
				localResult *= int64(intVal)
			}

			result += localResult
		}
	}
	return result
}

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewScanner(file)
	symbolSet := make(map[indexPair]symbol)
	numbers := make([]engineNumber, 0)
	row := 0

	for reader.Scan() {
		line := reader.Text()
		updateSymbolSet(symbolSet, line, row)
		updateEngineNumbers(&numbers, line, row)
		row += 1
	}
	fmt.Println("ratios: ", getGearRatios(numbers, symbolSet))
	fmt.Println("sum: ", sumNumbers(numbers, symbolSet))
}
