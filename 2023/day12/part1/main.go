package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// Pad left-pads s with zeroes, to length n.
// If n is smaller than s, Pad is a no-op.
func Pad(s string, n int) (string, error) {
	return PadChar(s, n, '0')
}

// PadChar left-pads s with the rune r, to length n.
// If n is smaller than s, PadChar is a no-op.
func PadChar(s string, n int, r rune) (string, error) {
	if n < 0 {
		return "", fmt.Errorf("invalid length %d", n)
	}
	if len(s) > n {
		return s, nil
	}
	return strings.Repeat(string(r), n-len(s)) + s, nil
}

func main() {
	file, err := os.Open("test2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var sum uint64 = 0

	scanner := bufio.NewScanner(file)
	splitter := regexp.MustCompile(` `)
	partsSplitter := regexp.MustCompile(`,`)
	brokenSplitter := regexp.MustCompile(`\.+`)
	for scanner.Scan() {
		text := scanner.Text()
		// fmt.Println(text)
		splitInput := splitter.Split(text, -1)
		// splitInput[0] = splitInput[0] + splitInput[0] + splitInput[0] + splitInput[0] + splitInput[0]
		// splitInput[1] = splitInput[1] + "," + splitInput[1] + "," + splitInput[1] + "," + splitInput[1] + "," + splitInput[1]
		parts := partsSplitter.Split(splitInput[1], -1)
		fmt.Println(splitInput[0])
		fmt.Println(splitInput[1])
		var wildcards []int
		for i := 0; i < len(splitInput[0]); i++ {
			if splitInput[0][i] == '?' {
				wildcards = append(wildcards, i)
				// fmt.Println("wildcard found at: ", i)
			}
		}
		if len(wildcards) == 0 {
			fmt.Println("no wildcards, 1 iteration")
			sum++
			continue
		}

		totalIterations := int(math.Pow(float64(2), float64(len(wildcards))))
		// fmt.Println("totalIterations: ", totalIterations)
		ch := make(chan int)
		var wg sync.WaitGroup

		var workers int = 5
		workerRange := int(math.Ceil(float64(totalIterations) / float64(workers)))
		
		for w := 0; w < workers; w++ {
			wg.Add(1)
			start := w * workerRange 
			end := int(math.Min(float64(w + 1) * float64(workerRange), float64(totalIterations)))
			// fmt.Println("starting: ", start, "ending: ", end, "length: ", totalIterations)

			go func(ch chan int, wg *sync.WaitGroup, start int, end int, splitInput *string, wildcards *[]int) {
				defer wg.Done()
				// fmt.Println("iterating between: ", start, "and ", end)
				// 0 - 111
				internalSum := 0
				for i := start; i < end; i++ {
					x := i
					bin_value := strconv.FormatInt(int64(x), 2)
					bin_value, _ = Pad(bin_value, len(*wildcards))
					var index int = 0
					input := make([]byte, len(*splitInput))
					copy(input, *splitInput)
					for k := 0; k < len(input); k++ {
						if input[k] == '?' {
							if bin_value[index] == '0' {
								input[k] = '.'
							} else {
								input[k] = '#'
							}
							index++
						}
					}
					strInput := string(input)
					test := brokenSplitter.Split(strInput, -1)
					for l := 0; l < len(test); l++ {
						if test[l] == "" {
							if l == len(test)-1 {
								test = test[:l]
							} else {
								test = append(test[:l], test[l+1:]...)
							}
						}
					}
					fmt.Println("input: ", string(input))
					if len(test) == len(parts) {
						var matched bool = false
						for eachPart := 0; eachPart < len(parts); eachPart++ {
							thing, _ := strconv.Atoi(parts[eachPart])
							fmt.Println("test: ", test[eachPart], "parts: ", parts[eachPart])
							if len(test[eachPart]) == thing {
								matched = true
							} else {
								matched = false
								break
							}
						}
						if matched {
							internalSum++
							fmt.Println("sum: ", sum)
						}
					}
				}
				ch <- internalSum
				// fmt.Println(string(input))
			}(ch, &wg, start, end, &splitInput[0], &wildcards)
		}

		go func() {
			wg.Wait()
			close(ch)
		}()

		// 0 - 8
		for res := range ch {
			fmt.Println("res: ", res)
			sum += uint64(res)
		}
	}
	fmt.Println("final sum: ", sum)
}
