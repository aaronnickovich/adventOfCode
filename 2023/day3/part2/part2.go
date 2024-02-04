package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func isGear(char byte) bool {
	if char == '*' {
		return true
	}
	return false
}

func isNumber(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

func isNumberAdjacentLeft(input []string, i int, j int) bool {
	if j <= 0 {
		return false
	}
	if isNumber(input[i][j-1]) {
		return true
	}
	return false
}

func isNumberAdjacentRight(input []string, i int, j int) bool {
	if j >= len(input[i])-1 {
		return false
	}
	if isNumber(input[i][j+1]) {
		return true
	}
	return false
}

func isNumberAdjacentUp(input []string, i int, j int) bool {
	if i <= 0 {
		return false
	}
	if isNumber(input[i-1][j]) {
		return true
	}
	return false
}

func isNumberAdjacentUpperLeft(input []string, i int, j int) bool {
	if i <= 0 || j <= 0 {
		return false
	}
	if isNumber(input[i-1][j-1]) {
		return true
	}
	return false
}

func isNumberAdjacentUpperRight(input []string, i int, j int) bool {
	if i <= 0 || j >= len(input[i])-1 {
		return false
	}
	if isNumber(input[i-1][j+1]) {
		return true
	}
	return false
}

func isNumberAdjacentDown(input []string, i int, j int) bool {
	if i >= len(input)-1 {
		return false
	}
	if isNumber(input[i+1][j]) {
		return true
	}
	return false
}

func isNumberAdjacentLowerLeft(input []string, i int, j int) bool {
	if i >= len(input)-1 || j <= 0 {
		return false
	}
	if isNumber(input[i+1][j-1]) {
		return true
	}
	return false
}

func isNumberAdjacentLowerRight(input []string, i int, j int) bool {
	if i >= len(input)-1 || j >= len(input[i])-1 {
		return false
	}
	if isNumber(input[i+1][j+1]) {
		return true
	}
	return false
}

type ScanStart struct {
	row    int
	column int
}

func run(input []string) uint64 {
	var finalSum uint64 = 0
	for i := 0; i < len(input); i++ {
		for j := len(input[i]) - 1; j >= 0; j-- {
			if isGear(input[i][j]) {
				var scanStarts []ScanStart
				var checkUpperLeft bool = true
				var checkUp bool = true
				var checkDown bool = true
				var checkLowerLeft bool = true

				if isNumberAdjacentLeft(input, i, j) {
					scanStarts = append(scanStarts, ScanStart{row: int(i), column: int(j - 1)})
					fmt.Printf("found left adjacent at i: %d, j %d\n", i, j)
				}
				if isNumberAdjacentRight(input, i, j) {
					scanStarts = append(scanStarts, ScanStart{row: int(i), column: int(j + 1)})
					fmt.Printf("found right adjacent at i: %d, j %d\n", i, j)
				}
				if isNumberAdjacentUpperRight(input, i, j) {
					scanStarts = append(scanStarts, ScanStart{row: int(i - 1), column: int(j + 1)})
					fmt.Printf("found UpperRight adjacent at i: %d, j %d\n", i, j)
					if !isNumberAdjacentUp(input, i, j) && isNumberAdjacentUpperLeft(input, i, j) {
						scanStarts = append(scanStarts, ScanStart{row: int(i - 1), column: int(j - 1)})
						checkUpperLeft = false
						checkUp = false
						fmt.Printf("found second UpperLeft at i: %d, j %d\n", i, j)
					}
					if isNumberAdjacentUp(input, i, j) {
						checkUp = false
						if isNumberAdjacentUpperLeft(input, i, j) {
							checkUpperLeft = false
						}
					}
				}
				if checkUp && isNumberAdjacentUp(input, i, j) {
					scanStarts = append(scanStarts, ScanStart{row: int(i - 1), column: int(j)})
					if isNumberAdjacentUpperLeft(input, i, j) {
						checkUpperLeft = false
					}
					fmt.Printf("found Up adjacent at i: %d, j %d\n", i, j)
				}
				if checkUpperLeft && isNumberAdjacentUpperLeft(input, i, j) {
					scanStarts = append(scanStarts, ScanStart{row: int(i - 1), column: int(j - 1)})
					fmt.Printf("found UpperLeft adjacent at i: %d, j %d\n", i, j)
				}
				if isNumberAdjacentLowerRight(input, i, j) {
					scanStarts = append(scanStarts, ScanStart{row: int(i + 1), column: int(j + 1)})
					fmt.Printf("found LowerRight adjacent at i: %d, j %d\n", i, j)
					if !isNumberAdjacentDown(input, i, j) && isNumberAdjacentLowerLeft(input, i, j) {
						scanStarts = append(scanStarts, ScanStart{row: int(i + 1), column: int(j - 1)})
						checkDown = false
						checkLowerLeft = false
						fmt.Printf("found second LowerLeft adjacent at i: %d, j %d\n", i, j)
					}
					if isNumberAdjacentDown(input, i, j) {
						checkDown = false
						if isNumberAdjacentLowerLeft(input, i, j) {
							checkLowerLeft = false
						}
					}
				}
				if checkDown && isNumberAdjacentDown(input, i, j) {
					scanStarts = append(scanStarts, ScanStart{row: int(i + 1), column: int(j)})
					if isNumberAdjacentLowerLeft(input, i, j) {
						checkLowerLeft = false
					}
					fmt.Printf("found Down adjacent at i: %d, j %d\n", i, j)
				}
				if checkLowerLeft && isNumberAdjacentLowerLeft(input, i, j) {
					scanStarts = append(scanStarts, ScanStart{row: int(i + 1), column: int(j - 1)})
					fmt.Printf("found LowerLeft adjacent at i: %d, j %d\n", i, j)
				}
				if len(scanStarts) == 2 {
					fmt.Print("found two adjacent\n")
					var parts []uint64
					for i := 0; i < len(scanStarts); i++ {
						fmt.Print(scanStarts[i].row, " ", scanStarts[i].column, "\n")
						beginningOfPartNumber := findPartnumberStart(input[int(scanStarts[i].row)], scanStarts[i])
						partNumber := getPartNumber(input[int(scanStarts[i].row)], beginningOfPartNumber)
						parts = append(parts, partNumber)
						fmt.Printf("part number: %d\n", partNumber)
					}
					finalSum += parts[0] * parts[1]
				}
			}
		}
	}
	return finalSum
}

func findPartnumberStart(input string, scanStart ScanStart) ScanStart {
	for i := int(scanStart.column); i < len(input); i++ {
		if !isNumber(input[i]) {
			scanStart = ScanStart{row: scanStart.row, column: int(i - 1)}
			return scanStart
		}
	}
	return ScanStart{row: scanStart.row, column: int(len(input) - 1)}
}

func getPartNumber(input string, scanStart ScanStart) uint64 {
	var partNumber uint64 = 0
	var numberIndex uint8 = 0
	for i := scanStart.column; i >= 0; i-- {
		if isNumber(input[i]) {
			number, _ := strconv.Atoi(string(input[i]))
			partNumber = partNumber + uint64(number)*uint64(math.Pow10(int(numberIndex)))
			numberIndex++
		} else {
			return partNumber
		}
	}
	return partNumber
}

func main() {
	var input []string
	file, err := os.Open("../input.txt")
	//file, err := os.Open("test2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		text := scanner.Text()

		input = append(input, text)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sum := run(input)
	fmt.Print("\nsum: ", sum)
}
