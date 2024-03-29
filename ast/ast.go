package ast

import (
	"fmt"
)

// Node is the interface that all AST nodes must implement.
type Node interface {
	// Accept is called to dispatch to the visitors. It must return the
	// resulting Node (which might be different in an AST transform).
	Accept(Visitor) Node

	// Pos returns the position of this node in some source.
	Pos() Pos

	// Type returns the type of this node for the given context.
	Type(Scope) (Type, error)
}

// Pos is the starting position of an AST node
type Pos struct {
	Column, Line int // Column/Line number, starting at 1
}

func (p Pos) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Column)
}

// Visitors are just implementations of this function.
//
// The function must return the Node to replace this node with. "nil" is
// _not_ a valid return value. If there is no replacement, the original node
// should be returned. We build this replacement directly into the visitor
// pattern since AST transformations are a common and useful tool and
// building it into the AST itself makes it required for future Node
// implementations and very easy to do.
//
// Note that this isn't a true implementation of the visitor pattern, which
// generally requires proper type dispatch on the function. However,
// implementing this basic visitor pattern style is still very useful even
// if you have to type switch.
type Visitor func(Node) Node

//go:generate stringer -type=Type

// Type is the type of any value.
type Type uint32

const (
	TUnsupported Type = 0
	TAny         Type = 1 << iota
	TString
	TInt
	TFloat
	TList
	TMap
	TBool
)

func (t Type) Printable() string {
	switch t {
	case TUnsupported:
		return "unsupported type"
	case TAny:
		return "any type"
	case TString:
		return "type string"
	case TInt:
		return "type int"
	case TFloat:
		return "type float"
	case TList:
		return "type list"
	case TMap:
		return "type map"
	case TBool:
		return "type bool"
	default:
		return "unknown type"
	}
}
