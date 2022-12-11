package compiler

import (
	"fmt"
	"strconv"
)

// Grammar description
// expr -> sum
// sum -> num rest
// rest -> op sum | E
// num -> NUMBER
// op -> PLUS

func (p *Parser) op() string {
	return p.read()
}

// num -> NUMBER
func (p *Parser) number() int {
	tok := p.read()
	val, _ := strconv.ParseInt(tok, 0, 64)
	return int(val)
}

// sum -> num (op num)*
func (p *Parser) sum() Sum {
	left := p.number()
	sum := Sum{left, Rest{}}

	for p.peek() == "+" {
		op := p.op()
		right := p.sum()
		rest := Rest{op, &right}
		sum = Sum{left, rest}
	}

	return sum
}

// expr -> sum
func (p *Parser) expr() Expression {
	return Expression{p.sum()}
}

func (p *Parser) peek() string {
	if p.Current >= len(p.Tokens) {
		return ""
	}

	return p.Tokens[p.Current]
}

func (p *Parser) read() string {
	tok := p.peek()
	p.Current++
	return tok
}

func (p *Parser) Parse() Expression {
	return p.expr()
}

type Parser struct {
	Current int
	Tokens  []string
}

type Expression struct {
	sum Sum
}

type Sum struct {
	left int
	rest Rest
}

type Rest struct {
	op  string
	sum *Sum
}

func (e Expression) Print() int {
	fmt.Printf("Parsed: ")
	return e.sum.print()
}

func (s Sum) print() int {
	var val int
	if s.left != 0 {
		fmt.Printf("%v ", s.left)
		val = s.left
	}

	if s.rest.op == "+" {
		return val + s.rest.print()
	}

	if s.rest.op == "-" {
		return val - s.rest.print()
	}

	return val
}

func (r Rest) print() int {
	if r.op != "" {
		fmt.Printf("%v ", r.op)
	}

	if r.sum != nil {
		return r.sum.print()
	}

	return 0
}
