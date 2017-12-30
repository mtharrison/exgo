package exgo

import "testing"

func TestLex(t *testing.T) {

	formula := "15 * 16"
	tokens := Lex(formula)

	expected := []Token{
		Token{INTEGER, "15"},
		Token{OPERATOR_MUL, "*"},
		Token{INTEGER, "16"},
	}

	if !equalTokens(tokens, expected) {
		t.Errorf("Tokens %v did not match expected %v", tokens, expected)
	}
}

func equalTokens(a, b []Token) bool {

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if (a[i].literal != b[i].literal) ||
			(a[i].tokenType != b[i].tokenType) {
			return false
		}
	}

	return true
}
