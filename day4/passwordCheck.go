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
  // there are 2 adjacent digits the same 
  adjacent := false 
  for i:=0; i<len(stringified)-1; i++ {
    current := stringified[i]
    next := stringified[i+1]
    if current == next {
      adjacent = true
    }
    // all are increasing
    // compare char val
    if current > next {
      return false
    }
  }
  if !adjacent {
    return false
  }
  return true  
}
