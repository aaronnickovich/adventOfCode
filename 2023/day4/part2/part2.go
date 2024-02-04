package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
)

func parseRowsInputMock() string {
	var input string = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	return input
}

func getCardMatches(input string) int {
	splitter := regexp.MustCompile(`: +`)
	cardSplitter := regexp.MustCompile(` \| +`)
	delimiter := regexp.MustCompile(` +`)

	card := splitter.Split(input, 2)
	valueTypes := cardSplitter.Split(card[1], 2)
	cardValues := delimiter.Split(valueTypes[0], -1)
	winningValues := delimiter.Split(valueTypes[1], -1)
	var sum int = 0
	for i := 0; i < len(cardValues); i++ {
		var check bool = slices.Contains(winningValues, cardValues[i])
		if check {
			sum += 1
			fmt.Println("Cards won: ", sum)
		}
	}
	fmt.Println(cardValues)
	fmt.Println(winningValues)
	fmt.Printf("sum: %d\n", sum)
	return sum
}

func run(input []string) []int {
	var cards []int = make([]int, 300)
	for i := 0; i < 199; i++ {
		cards[i] = 1
	}

	for i := 0; i < len(input); i++ {
		output := getCardMatches(input[i])
		for j := 0; j < output; j++ {
			if i+j+1 <= (300 - 1) {
				if i+j+1 >= 199 {
					cards[i+j+1] += 1
				} else {
					cards[i+j+1] += 1 * cards[i]
				}
			}
		}
	}
	return cards
}

func main() {
	var input []string
	file, err := os.Open("../input.txt")
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

	cards := run(input)
	var sum int = 0

	for i := 0; i < 300; i++ {
		sum += cards[i]
		fmt.Printf("card[%d] = %d\n", i, cards[i])
	}

	fmt.Print("\nFinal Sum: ", sum)
}
