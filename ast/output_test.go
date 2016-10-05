package ast

import (
	"testing"
)

func TestOutput_type(t *testing.T) {
	testCases := []struct {
		Name        string
		Output      *Output
		Scope       Scope
		ReturnType  Type
		ShouldError bool
	}{
		{
			Name:       "No expressions, for backward compatibility",
			Output:     &Output{},
			Scope:      nil,
			ReturnType: TString,
		},
		{
			Name: "Single string expression",
			Output: &Output{
				Exprs: []Node{
					&LiteralNode{
						Value: "Whatever",
						Typex: TString,
					},
				},
			},
			Scope:      nil,
			ReturnType: TString,
		},
		{
			Name: "Single list expression of strings",
			Output: &Output{
				Exprs: []Node{
					&VariableAccess{
						Name: "testvar",
					},
				},
			},
			Scope: &BasicScope{
				VarMap: map[string]Variable{
					"testvar": Variable{
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
			},
			ReturnType: TList,
		},
		{
			Name: "Single map expression",
			Output: &Output{
				Exprs: []Node{
					&VariableAccess{
						Name: "testvar",
					},
				},
			},
			Scope: &BasicScope{
				VarMap: map[string]Variable{
					"testvar": Variable{
						Type: TMap,
						Value: map[string]Variable{
							"key1": Variable{
								Type:  TString,
								Value: "Hello",
							},
							"key2": Variable{
								Type:  TString,
								Value: "World",
							},
						},
					},
				},
			},
			ReturnType: TMap,
		},
		{
			Name: "Multiple map expressions",
			Output: &Output{
				Exprs: []Node{
					&VariableAccess{
						Name: "testvar",
					},
					&VariableAccess{
						Name: "testvar",
					},
				},
			},
			Scope: &BasicScope{
				VarMap: map[string]Variable{
					"testvar": Variable{
						Type: TMap,
						Value: map[string]Variable{
							"key1": Variable{
								Type:  TString,
								Value: "Hello",
							},
							"key2": Variable{
								Type:  TString,
								Value: "World",
							},
						},
					},
				},
			},
			ShouldError: true,
			ReturnType:  TUnsupported,
		},
		{
			Name: "Multiple list expressions",
			Output: &Output{
				Exprs: []Node{
					&VariableAccess{
						Name: "testvar",
					},
					&VariableAccess{
						Name: "testvar",
					},
				},
			},
			Scope: &BasicScope{
				VarMap: map[string]Variable{
					"testvar": Variable{
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
			},
			ShouldError: true,
			ReturnType:  TUnsupported,
		},
		{
			Name: "Multiple string expressions",
			Output: &Output{
				Exprs: []Node{
					&VariableAccess{
						Name: "testvar",
					},
					&VariableAccess{
						Name: "testvar",
					},
				},
			},
			Scope: &BasicScope{
				VarMap: map[string]Variable{
					"testvar": Variable{
						Type:  TString,
						Value: "Hello",
					},
				},
			},
			ReturnType: TString,
		},
		{
			Name: "Multiple string expressions with coercion",
			Output: &Output{
				Exprs: []Node{
					&VariableAccess{
						Name: "testvar",
					},
					&VariableAccess{
						Name: "testint",
					},
				},
			},
			Scope: &BasicScope{
				VarMap: map[string]Variable{
					"testvar": Variable{
						Type:  TString,
						Value: "Hello",
					},
					"testint": Variable{
						Type:  TInt,
						Value: 2,
					},
				},
			},
			ReturnType: TString,
		},
	}

	for _, v := range testCases {
		actual, err := v.Output.Type(v.Scope)
		if err != nil && !v.ShouldError {
			t.Fatalf("case: %s\nerr: %s", v.Name, err)
		}
		if actual != v.ReturnType {
			t.Fatalf("case: %s\n     bad: %s\nexpected: %s\n", v.Name, actual, v.ReturnType)
		}
	}
}
