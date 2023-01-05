package main

import (
  //"fmt"
  "strconv"
)

const (
  NumberToken = iota
  IdentifierToken
  EqualsToken
  LParenToken
  RParenToken
)

type Any interface {}

type Token struct {
  value Any
  kind int
  pos int
}

func Lex(s string) []Token {
  tokens := make([]Token, 0)
  i := 0
  
  for i < len(s) {
    value := ""
    
    if IsChar(s[i]) {
      start := i
      
      for i < len(s) && IsChar(s[i]) {
        value += string(s[i])
        i++
      }
      
      tokens = append(tokens, Token { value, IdentifierToken, start })
      continue
    } else if IsDigit(s[i]) {
      start := i
      
      for i < len(s) && IsDigit(s[i]) {
        value += string(s[i])
        i++
      }
      
      num, _ := strconv.Atoi(value)
      
      tokens = append(tokens, Token { num, IdentifierToken, start })
      continue
    }
    
    i++
  }
  
  return tokens
}