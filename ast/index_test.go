package ast

import (
	"strings"
	"testing"
)

func TestIndexTMap_empty(t *testing.T) {
	i := &Index{
		Target: &VariableAccess{Name: "foo"},
		Key: &LiteralNode{
			Typex: TString,
			Value: "bar",
		},
	}

	scope := &BasicScope{
		VarMap: map[string]Variable{
			"foo": Variable{
				Type:  TMap,
				Value: map[string]Variable{},
			},
		},
	}

	actual, err := i.Type(scope)
	if err == nil || !strings.Contains(err.Error(), "does not have any elements") {
		t.Fatalf("bad err: %s", err)
	}
	if actual != TUnsupported {
		t.Fatalf("bad: %s", actual)
	}
}

func TestIndexTMap_string(t *testing.T) {
	i := &Index{
		Target: &VariableAccess{Name: "foo"},
		Key: &LiteralNode{
			Typex: TString,
			Value: "bar",
		},
	}

	scope := &BasicScope{
		VarMap: map[string]Variable{
			"foo": Variable{
				Type: TMap,
				Value: map[string]Variable{
					"baz": Variable{
						Type:  TString,
						Value: "Hello",
					},
					"bar": Variable{
						Type:  TString,
						Value: "World",
					},
				},
			},
		},
	}

	actual, err := i.Type(scope)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if actual != TString {
		t.Fatalf("bad: %s", actual)
	}
}

func TestIndexTMap_int(t *testing.T) {
	i := &Index{
		Target: &VariableAccess{Name: "foo"},
		Key: &LiteralNode{
			Typex: TString,
			Value: "bar",
		},
	}

	scope := &BasicScope{
		VarMap: map[string]Variable{
			"foo": Variable{
				Type: TMap,
				Value: map[string]Variable{
					"baz": Variable{
						Type:  TInt,
						Value: 1,
					},
					"bar": Variable{
						Type:  TInt,
						Value: 2,
					},
				},
			},
		},
	}

	actual, err := i.Type(scope)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if actual != TInt {
		t.Fatalf("bad: %s", actual)
	}
}

func TestIndexTMap_nonHomogenous(t *testing.T) {
	i := &Index{
		Target: &VariableAccess{Name: "foo"},
		Key: &LiteralNode{
			Typex: TString,
			Value: "bar",
		},
	}

	scope := &BasicScope{
		VarMap: map[string]Variable{
			"foo": Variable{
				Type: TMap,
				Value: map[string]Variable{
					"bar": Variable{
						Type:  TString,
						Value: "Hello",
					},
					"baz": Variable{
						Type:  TInt,
						Value: 43,
					},
				},
			},
		},
	}

	_, err := i.Type(scope)
	if err == nil || !strings.Contains(err.Error(), "homogenous") {
		t.Fatalf("expected error")
	}
}

func TestIndexTList_empty(t *testing.T) {
	i := &Index{
		Target: &VariableAccess{Name: "foo"},
		Key: &LiteralNode{
			Typex: TInt,
			Value: 1,
		},
	}

	scope := &BasicScope{
		VarMap: map[string]Variable{
			"foo": Variable{
				Type:  TList,
				Value: []Variable{},
			},
		},
	}

	actual, err := i.Type(scope)
	if err == nil || !strings.Contains(err.Error(), "does not have any elements") {
		t.Fatalf("bad err: %s", err)
	}
	if actual != TUnsupported {
		t.Fatalf("bad: %s", actual)
	}
}

func TestIndexTList_string(t *testing.T) {
	i := &Index{
		Target: &VariableAccess{Name: "foo"},
		Key: &LiteralNode{
			Typex: TInt,
			Value: 1,
		},
	}

	scope := &BasicScope{
		VarMap: map[string]Variable{
			"foo": Variable{
				Type: TList,
				Value: []Variable{
					Variable{
						Type:  TString,
						Value: "Hello",
					},
					Variable{
						Type:  TString,
						Value: "World",
					},
				},
			},
		},
	}

	actual, err := i.Type(scope)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if actual != TString {
		t.Fatalf("bad: %s", actual)
	}
}

func TestIndexTList_int(t *testing.T) {
	i := &Index{
		Target: &VariableAccess{Name: "foo"},
		Key: &LiteralNode{
			Typex: TInt,
			Value: 1,
		},
	}

	scope := &BasicScope{
		VarMap: map[string]Variable{
			"foo": Variable{
				Type: TList,
				Value: []Variable{
					Variable{
						Type:  TInt,
						Value: 34,
					},
					Variable{
						Type:  TInt,
						Value: 54,
					},
				},
			},
		},
	}

	actual, err := i.Type(scope)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if actual != TInt {
		t.Fatalf("bad: %s", actual)
	}
}

func TestIndexTList_nonHomogenous(t *testing.T) {
	i := &Index{
		Target: &VariableAccess{Name: "foo"},
		Key: &LiteralNode{
			Typex: TInt,
			Value: 1,
		},
	}

	scope := &BasicScope{
		VarMap: map[string]Variable{
			"foo": Variable{
				Type: TList,
				Value: []Variable{
					Variable{
						Type:  TString,
						Value: "Hello",
					},
					Variable{
						Type:  TInt,
						Value: 43,
					},
				},
			},
		},
	}

	_, err := i.Type(scope)
	if err == nil || !strings.Contains(err.Error(), "homogenous") {
		t.Fatalf("expected error")
	}
}
