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

func findRoutes(pipes [][]byte, start Node) ([]Node, byte) {
	var routes []Node = make([]Node, 2)
	var tryNode Node = start
	var next Node
	var index int = 0
	var up, down, left, right bool = false, false, false, false

	// try up
	tryNode = start
	tryNode.row = start.row - 1
	if !(tryNode.row < 0) {
		next = findNextNode(pipes, tryNode, start)
		if (next != Node{}) {
			routes[index] = tryNode
			index++
			up = true
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
			down = true
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
			left= true
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
			right = true
		}
	}
	if up {
		if down {
			pipes[start.row][start.column] = '|'
		}
		if right {
			pipes[start.row][start.column] = 'L'
		}
		if left {
			pipes[start.row][start.column] = 'J'
		}
	} else if down {
		if right {
			pipes[start.row][start.column] = 'F'
		}
		if left {
			pipes[start.row][start.column] = '7'
		}
	} else if right && left {
		pipes[start.row][start.column] = '-'
	}
	return routes, pipes[start.row][start.column]
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var originalPipes [][]byte
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

		copyRow := make([]byte, len(text))
		copy(copyRow, row) 
		originalPipes = append(originalPipes, copyRow)
		index++
	}
	var currentNodes []Node = make([]Node, 2)
	currentNodes[0] = start
	currentNodes[1] = start
	var nextNodes []Node
	nextNodes, originalPipes[start.row][start.column] = findRoutes(pipes, start)
	for {
		pipes[currentNodes[0].row][currentNodes[0].column] = 'X'
		pipes[currentNodes[1].row][currentNodes[1].column] = 'X'
		// fmt.Println("current nodes: ", currentNodes, " next nodes: ", nextNodes)
		var tmpNode0 Node
		var tmpNode1 Node
		tmpNode0 = findNextNode(pipes, nextNodes[0], currentNodes[0])
		// fmt.Println("tmpNode0: ", tmpNode0)
		if tmpNode0.row == nextNodes[1].row &&
			tmpNode0.column == nextNodes[1].column {
			pipes[nextNodes[1].row][nextNodes[1].column] = 'X'
			pipes[nextNodes[0].row][nextNodes[0].column] = 'X'
			pipes[tmpNode0.row][tmpNode0.column] = 'X'
			break
		}

		tmpNode1 = findNextNode(pipes, nextNodes[1], currentNodes[1])
		// fmt.Println("tmpNode1: ", tmpNode1)
		if tmpNode1.row == tmpNode0.row &&
			tmpNode1.column == tmpNode0.column {
			pipes[nextNodes[1].row][nextNodes[1].column] = 'X'
			pipes[nextNodes[0].row][nextNodes[0].column] = 'X'
			pipes[tmpNode1.row][tmpNode1.column] = 'X'
			break
		}

		//swap nodes
		currentNodes[0] = nextNodes[0]
		nextNodes[0] = tmpNode0

		currentNodes[1] = nextNodes[1]
		nextNodes[1] = tmpNode1
	}

	for i := 0; i < len(originalPipes); i++ {
		for j := 0; j < len(originalPipes[i]); j++ {
			if pipes[i][j] != 'X'{
				originalPipes[i][j] = '.'
			}
		}
		fmt.Println(string(originalPipes[i]))
	}

	var sum uint64 = 0
	
	//scan left to right for inner areas
	for i := 0; i < len(originalPipes); i++ {
		var foundEntry bool = false
		var dir string = ""
		for j := 0; j < len(originalPipes[i]); j++ {
			// is 'inside' the pipe
			if(foundEntry) {
				if originalPipes[i][j] == '.' {
					originalPipes[i][j] = '1'
					sum++
				} else if originalPipes[i][j] == '|' {
					foundEntry = false
				} else if originalPipes[i][j] == 'J' {
					if(dir == "up") {
						foundEntry = false
						dir = ""
					} else if(dir == "down") {
						foundEntry = true
						dir = ""
					}
				} else if originalPipes[i][j] == '7' {
					if(dir == "up") {
						foundEntry = true
						dir = ""
					} else if(dir == "down") {
						foundEntry = false
						dir = ""
					}
				} else if originalPipes[i][j] == 'F' {
					foundEntry = true
					dir = "up"
				} else if originalPipes[i][j] == 'L' {
					foundEntry = true
					dir = "down"
				}
			// is 'outside' the pipe
			} else {
				if originalPipes[i][j] == '.'{
					originalPipes[i][j] = '0'
				} else if originalPipes[i][j] == '|'{
					foundEntry = true
				} else if originalPipes[i][j] == 'F' {
					dir = "up"
					foundEntry = false
				} else if originalPipes[i][j] == 'L' {
					dir = "down"
					foundEntry = false
				} else if originalPipes[i][j] == 'J' {
					if(dir == "up") {
						foundEntry = true
						dir = ""
					} else if(dir == "down") {
						foundEntry = false
						dir = ""
					}
				} else if originalPipes[i][j] == '7' {
					if(dir == "up") {
						foundEntry = false
						dir = ""
					} else if(dir == "down") {
						foundEntry = true
						dir = ""
					}
				}
			}
		}
		fmt.Println(string(originalPipes[i]))
	}
	fmt.Println("sum: ",sum)
}
