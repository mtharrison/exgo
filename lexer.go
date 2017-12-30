package exgo

import "unicode"

type TokenType int

const (
	INTEGER TokenType = iota
	OPERATOR_ADD
	OPERATOR_SUB
	OPERATIR_DIV
	OPERATOR_MUL
	OPERATOR_MOD
)

type Token struct {
	tokenType TokenType
	literal   string
}

func Lex(input string) []Token {

	pos := 0
	runes := []rune(input)
	tokens := []Token{}

	readNumber := func() {
		endPos := pos
		for endPos < len(runes) && unicode.IsDigit(runes[endPos]) {
			endPos++
		}
		tokens = append(tokens, Token{INTEGER, string(runes[pos:endPos])})
		pos = endPos
	}

	for pos < len(runes) {
		r := runes[pos]
		if unicode.IsDigit(r) {
			readNumber()
		}

		if r == '*' {
			tokens = append(tokens, Token{OPERATOR_MUL, "*"})
		}

		pos++
	}

	return tokens
}
