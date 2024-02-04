package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("test2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var sum int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var index int = 0
		text := scanner.Text()
		//fmt.Println(text)
		inputSplitter := regexp.MustCompile(` +`)
		textValues := inputSplitter.Split(text, -1)
		var history [][]int
		history = append(history, make([]int, 0))
		for i := 0; i < len(textValues); i++ {
			value, _ := strconv.Atoi(textValues[i])
			history[index] = append(history[index], value)
			history = doStuff(history, i)
			fmt.Println(history)
		}
		//next := calculateNextVal(history)
		prev := calculatePrevVal(history)
		sum += prev
		fmt.Println("prev: ", prev)
		fmt.Println("result: ", sum)
	}
}

func prependIntArray(x []int, y int) []int {
	x = append(x, 0)
	copy(x[1:], x)
	x[0] = y
	return x
}

func calculatePrevVal(history [][]int) int {
	var last int = len(history) - 1
	history[last] = prependIntArray(history[last], 0)
	for i := len(history) - 1; i > 0; i-- {
		prev := history[i-1][0] - history[i][0]

		history[i-1] = prependIntArray(history[i-1], prev)
	}

	fmt.Println("hist: ", history)
	return history[0][0]
}

func calculateNextVal(history [][]int) int {
	var last int = len(history) - 1
	history[last] = append(history[last], 0)
	for i := len(history) - 1; i > 0; i-- {
		length := len(history[i])
		next := history[i][length-1] + history[i-1][length-1]

		history[i-1] = append(history[i-1], next)
	}

	fmt.Println("hist: ", history)
	return history[0][len(history[0])-1]
}

func extrapolateArray(history [][]int) [][]int {
	var index int = 0
	for {
		var tmp []int = []int{}
		var maximum int = -9223372036854775808
		for i := 0; i < len(history[index])-1; i++ {
			sum := history[index][i+1] - history[index][i]
			tmp = append(tmp, sum)
			if sum > maximum {
				maximum = sum
			}
		}
		history = append(history, tmp)
		if maximum == 0 {
			return history
		}
		index++
	}
}

func doStuff(history [][]int, index int) [][]int {
	// start from 1 since the first row is handled outside of this loop
	for i := 1; i <= len(history); i++ {
		if index-i < 0 {
			break
		}
		diff := history[i-1][index-i+1] - history[i-1][index-i]
		if len(history)-1 < i {
			var tmp []int = []int{diff}
			history = append(history, tmp)
		} else {
			history[i] = append(history[i], diff)
		}
	}
	return history
}
