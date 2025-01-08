package main

import (
	"jsonParser/lexer"
	"jsonParser/parser"
	"os"
)

func main() {
	filePath := os.Args[1]
	tokens := lexer.Lexer(filePath)
	parser.Parser(tokens)
}
