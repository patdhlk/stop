package stop

import (
	"testing"

	"github.com/patdhlk/stop/ast"
)

func TestTypeCheck(t *testing.T) {
	cases := []struct {
		Input string
		Scope ast.Scope
		Error bool
	}{
		{
			"foo",
			&ast.BasicScope{},
			false,
		},

		{
			"foo #{bar}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"bar": ast.Variable{
						Value: "baz",
						Type:  ast.TString,
					},
				},
			},
			false,
		},

		{
			"foo #{rand()}",
			&ast.BasicScope{
				FuncMap: map[string]ast.Function{
					"rand": ast.Function{
						ReturnType: ast.TString,
						Callback: func([]interface{}) (interface{}, error) {
							return "42", nil
						},
					},
				},
			},
			false,
		},

		{
			`foo #{rand("42")}`,
			&ast.BasicScope{
				FuncMap: map[string]ast.Function{
					"rand": ast.Function{
						ArgTypes:   []ast.Type{ast.TString},
						ReturnType: ast.TString,
						Callback: func([]interface{}) (interface{}, error) {
							return "42", nil
						},
					},
				},
			},
			false,
		},

		{
			`foo #{rand(42)}`,
			&ast.BasicScope{
				FuncMap: map[string]ast.Function{
					"rand": ast.Function{
						ArgTypes:   []ast.Type{ast.TString},
						ReturnType: ast.TString,
						Callback: func([]interface{}) (interface{}, error) {
							return "42", nil
						},
					},
				},
			},
			true,
		},

		{
			`foo #{rand()}`,
			&ast.BasicScope{
				FuncMap: map[string]ast.Function{
					"rand": ast.Function{
						ArgTypes:     nil,
						ReturnType:   ast.TString,
						Variadic:     true,
						VariadicType: ast.TString,
						Callback: func([]interface{}) (interface{}, error) {
							return "42", nil
						},
					},
				},
			},
			false,
		},

		{
			`foo #{rand("42")}`,
			&ast.BasicScope{
				FuncMap: map[string]ast.Function{
					"rand": ast.Function{
						ArgTypes:     nil,
						ReturnType:   ast.TString,
						Variadic:     true,
						VariadicType: ast.TString,
						Callback: func([]interface{}) (interface{}, error) {
							return "42", nil
						},
					},
				},
			},
			false,
		},

		{
			`foo #{rand("42", 42)}`,
			&ast.BasicScope{
				FuncMap: map[string]ast.Function{
					"rand": ast.Function{
						ArgTypes:     nil,
						ReturnType:   ast.TString,
						Variadic:     true,
						VariadicType: ast.TString,
						Callback: func([]interface{}) (interface{}, error) {
							return "42", nil
						},
					},
				},
			},
			true,
		},

		{
			"#{foo[0]}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"foo": ast.Variable{
						Type: ast.TList,
						Value: []ast.Variable{
							ast.Variable{
								Type:  ast.TString,
								Value: "Hello",
							},
							ast.Variable{
								Type:  ast.TString,
								Value: "World",
							},
						},
					},
				},
			},
			false,
		},

		{
			"#{foo[0]}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"foo": ast.Variable{
						Type: ast.TList,
						Value: []ast.Variable{
							ast.Variable{
								Type:  ast.TInt,
								Value: 3,
							},
							ast.Variable{
								Type:  ast.TString,
								Value: "World",
							},
						},
					},
				},
			},
			true,
		},

		{
			"#{foo[0]}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"foo": ast.Variable{
						Type:  ast.TString,
						Value: "Hello World",
					},
				},
			},
			true,
		},

		{
			"foo #{bar}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"bar": ast.Variable{
						Value: 42,
						Type:  ast.TInt,
					},
				},
			},
			true,
		},

		{
			"foo #{rand()}",
			&ast.BasicScope{
				FuncMap: map[string]ast.Function{
					"rand": ast.Function{
						ReturnType: ast.TInt,
						Callback: func([]interface{}) (interface{}, error) {
							return 42, nil
						},
					},
				},
			},
			true,
		},
	}

	for _, tc := range cases {
		node, err := Parse(tc.Input)
		if err != nil {
			t.Fatalf("Error: %s\n\nInput: %s", err, tc.Input)
		}

		visitor := &TypeCheck{Scope: tc.Scope}
		err = visitor.Visit(node)
		if err != nil != tc.Error {
			t.Fatalf("Error: %s\n\nInput: %s", err, tc.Input)
		}
	}
}

