package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func processMul(operand1, operand2 []string) int {
	op1, _ := strconv.Atoi(operand1[0])
	op2, _ := strconv.Atoi(operand2[0])
	return op1 * op2
}

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewScanner(file)
	enabled := true
	p1 := 0
	p2 := 0
	re := regexp.MustCompile(`\d+`)
	nums := []byte{',', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for reader.Scan() {
		line := reader.Text()
		for i := range line {
			if strings.HasPrefix(line[i:], "don't()") {
				enabled = false
			} else if strings.HasPrefix(line[i:], "do()") {
				enabled = true
			} else if strings.HasPrefix(line[i:], "mul(") {
				j := i + 4
				valid := true
				for line[j] != ')' {
					j++
					if !slices.Contains(nums, line[j-1]) {
						valid = false
						break
					}
				}

				if valid {
					matches := re.FindAllStringSubmatch(line[i:j+1], -1)
					fmt.Println(line[i : j+1])
					if len(matches) > 1 {
						x, y := matches[0], matches[1]
						fmt.Println(x, y)
						if enabled {
							p2 += processMul(x, y)
						}
						p1 += processMul(x, y)

					}
				}

			}

		}
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
