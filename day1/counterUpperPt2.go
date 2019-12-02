package main

import (
  "bufio"
  "fmt"
  "math"
  "os"
  "path/filepath"
  "strconv"
)

const INPUTFILE = "./input.txt"

func main() {
  total := 0
  inputFile, _ := filepath.Abs(INPUTFILE)
  file, err := os.Open(inputFile)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    curString := scanner.Text()
    curInt, _ := strconv.Atoi(curString)
    roundedInt := calculateMass(curInt) 
    for roundedInt > 0 {
      total += roundedInt
      roundedInt = calculateMass(roundedInt)
    }
  }
  fmt.Printf("total is: %v\n", total)
}

func calculateMass(baseMass int) int {
  divided := float64(baseMass) / 3
  rounded := math.Floor(divided) 
  subtracted := int(rounded) - 2
  return subtracted 
}
