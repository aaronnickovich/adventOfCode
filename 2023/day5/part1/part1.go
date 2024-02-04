package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "strconv"
)

func isNumber(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

// func genericMappingParser(input []string) (map[int]Mapping, []int) {
// 	var mapping map[int]Mapping = make(map[int]Mapping)
//
// 	splitter := regexp.MustCompile(` +`)
// 	for i := 1; i < len(input); i++ {
// 		splitMap := splitter.Split(input[i], 3)
// 		dest, _ := strconv.Atoi(splitMap[0])
// 		source, _ := strconv.Atoi(splitMap[1])
// 		diff, _ := strconv.Atoi(splitMap[2])
// 		mapping[source] = Mapping{dest: dest, diff: diff}
// 	}
// 	keys := make([]int, 0, len(mapping))
// 	for k := range mapping {
// 		keys = append(keys, k)
// 	}
// 	sortedKeys := quickSortStart(keys)
// 	return mapping, sortedKeys;
// }

// func partition(arr []int, low, high int) ([]int, int) {
// 	pivot := arr[high]
// 	i := low
// 	for j := low; j < high; j++ {
// 		if arr[j] < pivot {
// 			arr[i], arr[j] = arr[j], arr[i]
// 			i++
// 		}
// 	}
// 	arr[i], arr[high] = arr[high], arr[i]
// 	return arr, i
// }
//
// func quickSort(arr []int, low, high int) []int {
// 	if low < high {
// 		arr, p := partition(arr, low, high)
// 		arr = quickSort(arr, low, p-1)
// 		arr = quickSort(arr, p+1, high)
// 	}
// 	return arr
// }
//
// func quickSortStart(arr []int) []int {
// 	return quickSort(arr, 0, len(arr)-1)
// }

type Conversion struct {
	source int
	dest   int
	ran    int
}

type Mapping struct {
	seeds		[]int
	seed		[]Conversion
	soil		[]Conversion
	fertilizer	[]Conversion
	water		[]Conversion
	light		[]Conversion
	temperature	[]Conversion
	humidity	[]Conversion
}

func handleConversion(c []Conversion, input int) int {
	var output int = input
	for i := 0; i < len(c); i++ {
		if(output >= c[i].source && output < c[i].source + c[i].ran) {
			diff := output - c[i].source
			output = c[i].dest + diff
			break
		}
	}
	return output
}

func convertSeedsToLocation(m Mapping, input int) int {
	var output int = input
	output = handleConversion(m.seed, input)
	output = handleConversion(m.soil, output)
	output = handleConversion(m.fertilizer, output)
	output = handleConversion(m.water, output)
	output = handleConversion(m.light, output)
	output = handleConversion(m.temperature, output)
	output = handleConversion(m.humidity, output)
	return output
}

func main() {
	var input string = ""
	var mapping Mapping = Mapping{}
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	seedSplitter := regexp.MustCompile(`seeds: `)
	linesSplitter := regexp.MustCompile(`\n+`)
	conversionSplitter := regexp.MustCompile(`-to-`)
	spaceSplitter := regexp.MustCompile(` +`)

	seedMap := seedSplitter.Split(input, -1)
	lines := linesSplitter.Split(seedMap[1], -1)
	values := spaceSplitter.Split(lines[0], -1)
	mapping.seeds = make([]int, len(values))
	for i := 0; i < len(values); i++ {
		mapping.seeds[i], _ = strconv.Atoi(values[i])
	}

	lines = lines[1:]
	var types string = ""
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) > 0 && !isNumber(lines[i][0]) {
			types = conversionSplitter.Split(lines[i], 2)[0]
		} else {
			values := spaceSplitter.Split(lines[i], -1)
			if len(values) == 3 {
				source, _ := strconv.Atoi(values[1])
				dest, _ := strconv.Atoi(values[0])
				ran, _ := strconv.Atoi(values[2])
				if types == "seed" {
					mapping.seed = append(mapping.seed, Conversion{source: source, dest: dest, ran: ran})
				} else if types == "soil" {
					mapping.soil = append(mapping.soil, Conversion{source: source, dest: dest, ran: ran})
				} else if types == "fertilizer" {
					mapping.fertilizer = append(mapping.fertilizer, Conversion{source: source, dest: dest, ran: ran})
				} else if types == "water" {
					mapping.water = append(mapping.water, Conversion{source: source, dest: dest, ran: ran})
				} else if types == "light" {
					mapping.light = append(mapping.light, Conversion{source: source, dest: dest, ran: ran})
				} else if types == "temperature" {
					mapping.temperature = append(mapping.temperature, Conversion{source: source, dest: dest, ran: ran})
				} else if types == "humidity" {
					mapping.humidity = append(mapping.humidity, Conversion{source: source, dest: dest, ran: ran})
				}
			}
		}
	}

	fmt.Println(mapping.seeds)
	var smallest int = -1;
	for i := 0; i < len(mapping.seeds); i++ {
		var location int = convertSeedsToLocation(mapping, mapping.seeds[i])
		fmt.Println(location)
		if smallest == -1 || location < smallest {
			smallest = location
		} 
	}
	fmt.Println(smallest)

}
