package main

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"os"
	"io"
)

func readFirstValueFromString(test string) uint8 {
  var length int = len(test);
  for i := 0; i < length; i++ {
    if '0' <= test[i] && test[i] <= '9' {
      int, _ := strconv.ParseInt(string(test[i]), 10, 0)
      return uint8(int);
    }
    switch test[i] {
      case 'o': {
        //test one
        if(length > (i+2)) {
          if(test[i+1] == 'n' && test[i+2] == 'e') {
            return uint8(1);
          }
        }
      }
      case 't': {
        // test 'two'
        if(length > (i+2)) {
          if(test[i+1] == 'w' && test[i+2] == 'o') {
            return uint8(2);
          }
        }
        // test 'three'
        if(length > (i+4)) {
          if(test[i+1] == 'h' && test[i+2] == 'r' && test[i+3] == 'e' && test[i+4] == 'e') {
            return uint8(3);
          }
        }
      }
      case 'f': {
        if(length > (i+3)) {
          // test 'four'
          if(test[i+1] == 'o' && test[i+2] == 'u' && test[i+3] == 'r') {
            return uint8(4);
          }
          // test 'five'
          if(test[i+1] == 'i' && test[i+2] == 'v' && test[i+3] == 'e') {
            return uint8(5);
          }
        }
        }
      case 's': {
        // test 'six'
        if(length > (i+2)) {
          if(test[i+1] == 'i' && test[i+2] == 'x') {
            return uint8(6);
          }
        }
        // test 'seven'
        if(length > (i+4)) {
          if(test[i+1] == 'e' && test[i+2] == 'v' && test[i+3] == 'e' && test[i+4] == 'n') {
            return uint8(7);
          }
        }
      }
      case 'e': {
        // covers 'eight'
        if(length > (i+4)) {
          if(test[i+1] == 'i' && test[i+2] == 'g' && test[i+3] == 'h' && test[i+4] == 't') {
            return uint8(8);
          }
        }
      }
      case 'n': {
        // covers 'nine'
        if(length > (i+3)) {
          if(test[i+1] == 'i' && test[i+2] == 'n' && test[i+3] == 'e') {
            return uint8(9);
          }
        }
      }
      default:
        continue;
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
    switch test[i] {
      case 'e': {
        // test one
        if(0 <= i-2) {
          if(test[i-1] == 'n' && test[i-2] == 'o') {
            return uint8(1);
          }
        }
        // test three
        if(0 <= i-4) {
          if(test[i-1] == 'e' && test[i-2] == 'r' && test[i-3] == 'h' && test[i-4] == 't') {
            return uint8(3);
          }
        }
        // test five
        if(0 <= i-3) {
          if(test[i-1] == 'v' && test[i-2] == 'i' && test[i-3] == 'f') {
            return uint8(5);
          }
        }
        // test nine
        if(0 <= i-3) {
          if(test[i-1] == 'n' && test[i-2] == 'i' && test[i-3] == 'n') {
            return uint8(9);
          }
        }
      }

      case 'o': {
        // test 'two'
        if(0 <= (i-2)) {
          if(test[i-1] == 'w' && test[i-2] == 't') {
            return uint8(2);
          }
        }
      }
      case 'r': {
        if(0 <= i-3) {
          // test 'four'
          if(test[i-1] == 'u' && test[i-2] == 'o' && test[i-3] == 'f') {
            return uint8(4);
          }
        }
      }

      case 'x': {
        // test 'six'
        if(0 <= i-2) {
          if(test[i-1] == 'i' && test[i-2] == 's') {
            return uint8(6);
          }
        }
      }
      case 'n': {
        // test 'seven'
        if(0 <= i-4) {
          if(test[i-1] == 'e' && test[i-2] == 'v' && test[i-3] == 'e' && test[i-4] == 's') {
            return uint8(7);
          }
        }
      }
      case 't': {
        // covers 'eight'
        if(0 <= i-4) {
          if(test[i-1] == 'h' && test[i-2] == 'g' && test[i-3] == 'i' && test[i-4] == 'e') {
            return uint8(8);
          }
        }
      }
      default:
        continue;
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
