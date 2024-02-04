package main

import (
	"testing"
)

const testSimple = "Game 1: 1 blue; 1 red; 1 green";
const testSimple2 = "Game 1: 1 blue, 2 red, 1 green; 2 blue, 1 red, 2 green";
const testComplexString = "Game 1: 17 green; 12 red, 14 blue; 10 red, 1 green, 1 blue";

func TestSimpleInput(t *testing.T) {
  var run gameRun = parseInput(testSimple);
  result := powerOfGameRun(run);
  if result != 1 {
    t.Fatalf(`result: %d, want 1`, result)
  }
}


func TestSimple2Input(t *testing.T) {
  var run gameRun = parseInput(testSimple2);
  result := powerOfGameRun(run);
  if result != 8 {
    t.Fatalf(`result: %d, want 8`, result)
  }
}

func TestComplexInput(t *testing.T) {
  var run gameRun = parseInput(testComplexString);
  result := powerOfGameRun(run);
  if result != 2856 {
    t.Fatalf(`result: %d, want 2856`, result)
  }
}
// func TestParseInputFailure(t *testing.T) {
//   var run gameRun = parseInput(testFailureString);
//   result := powerOfGameRun(run);
//   if result != 0 {
//     t.Fatalf(`result: %d, want 0`, result)
//   }
// }
//
// func TestParseInputSuccessfulEdgeCase(t *testing.T) {
//   var run gameRun = parseInput(testSuccessfulEdgeCaseString);
//   result := powerOfGameRun(run);
//   if result != 1 {
//     t.Fatalf(`result: %d, want 1`, result)
//   }
// }
