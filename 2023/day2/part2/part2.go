package main

import (
	// "fmt"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"regexp"
)

type gameRun struct {
  green int
  red int
  blue int
}

func parseInput(input string) gameRun {
  gameSplit := regexp.MustCompile(`: `);
  runSeparator := regexp.MustCompile(`; `); 
  colorSeparator := regexp.MustCompile(`, `); 
  typeSeparator := regexp.MustCompile(` `); 
  var idSplit []string = gameSplit.Split(input[5:], 2);
  gameRun := gameRun{0,0,0};
  var runSplit []string = runSeparator.Split(idSplit[1], -1);
  for i := 0; i < len(runSplit); i++ {
    var types []string = colorSeparator.Split(runSplit[i], -1);
    for j := 0; j < len(types); j++ {
      var colorCount []string = typeSeparator.Split(types[j], 2);
      var rgb byte = colorCount[1][0];
      switch rgb {
        case 'g':
          run, _ := strconv.Atoi(colorCount[0]);
          if(gameRun.green < run) {
            gameRun.green = run
          }
        case 'r':
          run, _ := strconv.Atoi(colorCount[0]);
          if(gameRun.red < run) {
            gameRun.red = run
          }
        case 'b':
          run, _ := strconv.Atoi(colorCount[0]);
          if(gameRun.blue < run) {
            gameRun.blue = run
          }
      }
    }
  }
  return gameRun;
}

func powerOfGameRun(run gameRun) uint64 {
  var power uint64 = uint64(run.green) * uint64(run.red) * uint64(run.blue);
  return power;
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
    fmt.Println(input)

    var run gameRun = parseInput(scanner.Text());
    result := powerOfGameRun(run);
    fmt.Println(result)
    final += result;
  }

  if err := scanner.Err(); err != nil {
    panic(err);
  }

  fmt.Print(final);
}
