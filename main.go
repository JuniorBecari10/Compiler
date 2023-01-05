package main

import (
  "fmt"
  "os"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Usage: compiler <file>")
    os.Exit(1)
  }
  
  lines := ReadFile(os.Args[1])
  
  for i, l := range lines {
    tokens := Lex(l, i)
    fmt.Println(tokens)
  }
}