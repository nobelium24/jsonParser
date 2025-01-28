package main

import (
	"fmt"
	"jsonParser/lexer"
	"jsonParser/parser"
	"os"
)

func main() {
	filePath := os.Args[1]
	tokens := lexer.Lexer(filePath)
	fmt.Printf("Tokens: %v\n", tokens)
	if len(tokens) == 0 {
		fmt.Println("Error: No tokens were parsed. Check your lexer function.")
		return
	}
	index := 0
	parser, err := parser.Parser(tokens, &index)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("Parsed tokens:%v\n", parser)
}
