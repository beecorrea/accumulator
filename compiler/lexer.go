package compiler

import (
	"regexp"
	"strings"
)

type Token struct {
	Val  string
	Type string
}

type Lexer struct {
	Stream string
}

func (l *Lexer) Tokenize() []Token {
	var b strings.Builder
	tokens := make([]Token, 0)

	for i, char := range l.Stream {

		if isNum(char) {
			b.WriteString(string(char))
			if i == len(l.Stream)-1 {
				num := Token{Val: b.String(), Type: "NUMBER"}
				tokens = append(tokens, num)
				b.Reset()
			}
		} else if b.Len() > 0 {
			num := Token{Val: b.String(), Type: "NUMBER"}
			tokens = append(tokens, num)
			b.Reset()
		}

		if isOp(char) {
			op := Token{Val: string(char), Type: "OPERATOR"}
			tokens = append(tokens, op)
		}

	}

	return tokens
}

func (l *Lexer) GetTokensAsString(ts []Token) []string {
	tokens := make([]string, 0)
	for _, v := range ts {
		tokens = append(tokens, v.Val)
	}

	return tokens
}

func isOp(char rune) bool {
	if char == '+' || char == '-' {
		return true
	}
	return false
}

func isNum(char rune) bool {
	if ok, _ := regexp.MatchString("[0-9]", string(char)); ok {
		return true
	}
	return false
}
