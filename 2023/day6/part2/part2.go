package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	var time int = 0
	var distance int = 0
	splitter := regexp.MustCompile(`: +`)
	valueSplitter := regexp.MustCompile(` +`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		text := scanner.Text()
		result := splitter.Split(text, 2)
		values := valueSplitter.Split(result[1], -1)
		value := ""
		for i := 0; i < len(values); i++ {
			value += values[i]
		}
		valueInt, _ := strconv.Atoi(value)
		if(result[0] == "Time"){
			time = valueInt
		} else {
			distance = valueInt
		}
	}
	fmt.Println(time)
	fmt.Println(distance)

	var count int = calculateCountByDistance(time, distance)
	fmt.Println(count)
}
