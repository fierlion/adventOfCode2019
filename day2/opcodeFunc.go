package main

import (
  "bufio"
  "fmt"
  "os"
  "path/filepath"
  "strings"
  "strconv"
)

const INPUTFILE = "./input.txt"

func main() {
  inputFile, _ := filepath.Abs(INPUTFILE)
  file, err := os.Open(inputFile)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    rawInputArr := scanner.Text()
    splitInputArr := strings.Split(rawInputArr, ",")

    // find the noun/verb
    // update initial array
    // BRUTE FORCE BABY
    for n := 0; n < 100; n++ {
      for v := 0; v < 100; v++ {
        intInputArr := []int{}
        for i := 0; i < len(splitInputArr); i++ {
          converted, _ := strconv.Atoi(splitInputArr[i])      
          intInputArr = append(intInputArr, converted) 
        }
        intInputArr[1] = n
        intInputArr[2] = v 
        // opcode time
        // iterate over list in fours
        for o := 0; o < len(intInputArr); o += 4 {
          switch opcode := intInputArr[o]; opcode {
            case 1:
              // these should all be stuffed into a function...
              input1Pos := intInputArr[o+1]
              input2Pos := intInputArr[o+2]
              input1 := intInputArr[input1Pos]
              input2 := intInputArr[input2Pos]
              outputPos := intInputArr[o+3]
              intInputArr[outputPos] = input1 + input2
            case 2:
              input1Pos := intInputArr[o+1]
              input2Pos := intInputArr[o+2]
              input1 := intInputArr[input1Pos]
              input2 := intInputArr[input2Pos]
              outputPos := intInputArr[o+3]
              intInputArr[outputPos] = input1 * input2
            case 99:
              if intInputArr[0] == 19690720 {
                fmt.Printf("%v\n", (n*100)+v)
              }
            default:
              fmt.Printf("")
          }
        }
      }
    }
  }
}
