package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const (
	one   = "one"
	two   = "two"
	three = "three"
	four  = "four"
	five  = "five"
	six   = "six"
	seven = "seven"
	eight = "eight"
	nine  = "nine"
)

var numbers = []string{one, two, three, four, five, six, seven, eight, nine}
var numMap = map[string]int{
	one:   1,
	two:   2,
	three: 3,
	four:  4,
	five:  5,
	six:   6,
	seven: 7,
	eight: 8,
	nine:  9,
}

func getFirstNumber(text string) int {
	for index, char := range text {
		if unicode.IsDigit(char) {
			return int(rune(text[index]) - '0')
		} else {
			for _, number := range numbers {
				if strings.Contains(text[:index+1], number) {
					return numMap[number]
				}
			}
		}
	}
	return -1
}

func getLastNumber(text string) int {
	for i := len(text) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(text[i])) {
			return int(rune(text[i]) - '0')
		} else {
			for _, number := range numbers {
				if strings.Contains(text[i:], number) {
					return numMap[number]
				}
			}
		}
	}
	return -1
}

func getNumbers(text string, numChan chan int) {
	first := getFirstNumber(text)
	last := getLastNumber(text)
	result := first*10 + last
	numChan <- result
}

func main() {
	file, _ := os.Open("input")
	reader := bufio.NewScanner(file)
	var total int
	numChan := make(chan int)

	for reader.Scan() {
		go getNumbers(reader.Text(), numChan)
		nums := <-numChan
		total += nums
	}

	fmt.Println(total)
}
