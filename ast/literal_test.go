package ast

import (
	"testing"
)

func TestLiteralNodeType(t *testing.T) {
	c := &LiteralNode{Typex: TString}
	actual, err := c.Type(nil)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if actual != TString {
		t.Fatalf("bad: %s", actual)
	}
}
