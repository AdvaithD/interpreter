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

	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
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
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok // early exit is necessary. Makes sure we dont call readChar() after switch again
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readIdentifier - reads in an identifier and advance lexer's position
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) { // iterate over all letters. used in default NextToken() case
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// isLetter - definition of a letter within the context of monkeylang. underscores allow for snake case
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// peekChar - we want to peek ahead in input and not traverse through the input
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// skipWhitespace - ignore whitespace
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
