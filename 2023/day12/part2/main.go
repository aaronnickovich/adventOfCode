package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type State struct {
	group  int
	amount int
}

func main() {
	file, err := os.Open("test2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		valSplitter := regexp.MustCompile(` `)
		conditionSplitter := regexp.MustCompile(`,`)

		values := valSplitter.Split(line, -1)
		springs := values[0]
		springCondition := values[1]

		// convert to []int
		var splitConditions []string = conditionSplitter.Split(springCondition, -1)
		var conditions []int = make([]int, len(splitConditions))
		for i := 0; i < len(splitConditions); i++ {
			conditions[i], _ = strconv.Atoi(splitConditions[i])
		}

		var nextStates map[State]int = make(map[State]int)
		var states map[State]int = make(map[State]int)
		var init State = State{group: 0, amount: 0}
		states[init] = 1
		var brokenSpringsLeft int = 0
		for i := 0; i < len(springs); i++ {
			if springs[i] != '.' {
				brokenSpringsLeft++
			}
		}
		var minRequiredBrokenSpringsLeft []int = make([]int, 0)

		for i := 0; i <= len(conditions); i++ {
			var tmp []int = []int{}
			for j := i; j < len(conditions); j++ {
				tmp = append(tmp, conditions[j])
			}
			minRequiredBrokenSpringsLeft = append(minRequiredBrokenSpringsLeft, len(tmp))
		}

		for _, spring := range springs {
			if spring != '.' {
				brokenSpringsLeft--
			}
			for state, permutations := range states {
				// increase amount of broken springs for the current group
				// but only if the maximum of broken springs the current group hasn't been reached
				if spring == '#' || spring == '?' {
					if state.group < len(conditions) && state.amount < conditions[state.group] {
						var tmp State = State{group: state.group, amount: state.amount + 1}
						nextStates[tmp] = permutations
					}
				}

				// end current group if amount of broken springs equals the required amount
				// or keep going if there are no broken springs in the current group
				if spring == '.' || spring == '?' {
					if state.amount == 0 {
						// merge permutation count with other group that ended this loop
						var tmp State = State{group: state.group, amount: 0}
						permutations, exists := nextStates[tmp]
						if exists == true {
							nextStates[tmp] = permutations + nextStates[tmp]
						} else {
							nextStates[tmp] = permutations
						}
					} else if state.amount == conditions[state.group] {
						// merge permutation count with other group that's already been ended in a previous loop
						var tmp State = State{group: state.group + 1, amount: 0}
						permutations, exists := nextStates[tmp]
						if exists {
							nextStates[tmp] = permutations + nextStates[tmp]
						} else {
							nextStates[tmp] = permutations
						}
					}
				}
			}

			// remove all states that can't be finished because there aren't enough broken springs left
			for nextState := range nextStates {
				if brokenSpringsLeft+nextState.amount < minRequiredBrokenSpringsLeft[nextState.group] {
					delete(nextStates, nextState)
				}
			}

			states = make(map[State]int)
			for k, v := range nextStates {
				states[k] = v
			}
			nextStates = make(map[State]int)
		}
		var sum uint64 = 0
		for _, permutations := range states {
			sum += uint64(permutations)
		}

		fmt.Println("sum: ", sum)
	}
}

// func stuff() {
// 			ch := make(chan int)
// 		var wg sync.WaitGroup
//
// 		var workers int = 100
//
// 		for w := 0; w < workers; w++ {
// 			wg.Add(1)
// 			go func(ch chan int, wg *sync.WaitGroup, splitInput *string, wildcards *[]int) {
// 				defer wg.Done()
// 				ch <- internalSum
// 				// fmt.Println(string(input))
// 			}(ch, &wg, &splitInput[0], &wildcards)
// 		}
// 		go func() {
// 			wg.Wait()
// 			close(ch)
// 		}()
//
// 		for res := range ch {
// 			fmt.Println("res: ", res)
// 			sum += uint64(res)
// 		}
//
// }
