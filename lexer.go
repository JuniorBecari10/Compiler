package main

import (
  "fmt"
  "os"
  "strconv"
)

const (
  NumberToken = iota
  IdentifierToken
  
  AssignToken
  
  LParenToken
  RParenToken
  
  PlusToken
  MinusToken
  TimesToken
  DivideToken
  
  IntTypeToken
  StrTypeToken
  FloatTypeToken
  
  NewLineToken
)

type Any interface {}

type Token struct {
  value Any
  kind int
  pos int
}

func Lex(s string, line int) []Token {
  tokens := make([]Token, 0)
  i := 0
  
  for i < len(s) {
    value := ""
    
    if s[i] == ' ' {
      i++
      continue
    } else if IsChar(s[i]) {
      start := i
      cont := false
      
      for i < len(s) && IsChar(s[i]) {
        value += string(s[i])
        i++
        
        if value == "int" {
          tokens = append(tokens, Token { value, IntTypeToken, start })
          cont = true
          break
        } else if value == "str" {
          tokens = append(tokens, Token { value, StrTypeToken, start })
          cont = true
          break
        } else if value == "float" {
          tokens = append(tokens, Token { value, FloatTypeToken, start })
          cont = true
          break
        }
      }
      
      if cont {
        continue
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
      
      tokens = append(tokens, Token { num, NumberToken, start })
      continue
    } else if s[i] == '=' {
      tokens = append(tokens, Token { s[i], AssignToken, i })
    } else if s[i] == '(' {
      tokens = append(tokens, Token { s[i], LParenToken, i })
    } else if s[i] == ')' {
      tokens = append(tokens, Token { s[i], RParenToken, i })
    } else if s[i] == '+' {
      tokens = append(tokens, Token { s[i], PlusToken, i })
    } else if s[i] == '-' {
      tokens = append(tokens, Token { s[i], MinusToken, i })
    } else if s[i] == '*' {
      tokens = append(tokens, Token { s[i], TimesToken, i })
    } else if s[i] == '/' {
      tokens = append(tokens, Token { s[i], DivideToken, i })
    } else {
      fmt.Printf("Syntax error on token '%s', on line %d, position %d.\n", string(s[i]), line + 1, i + 1)
      os.Exit(1)
    }
    
    i++
  }
  
  tokens = append(tokens, Token { "", NewLineToken, len(s) })
  
  return tokens
}