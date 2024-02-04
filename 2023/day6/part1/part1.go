package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "strconv"
)

func calculateCountByDistance(totalTime int, recordDistance int) int {
	count := 0
	for i := 0; i < totalTime; i++ {
		rate := i
		remainingTime := totalTime - i

		tempDistance := rate * remainingTime;
		if(tempDistance > recordDistance){
			count++
			continue
		}
	}
	return count
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var times []int = []int{}
	var distances []int = []int{}
	splitter := regexp.MustCompile(`: +`)
	valueSplitter := regexp.MustCompile(` +`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		text := scanner.Text()
		result := splitter.Split(text, 2)
		values := valueSplitter.Split(result[1], -1)
		for i := 0; i < len(values); i++ {
			value, _ := strconv.Atoi(values[i])
			if(result[0] == "Time"){
				times = append(times, value)
			} else {
				distances = append(distances, value)
			}
		}
	}
	fmt.Println(times)
	fmt.Println(distances)

	var count int = 1
	for i := 0; i < len(times); i++ {
		count *= calculateCountByDistance(times[i], distances[i])
	}
	fmt.Println(count)
}
