package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

const testSuccessfulString = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green";
const testFailureString = "Game 1: 36 blue, 4 red; 1 red, 2 green, 6 blue; 2 green";
const testSuccessfulEdgeCaseString = "Game 1: 13 green, 12 red, 14 blue; 1 red, 1 green, 1 blue";

func TestParseInputSuccessful(t *testing.T) {
  result := parseInput(testSuccessfulString);
  if result != 1 {
    t.Fatalf(`result: %d, want 1`, result)
  }
}

func TestParseInputFailure(t *testing.T) {
  result := parseInput(testFailureString);
  if result != 0 {
    t.Fatalf(`result: %d, want 0`, result)
  }
}

func TestParseInputSuccessfulEdgeCase(t *testing.T) {
  result := parseInput(testSuccessfulEdgeCaseString);
  if result != 1 {
    t.Fatalf(`result: %d, want 1`, result)
  }
}

func TestInputsByFile(t *testing.T) {
  var final uint64 = 0;
  file, err := os.Open("test.txt");
  if err != nil {
    panic(err);
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  // optionally, resize scanner's capacity for lines over 64K, see next example
  for scanner.Scan() {
    input := scanner.Text()
    fmt.Println(input)

    result := parseInput(input);
    fmt.Println(result)
    final += result;
  }

  if err := scanner.Err(); err != nil {
    panic(err);
  }

  if final != 8 {
    t.Fatalf(`result: %d, want 8`, final)
  }
}
