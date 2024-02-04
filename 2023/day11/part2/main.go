package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Galaxy struct {
	row    int
	column int
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	distanceMultiplier := 1000000
	scanner := bufio.NewScanner(file)
	var input [][]byte
	var emptyRows map[int]bool = make(map[int]bool)
	var emptyColumns map[int]bool = make(map[int]bool)
	var rowIndex int = 0
	var galaxies []Galaxy
	for scanner.Scan() {
		var isEmpty bool = true
		text := scanner.Text()
		fmt.Println(text)
		var line []byte = make([]byte, len(text))
		for i := 0; i < len(text); i++ {
			line[i] = text[i]
			_, exists := emptyColumns[i]
			if text[i] != '#' {
				if !exists {
					emptyColumns[i] = true
				}
			} else {
				galaxies = append(
					galaxies,
					Galaxy{row: rowIndex, column: i})
				isEmpty = false
				emptyColumns[i] = false
			}
		}
		if isEmpty {
			emptyRows[rowIndex] = true
		}
		input = append(input, line)
		rowIndex++
	}
	var sum uint64 = 0
	each := 0
	for i := 0; i < len(galaxies); i++ {
		minDistance := math.MaxInt64
		fmt.Println(galaxies[i])
		for j := 0; j < len(galaxies); j++ {
			if j <= i {
				continue
			}
			each++
			currentHeight := galaxies[i].row - galaxies[j].row
			if currentHeight > 0 {
				for k, v := range emptyRows {
					if v && k < galaxies[i].row && k > galaxies[j].row {
						currentHeight = currentHeight + distanceMultiplier - 1
					}
				}
			}
			if currentHeight < 0 {
				currentHeight = currentHeight * -1
				for k, v := range emptyRows {
					if v && k > galaxies[i].row && k < galaxies[j].row {
						currentHeight = currentHeight + distanceMultiplier - 1
					}
				}
			}
			currentWidth := galaxies[i].column - galaxies[j].column
			if currentWidth > 0 {
				for k, v := range emptyColumns {
					if v && k < galaxies[i].column && k > galaxies[j].column {
						currentWidth = currentWidth + distanceMultiplier - 1
					}
				}
			}
			if currentWidth < 0 {
				currentWidth = currentWidth * -1
				for k, v := range emptyColumns {
					if v && k > galaxies[i].column && k < galaxies[j].column {
						currentWidth = currentWidth + distanceMultiplier - 1
					}
				}
			}
			distance := currentHeight + currentWidth
			sum += uint64(distance)
			if minDistance > distance {
				minDistance = distance
			}
			fmt.Println("Galaxies: ", galaxies[i], " --> ", galaxies[j], ", H: ", currentHeight, ", W: ", currentWidth, ", Dist: ", distance)
		}
		// sum += minDistance
	}
	fmt.Println("emptyRows: ", emptyRows)
	fmt.Println("emptyColumns", emptyColumns)
	fmt.Println("each: ", each)
	fmt.Println("sum: ", sum)
}
