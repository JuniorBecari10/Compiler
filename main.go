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
  
  for _, l := range lines {
    tokens := Lex(l)
    fmt.Println(tokens)
  }
}