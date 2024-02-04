package main

import (
	"bufio"
	"fmt"
	"os"
)

func findSum(root string, end string, instructions []int, nodes map[string][]string) uint64 {
	next := root
	var sum uint64 = 0
	var length int = len(instructions)
	for {
		for i := 0; i < len(instructions); i++ {
			next = nodes[next][instructions[i]]

			// fmt.Printf("next: %s, end: %s\n", next, end)
			if(next == end) {
				sum += uint64(i + 1)
				// fmt.Printf("sum: %d\n", sum)
				return sum
			}
		}
		sum += (uint64(length))
		// fmt.Println("next iteration")
	}

}
func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	text := scanner.Text()
	var instructions []int = []int{}
	for i := 0; i < len(text); i++ {
		if(text[i] == 'L') {
			instructions = append(instructions, 0)
		}

		if(text[i] == 'R') {
			instructions = append(instructions, 1)
		}
	}
	fmt.Println(instructions)
	scanner.Scan()
	scanner.Text()
	
	var nodes map[string][]string = make(map[string][]string)

	var root string = "AAA"
	var end string = "ZZZ"
	for scanner.Scan(){
		input := scanner.Text()
		if(len(input) < 14) {
			continue
		}

		start := input[0:3]
		left := input[7:10]
		right := input[12:15]
		nodes[start] = []string{left, right}
	}
	sum := findSum(root, end, instructions, nodes)
	fmt.Println(sum)
}
