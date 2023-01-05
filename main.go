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
  tokens := make([]Token, 0)
  
  for i, l := range lines {
    tks := Lex(l, i)
    fmt.Println(tks)
    
    tokens = append(tokens, tks)
  }
  
  
}