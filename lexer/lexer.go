package lexer

import "monkeylang/token"

// Lexer - lexer for the monkeylanguage
type Lexer struct {
	input        string
	position     int  // current position in input (points to current chat)
	readPosition int  // current reading position in input
	ch           byte // current char under examination
}

// New - returns a new lexer instance
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar - give us the next characted and advance position in the input string.
// In order to support full Unicode and UTF-8 (currently only ASCII) we need to change `l.ch` from
// byte to rune, and chance the way next char is read
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // check to see if we have reached end of input
		l.ch = 0 // set ch to 0 ~ ASCII code for 'NUL'
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition // position should point to the last read token
	l.readPosition += 1         // increment readposition so we know what comes next
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