func TestTypeCheck_implicit(t *testing.T) {
	implicitMap := map[ast.Type]map[ast.Type]string{
		ast.TInt: {
			ast.TString: "intToString",
		},
	}

	cases := []struct {
		Input string
		Scope *ast.BasicScope
		Error bool
	}{
		{
			"foo #{bar}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"bar": ast.Variable{
						Value: 42,
						Type:  ast.TInt,
					},
				},
			},
			false,
		},

		{
			"foo #{foo(42)}",
			&ast.BasicScope{
				FuncMap: map[string]ast.Function{
					"foo": ast.Function{
						ArgTypes:   []ast.Type{ast.TString},
						ReturnType: ast.TString,
					},
				},
			},
			false,
		},

		{
			`foo #{foo("42", 42)}`,
			&ast.BasicScope{
				FuncMap: map[string]ast.Function{
					"foo": ast.Function{
						ArgTypes:     []ast.Type{ast.TString},
						Variadic:     true,
						VariadicType: ast.TString,
						ReturnType:   ast.TString,
					},
				},
			},
			false,
		},

		{
			"#{foo[1]}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"foo": ast.Variable{
						Type: ast.TList,
						Value: []ast.Variable{
							ast.Variable{
								Type:  ast.TInt,
								Value: 42,
							},
							ast.Variable{
								Type:  ast.TInt,
								Value: 23,
							},
						},
					},
				},
			},
			false,
		},

		{
			`#{foo[bar[var.keyint]]}`,
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"foo": ast.Variable{
						Type: ast.TMap,
						Value: map[string]ast.Variable{
							"foo": ast.Variable{
								Type:  ast.TString,
								Value: "hello",
							},
							"bar": ast.Variable{
								Type:  ast.TString,
								Value: "world",
							},
						},
					},
					"bar": ast.Variable{
						Type: ast.TList,
						Value: []ast.Variable{
							ast.Variable{
								Type:  ast.TString,
								Value: "i dont exist",
							},
							ast.Variable{
								Type:  ast.TString,
								Value: "bar",
							},
						},
					},
					"var.key": ast.Variable{
						Type:  ast.TInt,
						Value: 1,
					},
				},
			},
			false,
		},
	}

	for _, tc := range cases {
		node, err := Parse(tc.Input)
		if err != nil {
			t.Fatalf("Error: %s\n\nInput: %s", err, tc.Input)
		}

		// Modify the scope to add our conversion functions.
		if tc.Scope.FuncMap == nil {
			tc.Scope.FuncMap = make(map[string]ast.Function)
		}
		tc.Scope.FuncMap["intToString"] = ast.Function{
			ArgTypes:   []ast.Type{ast.TInt},
			ReturnType: ast.TString,
		}

		// Do the first pass...
		visitor := &TypeCheck{Scope: tc.Scope, Implicit: implicitMap}
		err = visitor.Visit(node)
		if err != nil != tc.Error {
			t.Fatalf("Error: %s\n\nInput: %s", err, tc.Input)
		}
		if err != nil {
			continue
		}

		// If we didn't error, then the next type check should not fail
		// WITHOUT implicits.
		visitor = &TypeCheck{Scope: tc.Scope}
		err = visitor.Visit(node)
		if err != nil {
			t.Fatalf("Error: %s\n\nInput: %s", err, tc.Input)
		}
	}
}
