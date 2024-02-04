package main

import (
	// "encoding/csv"
	// "fmt"
	"bufio"
	"fmt"
	"os"
	"strconv"

	// "io"
	// "os"
	"regexp"
)

type gameRun struct {
  green int
  red int
  blue int
}

func parseInput(input string) uint64 {
  gameSplit := regexp.MustCompile(`: `);
  runSeparator := regexp.MustCompile(`; `); 
  colorSeparator := regexp.MustCompile(`, `); 
  typeSeparator := regexp.MustCompile(` `); 
  var idSplit []string = gameSplit.Split(input[5:], 2);
  var id, _ = strconv.Atoi(idSplit[0]);
  var runSplit []string = runSeparator.Split(idSplit[1], -1);
  for i := 0; i < len(runSplit); i++ {
    var types []string = colorSeparator.Split(runSplit[i], -1);
    for j := 0; j < len(types); j++ {
      var colorCount []string = typeSeparator.Split(types[j], 2);
      var rgb byte = colorCount[1][0];
      switch rgb {
        case 'g':
          run, _ := strconv.Atoi(colorCount[0]);
          if (run > 13) {
            return uint64(0);
          }
        case 'r':
          run, _ := strconv.Atoi(colorCount[0]);
          if(run > 12) {
            return uint64(0);
          }
        case 'b':
          run, _ := strconv.Atoi(colorCount[0]);
          if(run > 14) {
            return uint64(0);
          }
      }
    }
  }
  return uint64(id);
}

func main() {
  var final uint64 = 0;
  file, err := os.Open("../input.txt");
  if err != nil {
    panic(err);
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  // optionally, resize scanner's capacity for lines over 64K, see next example
  for scanner.Scan() {
    input := scanner.Text()

    result := parseInput(input);
    final += result;
  }

  if err := scanner.Err(); err != nil {
    panic(err);
  }

  fmt.Print(final);
}
