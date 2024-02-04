package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
)

func isNumber(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

func parseRowsInputMock() string {
	var input string = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	return input
}

func run(input string) uint64 {
	splitter := regexp.MustCompile(`: +`)
	cardSplitter := regexp.MustCompile(` \| +`)
	delimiter := regexp.MustCompile(` +`)

	card := splitter.Split(input, 2)
	valueTypes := cardSplitter.Split(card[1], 2)
	cardValues := delimiter.Split(valueTypes[0], -1)
	winningValues := delimiter.Split(valueTypes[1], -1)
	var sum uint64 = 0
	for i := 0; i < len(cardValues); i++ {
		var check bool = slices.Contains(winningValues, cardValues[i])
		if check {
			fmt.Println("Matched! ", cardValues[i])
			if sum == 0 {
				sum = 1
			} else {
				sum *= 2
			}
		}
	}
	fmt.Println(cardValues)
	fmt.Println(winningValues)
	fmt.Printf("sum: %d\n", sum)
	return sum
}

func main() {
	var sum uint64 = 0
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		input := scanner.Text()

		sum += run(input)
		fmt.Printf("Total: %d\n", sum)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Print("\nFinal Sum: ", sum)
}
