package ast

import "github.com/cch123/monkey/token"

// Node is ast node
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

// Expression ...
type Expression interface {
	Node
	expressionNode()
}

// Program ...
type Program struct {
	Statements []Statement
}

// LetStatement ...
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// Identifier 标识符
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral ...
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral ...
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
