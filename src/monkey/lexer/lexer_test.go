package lexer 

// import the token package and the testing package 
import (
	"testing"
	"monkey/token"
)

// function that tests the nextToken function
func TestNextToken(t *testing.T){
	// the input for the test 
	input := '=+(){},;'

	
	tests := []struct {
		expectedType token.TokenType // the expected type of the token
		expectedLiteral string // the expected value of the token
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""}, // the end of the file
	}

	l := New(input) // create a new lexer with the input

	// loop through the tests 
	for i, tt := range tests {
		tok := l.NextToken() // get the next token and store it in tok

		// check if the type of the token is the same as the expected type
		if tok.Type != tt.expectedType {
			// if not, print an error message
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		// check if the literal of the token is the same as the expected literal
		if tok.Literal != tt.expectedLiteral {
			// if not, print an error message
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}