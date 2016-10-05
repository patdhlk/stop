package ast

import (
	"testing"
)

func TestVariableAccessType(t *testing.T) {
	c := &VariableAccess{Name: "foo"}
	scope := &BasicScope{
		VarMap: map[string]Variable{
			"foo": Variable{Type: TString},
		},
	}

	actual, err := c.Type(scope)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if actual != TString {
		t.Fatalf("bad: %s", actual)
	}
}

func TestVariableAccessType_invalid(t *testing.T) {
	c := &VariableAccess{Name: "bar"}
	scope := &BasicScope{
		VarMap: map[string]Variable{
			"foo": Variable{Type: TString},
		},
	}

	_, err := c.Type(scope)
	if err == nil {
		t.Fatal("should error")
	}
}

func TestVariableAccessType_list(t *testing.T) {
	c := &VariableAccess{Name: "baz"}
	scope := &BasicScope{
		VarMap: map[string]Variable{
			"baz": Variable{Type: TList},
		},
	}

	actual, err := c.Type(scope)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if actual != TList {
		t.Fatalf("bad: %s", actual)
	}
}
