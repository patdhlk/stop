package ast

import (
	"bytes"
	"fmt"
)

// Output represents the root node of all interpolation evaluations. If the
// output only has one expression which is either a TList or TMap, the
// Output can be type-asserted to []interface{} or map[string]interface{}
// respectively. Otherwise the Output evaluates as a string, and concatenates
// the evaluation of each expression.
type Output struct {
	Exprs []Node
	Posx  Pos
}

func (n *Output) Accept(v Visitor) Node {
	for i, expr := range n.Exprs {
		n.Exprs[i] = expr.Accept(v)
	}

	return v(n)
}

func (n *Output) Pos() Pos {
	return n.Posx
}

func (n *Output) GoString() string {
	return fmt.Sprintf("*%#v", *n)
}

func (n *Output) String() string {
	var b bytes.Buffer
	for _, expr := range n.Exprs {
		b.WriteString(fmt.Sprintf("%s", expr))
	}

	return b.String()
}

func (n *Output) Type(s Scope) (Type, error) {
	// Special case no expressions for backward compatibility
	if len(n.Exprs) == 0 {
		return TString, nil
	}

	// Special case a single expression of types list or map
	if len(n.Exprs) == 1 {
		exprType, err := n.Exprs[0].Type(s)
		if err != nil {
			return TUnsupported, err
		}
		switch exprType {
		case TList:
			return TList, nil
		case TMap:
			return TMap, nil
		}
	}

	// Otherwise ensure all our expressions are strings
	for index, expr := range n.Exprs {
		exprType, err := expr.Type(s)
		if err != nil {
			return TUnsupported, err
		}
		// We only look for things we know we can't coerce with an implicit conversion func
		if exprType == TList || exprType == TMap {
			return TUnsupported, fmt.Errorf(
				"multi-expression STOP outputs may only have string inputs: %d is type %s",
				index, exprType)
		}
	}

	return TString, nil
}
