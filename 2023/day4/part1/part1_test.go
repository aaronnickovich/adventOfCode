package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestParseInput(t *testing.T) {
  var input []string;
  file, err := os.Open("test.txt");
  if err != nil {
    panic(err);
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  // optionally, resize scanner's capacity for lines over 64K, see next example
  for scanner.Scan() {
    text := scanner.Text()

    input = append(input, text);
  }

  if err := scanner.Err(); err != nil {
    panic(err);
  }

  sum := run(input);
  fmt.Print("\nsum: ", sum)

  if sum != 4361 {
    t.Fatalf(`result: %d, want 4361`, sum)
  }
}

