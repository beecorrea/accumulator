package main

import (
	"fmt"

	"github.com/beecorrea/accumulator/compiler"
)

func main() {
	stream := "5 + 5 + 50"
	fmt.Printf("Input expression: %s\n", stream)

	l := &compiler.Lexer{Stream: stream}
	tokens := l.Tokenize()
	fmt.Printf("Tokens: %v\n", tokens)

	stringfied := l.GetTokensAsString(tokens)
	p := &compiler.Parser{Current: 0, Tokens: stringfied}
	ast := p.Parse()

	total := ast.Print()

	fmt.Printf("\n\nTotal = %v\n", total)
}
