package ast

import "monkeylang/token"

// AST consists of interconnected nodes, forming a tree.

type Node interface {
	TokenLiteral() string // used for debugging
}

// Statement - do not produce values
type Statement interface {
	Node
	statementNode()
}

type LetStatement struct {
	Token token.Token // token.LET type
	Name  *Identifier // identifier
	Value Expression  // expression that produces the value
}

type ReturnStatement struct {
	Token       token.Token // return token
	ReturnValue Expression  // expression to be returned
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// Expression - return values
type Expression interface {
	Node
	expressionNode()
}

// Program - root node of every AST
// A program contains an array of connected nodes
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
