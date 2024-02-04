package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "strconv"
)
// 32T3K 765
// T55J5 684
// KK677 28
// KTJJT 220
// QQQJA 483

type node struct {
	cards string
	bid int
}

func convertCardLevel(a byte) int {
	switch a {
		case 'J':
			return 0
		case '2':
			return 1
		case '3':
			return 2
		case '4':
			return 3
		case '5':
			return 4
		case '6':
			return 5
		case '7':
			return 6
		case '8':
			return 7
		case '9':
			return 8
		case 'T':
			return 9
		case 'Q':
			return 10
		case 'K':
			return 11
		case 'A':
			return 12
	}
	return 0
}

func compareNodes(a node, b node) bool {
	aType := findRank(a.cards)
	bType := findRank(b.cards)
	if aType == bType {
		for i := 0; i < len(a.cards); i++ {
			if a.cards[i] == b.cards[i] {
				continue
			}
			return convertCardLevel(a.cards[i]) < convertCardLevel(b.cards[i])
		}
		return false
	}
	return aType < bType
}

func partition(arr []node, low, high int) ([]node, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if compareNodes(arr[j], pivot) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []node, low, high int) []node {
	if low < high {
		arr, p := partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []node) []node {
	return quickSort(arr, 0, len(arr)-1)
}

func findRank( input string ) int {
	var max int = -1;
	maps := make(map[byte]int)
	for i := 0; i < len(input); i++ {
		maps[input[i]] += 1
		if(input[i] != 'J' && maps[input[i]] > max) {
			max = maps[input[i]]
		}
	}

	length := len(maps)
	if(length > 1 && maps['J'] > 0 ){
		length -= 1
		max += maps['J']
        delete(maps, 'J')
	}
	if(length == 1) {
		return 6
	}
	if(length == 2) {
		if max == 4 || max == 1 {
			return 5
		}
		return 4
	}
	if(length == 3) {
		if max == 3 {
			return 3
		}
		if max == 2 {
			return 2
		}
	}
	if(length == 4) {
		return 1
	}
	return 0
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var nodes []node = []node{}
	valueSplitter := regexp.MustCompile(` +`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		text := scanner.Text()
		values := valueSplitter.Split(text, 2)
		bid, _ := strconv.Atoi(values[1])
		nodes = append(nodes, node{values[0], bid})
	}

	quickSortStart(nodes)
	sum := 0
	for i := 0; i < len(nodes); i++ {
	 sum += nodes[i].bid * (i+1)
	}
	fmt.Println(sum)
}
