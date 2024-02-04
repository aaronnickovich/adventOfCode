package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func parseRowsInputMock() []string {
	var input []string = []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598.."}
	return input
}

func isNumber(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

func isSymbolAdjacentLeft(input []string, i int, j int) bool {
	if j <= 0 {
		return false
	}
	if input[i][j-1] != '.' && !isNumber(input[i][j-1]) {
		return true
	}
	return false
}

func isSymbolAdjacentRight(input []string, i int, j int) bool {
	if j >= len(input[i])-1 {
		return false
	}
	if input[i][j+1] != '.' && !isNumber(input[i][j+1]) {
		return true
	}
	return false
}

func isSymbolAdjacentUp(input []string, i int, j int) bool {
	if i <= 0 {
		return false
	}
	if input[i-1][j] != '.' && !isNumber(input[i-1][j]) {
		return true
	}
	return false
}

func isSymbolAdjacentDown(input []string, i int, j int) bool {
	if i >= len(input)-1 {
		return false
	}
	if input[i+1][j] != '.' && !isNumber(input[i+1][j]) {
		return true
	}
	return false
}

func isSymbolAdjacentUpperLeft(input []string, i int, j int) bool {
	if i <= 0 || j <= 0 {
		return false
	}
	if input[i-1][j-1] != '.' && !isNumber(input[i-1][j-1]) {
		return true
	}
	return false
}

func isSymbolAdjacentUpperRight(input []string, i int, j int) bool {
	if i <= 0 || j >= len(input[i])-1 {
		return false
	}
	if input[i-1][j+1] != '.' && !isNumber(input[i-1][j+1]) {
		return true
	}
	return false
}

func isSymbolAdjacentLowerLeft(input []string, i int, j int) bool {
	if i >= len(input)-1 || j <= 0 {
		return false
	}
	if input[i+1][j-1] != '.' && !isNumber(input[i+1][j-1]) {
		return true
	}
	return false
}

func isSymbolAdjacentLowerRight(input []string, i int, j int) bool {
	if i >= len(input)-1 || j >= len(input[i])-1 {
		return false
	}
	if input[i+1][j+1] != '.' && !isNumber(input[i+1][j+1]) {
		return true
	}
	return false
}

func run(input []string) uint64 {
	var finalSum uint64 = 0
	for i := 0; i < len(input); i++ {
		var foundAdjacentSymbol bool = false
		var numberIndex uint8 = 0
		var partNumber uint64 = 0
		for j := len(input[i]) - 1; j >= 0; j-- {
			if isNumber(input[i][j]) {
				number, _ := strconv.Atoi(string(input[i][j]))
				partNumber = partNumber + uint64(number)*uint64(math.Pow10(int(numberIndex)))
				numberIndex++
				if foundAdjacentSymbol {
					if j <= 0 {
						fmt.Printf("number: %d", partNumber)
						finalSum += partNumber
						foundAdjacentSymbol = false
					}
					continue
				}
				if isSymbolAdjacentLeft(input, i, j) ||
					isSymbolAdjacentRight(input, i, j) ||
					isSymbolAdjacentUp(input, i, j) ||
					isSymbolAdjacentDown(input, i, j) ||
					isSymbolAdjacentUpperLeft(input, i, j) ||
					isSymbolAdjacentUpperRight(input, i, j) ||
					isSymbolAdjacentLowerLeft(input, i, j) ||
					isSymbolAdjacentLowerRight(input, i, j) {
					foundAdjacentSymbol = true
				}
			} else {
				if foundAdjacentSymbol {
					fmt.Printf("number: %d", partNumber)
					finalSum += partNumber
					foundAdjacentSymbol = false
				}
				numberIndex = 0
				partNumber = 0
			}
		}
	}
	return finalSum
}

func main() {
	var input []string;
	file, err := os.Open("../input.txt");
	if err != nil {
	  panic(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
	  text := scanner.Text()

	  input = append(input, text);
	}

	if err := scanner.Err(); err != nil {
	  panic(err);
	}

	sum := run(input);
	fmt.Print("\nsum: ", sum)
}
