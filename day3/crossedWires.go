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

type Direction int

const (
  U Direction = iota
  D
  L
  R
)

type Path struct {
  Goes Direction
  Distance int
}

func main() {
  inputFile, _ := filepath.Abs(INPUTFILE)
  file, err := os.Open(inputFile)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()
  
  reader := bufio.NewReader(file)
  rawLine1, err := reader.ReadString('\n')
  rawLine1 = strings.TrimSuffix(rawLine1, "\n")
  rawLine2, err := reader.ReadString('\n')
  rawLine2 = strings.TrimSuffix(rawLine2, "\n")
  wire1 := getWire(rawLine1)
  wire2 := getWire(rawLine2)

  // for each wire create set of all points visited
  path1 := getWirePath(wire1)
  path2 := getWirePath(wire2)
  overlaps := []string{}

  // fake set "or" operation
  for k, _ := range path1 {
    if _, ok := path2[k]; ok {
      overlaps = append(overlaps, k)  
    } 
  }

  // get min distance
  minDistance := int32(1<<31 - 1)
  for _, overlap := range overlaps {
    taxiDistance := getTaxiDistance(overlap)  
    if taxiDistance > int32(0) && taxiDistance < minDistance {
      minDistance = taxiDistance
    }
  }
  fmt.Printf("minDistance: %v\n", minDistance)
}

func getTaxiDistance(point string) int32 {
  splitPoint := strings.Split(point, "|")  
  pointX, _ := strconv.Atoi(splitPoint[0])
  pointY, _ := strconv.Atoi(splitPoint[1])
  return intAbsVal(int32(pointX)) + intAbsVal(int32(pointY))
}

func intAbsVal(val int32) int32 {
  if val > 0 {
    return val
  }
  return val * -1
}


func getWirePath(wire []Path) map[string]bool {
  wirePath := map[string]bool{}
  wireX := 0
  wireY := 0
  // for each point in wire update x/y points and add the string to the set  
  for i:=0; i<len(wire); i++ {
    thisPath := wire[i]
    switch move := thisPath.Goes; move {
      case U:
        endWire := wireY + thisPath.Distance
        for j:=wireY;j<endWire;j++ {
          pathString := strconv.Itoa(wireX) + "|" + strconv.Itoa(j)
          wirePath[pathString] = true
        }
        wireY += thisPath.Distance
      case D:
        endWire := wireY - thisPath.Distance
        for j:=wireY;j>endWire;j-- {
          pathString := strconv.Itoa(wireX) + "|" + strconv.Itoa(j)
          wirePath[pathString] = true
        }
        wireY -= thisPath.Distance
      case R:
        endWire := wireX + thisPath.Distance
        for j:=wireX;j<endWire;j++ {
          pathString := strconv.Itoa(j) + "|" + strconv.Itoa(wireY)
          wirePath[pathString] = true
        }
        wireX += thisPath.Distance
      case L:
        endWire := wireX - thisPath.Distance
        for j:=wireX;j>endWire;j-- {
          pathString := strconv.Itoa(j) + "|" + strconv.Itoa(wireY)
          wirePath[pathString] = true
        }
        wireX -= thisPath.Distance 
    }
  }
  return wirePath
}

func getWire(rawPath string) []Path {
  wireSplitArr := strings.Split(rawPath, ",")
  wire := []Path{}
  for i:=0; i<len(wireSplitArr); i++ {
    rawPath := wireSplitArr[i]
    distance, _ := strconv.Atoi(rawPath[1:])
    thisPath := Path{U, distance}
    switch direction := rawPath[0]; direction {
      case 'U':
        thisPath.Goes = U
      case 'D':
        thisPath.Goes = D
      case 'L':
        thisPath.Goes = L
      case 'R':
        thisPath.Goes = R
    }
    wire = append(wire, thisPath)
  }
  return wire
}

