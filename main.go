package main

import (
	"fmt"
	"jsonParser/lexer"
	"os"
)

func main() {
	filePath := os.Args[1]
	tokens := lexer.Lexer(filePath)
	fmt.Printf("Tokens: %v\n", tokens)
	// parser.Parser(tokens)
}
