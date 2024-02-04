package main

import (
	"bufio"
	"fmt"
	"os"
)

func findSum(ghost string, instructions []int, nodes map[string][]string) (uint64, int) {
	var initialSum int = 0
	var sum uint64 = 0 
	var start string = ""
	var length int = len(instructions)
	for {
		for i := 0; i < length; i++ {
			fmt.Println(ghost)
			fmt.Println(sum)
			ghost = nodes[ghost][instructions[i]]
            if sum > 0 {
				sum++
			} else {
				initialSum++
			} 
			if(ghost[2] == 'Z') {
				if(start == "" && sum <= 0) {
					// fmt.Println("initialSum: ", initialSum)
					start = ghost
					sum++
				} else if(start == ghost) {
					sum = sum - 1
					return sum, initialSum
				}
			}
		}
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
	// fmt.Println(instructions)
	scanner.Scan()
	scanner.Text()
	
	var nodes map[string][]string = make(map[string][]string)

	var ghosts []string
	for scanner.Scan(){
		input := scanner.Text()

		start := input[0:3]
		left := input[7:10]
		right := input[12:15]
		nodes[start] = []string{left, right}
		if(input[2] == 'A'){
			ghosts = append(ghosts, start)
		}
	}
	var sums []int
	var starts []int
	for i := 0; i < len(ghosts); i++ {
		sum, initial := findSum(ghosts[i], instructions, nodes)
		sums = append(sums, int(sum))
		starts = append(starts, int(initial))
	}
	fmt.Println(sums)
	fmt.Println(starts)
	var current int = int(starts[0])
	for{
		// fmt.Println(current)
		for i := 0; i < len(ghosts); i++ {
			if(starts[i] == current) {
				if i == len(ghosts) - 1 {
					fmt.Println(current)
					return
				}
			} else {
				break
			}
		}
		for i := 0; i < len(ghosts); i++ {
			for (starts[i] < current) {
				starts[i] = starts[i] + sums[i]
			}
			current = starts[i]
		}
	}
}
