package parser

import (
	"monkeylang/ast"
	"monkeylang/lexer"
	"monkeylang/token"
)

// Parser - l curToken and peekToken
// l - lexer instance
// curToken & peekToken - pointers similar to position
// and readPosition on the lexer
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

// New - create a new parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// read two tokens - prepopulate curToken and peekToken
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram - recursive descent parser
func (p *Parser) ParseProgram() *ast.Program {
	program = newProgramASTNode()

	advanceTokens()

	for currentToken() != EOF_TOKEN {
		statement = null
		if currentToken() == LET_TOKEN {
			statement = parseLetStatement()
		} else if currentToken() == RETURN_TOKEN {
			statement = parseReturnStatement()
		} else if currentToken() == IF_TOKEN {
			statement = parseIfStatement()
		}

		if statement != null {
			program.Statements.push(statement)
		}
		advanceTokens()
	}
	return program
}
