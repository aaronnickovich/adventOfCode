package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	row    int
	column int
}

type Routes struct {
	first  Node
	second Node
}

func findNextNode(pipes [][]byte, next Node, current Node) Node {
	fmt.Println("FindNextNode - current: ", current, "next", next)
	// moving up
	if next.row == current.row-1 {
		switch pipes[next.row][next.column] {
		case '|':
			return Node{row: next.row - 1, column: next.column}
		case '7':
			return Node{row: next.row, column: next.column - 1}
		case 'F':
			return Node{row: next.row, column: next.column + 1}
		}
	}

	// moving down
	if next.row == current.row+1 {
		switch pipes[next.row][next.column] {
		case '|':
			return Node{row: next.row + 1, column: next.column}
		case 'J':
			return Node{row: next.row, column: next.column - 1}
		case 'L':
			return Node{row: next.row, column: next.column + 1}
		}
	}

	// moving left
	if next.column == current.column-1 {
		switch pipes[next.row][next.column] {
		case '-':
			return Node{row: next.row, column: next.column - 1}
		case 'F':
			return Node{row: next.row + 1, column: next.column}
		case 'L':
			return Node{row: next.row - 1, column: next.column}
		}
	}

	// moving right
	if next.column == current.column+1 {
		switch pipes[next.row][next.column] {
		case '-':
			return Node{row: next.row, column: next.column + 1}
		case 'J':
			return Node{row: next.row - 1, column: next.column}
		case '7':
			return Node{row: next.row + 1, column: next.column}
		}

	}
	return Node{}
}

func findRoutes(pipes [][]byte, start Node) []Node {
	var routes []Node = make([]Node, 2)
	var tryNode Node = start
	var next Node
	var index int = 0

	// try up
	tryNode = start
	tryNode.row = start.row - 1
	if !(tryNode.row < 0) {
		next = findNextNode(pipes, tryNode, start)
		if (next != Node{}) {
			routes[index] = tryNode
			index++
		}
	}

	// try down
	tryNode = start
	tryNode.row = start.row + 1
	if !(tryNode.row >= len(pipes)) {
		next = findNextNode(pipes, tryNode, start)
		if (next != Node{}) {
			routes[index] = tryNode
			index++
		}
	}

	// try left
	tryNode = start
	tryNode.column = start.column - 1
	if !(tryNode.column < 0) {
		next = findNextNode(pipes, tryNode, start)
		if (next != Node{}) {
			routes[index] = tryNode
			index++
		}
	}

	// try right
	tryNode = start
	tryNode.column = start.column + 1
	if !(tryNode.column >= len(pipes)) {
		next = findNextNode(pipes, tryNode, start)
		if (next != Node{}) {
			routes[index] = tryNode
			index++
		}
	}
	return routes
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var pipes [][]byte
	var index int = 0
	var start Node
	for scanner.Scan() {
		text := scanner.Text()
		row := make([]byte, len(text))
		for i := 0; i < len(text); i++ {
			row[i] = text[i]
			if (row[i] == 'S' && start == Node{}) {
				start = Node{
					row:    index,
					column: i,
				}
			}
		}
		pipes = append(pipes, row)
		index++
	}
	var currentNodes []Node = make([]Node, 2)
	currentNodes[0] = start
	currentNodes[1] = start
	var nextNodes []Node = findRoutes(pipes, start)
	index = 1
	for {
		index++
		// fmt.Println("current nodes: ", currentNodes, " next nodes: ", nextNodes)
		var tmpNode0 Node
		var tmpNode1 Node
		tmpNode0 = findNextNode(pipes, nextNodes[0], currentNodes[0])
		// fmt.Println("tmpNode0: ", tmpNode0)
		if tmpNode0.row == nextNodes[1].row &&
			tmpNode0.column == nextNodes[1].column {
			break
		}

		tmpNode1 = findNextNode(pipes, nextNodes[1], currentNodes[1])
		// fmt.Println("tmpNode1: ", tmpNode1)
		if tmpNode1.row == tmpNode0.row &&
			tmpNode1.column == tmpNode0.column {
			break
		}

		//swap nodes
		currentNodes[0] = nextNodes[0]
		nextNodes[0] = tmpNode0

		currentNodes[1] = nextNodes[1]
		nextNodes[1] = tmpNode1
	}
	fmt.Println(pipes)
	fmt.Println("farthest count: ", index)
}
