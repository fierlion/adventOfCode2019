package main

import (
  "fmt"
  "strconv"
)

const START = 134564
const END = 585159

func main() {
  result := 0
  for i:=START; i<=END; i++ {
    if verifyPassword(i) {
      result += 1
    }
  } 
  fmt.Printf("result: %v\n", result)
}

func verifyPassword(password int) bool {
  // it's a six-digit number:
  if (!(password > 99999) && (password < 1000000)) {
    return false
  }

  stringified := strconv.Itoa(password)
  // there are never more than 2 repeating digits at least once
  counts := map[rune]int{}
  for _, v := range(stringified) {
    if count, ok := counts[v]; ok {
      counts[v] = count+1
    } else {
      counts[v] = 1
    }
  }
  adjacent := false
  for _, v := range(counts) {
    if v == 2 {
      adjacent = true
    }
  }
  if !adjacent {
    return false
  }

  // all are increasing
  for j:=0; j<len(stringified)-1; j++ {
    current := stringified[j]
    next := stringified[j+1]
    // all are increasing
    if current > next {
      return false
    }
  }

  
  if !adjacent {
    return false
  }
  return true  
}
