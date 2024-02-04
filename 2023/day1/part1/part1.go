package main

import (
  "encoding/csv"
  "fmt"
  "strconv"
  "os"
  "io"
)

func readFirstValueFromString(test string) uint8 {
  for i := 0; i < len(test); i++ {
    if '0' <= test[i] && test[i] <= '9' {
      int, _ := strconv.ParseInt(string(test[i]), 10, 0)
      return uint8(int);
    }
  }
  return uint8(0);
}

func readLastValueFromString(test string) uint8 {
  for i := len(test) - 1; i >= 0; i-- {
    if '0' <= test[i] && test[i] <= '9' {
      int, _ := strconv.ParseInt(string(test[i]), 10, 0)
      return uint8(int);
    }
  }
  return uint8(0);
}

func combineValuesFromString(test string) uint8 {
  var first uint8 = readFirstValueFromString(test);
  var last uint8 = readLastValueFromString(test);

  return ((first * 10) + last);
}

func main() {
  var final uint64 = 0;
  f, err := os.Open("../input.csv");
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

  fmt.Printf("result %d\n", final);
}
