package main

import (
  "fmt"
  "os"
  "bufio"
)


func IsDigit(r byte) bool {
  return r >= '0' && r <= '9'
}

func IsChar(r byte) bool {
  return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func ReadFile(name string) []string {
  lines := make([]string, 0)
  
  read, err := os.Open(name)
  
  if err != nil {
    fmt.Println("An error occurred while reading the file.")
    os.Exit(1)
  }
  
  defer read.Close()
  
  scan := bufio.NewScanner(read)
  scan.Split(bufio.ScanLines)
  
  for scan.Scan() {
    lines = append(lines,  scan.Text())
  }
  
  return lines
}