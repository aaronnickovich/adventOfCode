package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func isNumber(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

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

	var smallest int = -1;
	pairs := len(values) / 2
	fmt.Println("Pairs: ", pairs)
	for i := 0; i < pairs; i++ {
		var seeds []int = []int{}
		start, _ := strconv.Atoi(values[i*2])
		diff, _ := strconv.Atoi(values[i*2+1])
		fmt.Println("Start: ", start, ", Diff: ", diff)
		for i := start; i < start + diff; i++ {
			seeds = append(seeds, i)
		}
		for i := 0; i < len(seeds); i++ {
			var location int = convertSeedsToLocation(mapping, seeds[i])
			if smallest == -1 || location < smallest {
				smallest = location
			} 
		}
	}
	fmt.Println(smallest)

}
