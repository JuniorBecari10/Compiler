package main

const (
  Expression = iota
  VarDeclStat
)

type ParseNode struct {
  kind int
  tokens []Token
}

func Parse(tokens []Token) ParseNode {
   
}

/*
Rules:

Expression: Identifiers, Integers, Floats, Parenthesis Plus, Minus, Times or Divide.
VarDeclStat: <name> <type: int | float | str> = <Expression>


*/