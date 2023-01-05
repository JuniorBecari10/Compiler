package main

import (
  //"fmt"
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
    }
    
    i++
  }
  
  return tokens
}