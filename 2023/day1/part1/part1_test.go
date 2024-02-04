package main

import (
	"encoding/csv"
	"io"
	"os"
	"testing"
)

const testString = "one2three4"
const testEmptyString = "onetwo"

func TestFirstValueWithValues(t *testing.T) {
  result := readFirstValueFromString(testString)
  if result != 2 {
    t.Fatalf(`readFirstValueFromString(%s) = %d, want 2`, testString, result)
  }
}

func TestFirstValueWithoutValues(t *testing.T) {
  result := readFirstValueFromString(testEmptyString)
  if result != 0 {
    t.Fatalf(`readFirstValueFromString(%s) = %d, want 0`, testEmptyString, result)
  }
}


func TestLastValueWithValues(t *testing.T) {
  result := readLastValueFromString(testString)
  if result != 4 {
    t.Fatalf(`readLastValueFromString(%s) = %d, want 4`, testString, result)
  }
}

func TestLastValueWithoutValues(t *testing.T) {
  result := readLastValueFromString(testEmptyString)
  if result != 0 {
    t.Fatalf(`readLastValueFromString(%s) = %d, want 0`, testEmptyString, result)
  }
}

func TestFinalValuesFromCSV(t *testing.T) {
  var final uint64 = 0;
  f, err := os.Open("test.csv");
  if err != nil {
    panic(err);
  }
  defer f.Close();

  csvReader := csv.NewReader(f);
  for {
    record, err := csvReader.Read();
    if err == io.EOF {
      break;
    }
    if err != nil {
      panic(err);
    }
    var result uint8 = combineValuesFromString(record[0]);
    final += uint64(result);
  }

  if final != 142 {
    t.Fatalf(`final = %d, want 142`, final);
  }
}
