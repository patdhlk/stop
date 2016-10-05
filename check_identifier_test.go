package stop

import (
	"testing"

	"github.com/patdhlk/stop/ast"
)

func TestIdentifierCheck(t *testing.T) {
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
			"foo #{bar} success",
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
			"foo #{bar}",
			&ast.BasicScope{},
			true,
		},

		{
			"foo #{rand()} success",
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
			"foo #{rand()}",
			&ast.BasicScope{},
			true,
		},

		{
			"foo #{rand(42)} ",
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
			true,
		},

		{
			"foo #{rand()} ",
			&ast.BasicScope{
				FuncMap: map[string]ast.Function{
					"rand": ast.Function{
						ReturnType:   ast.TString,
						Variadic:     true,
						VariadicType: ast.TInt,
						Callback: func([]interface{}) (interface{}, error) {
							return "42", nil
						},
					},
				},
			},
			false,
		},

		{
			"foo #{rand(42)} ",
			&ast.BasicScope{
				FuncMap: map[string]ast.Function{
					"rand": ast.Function{
						ReturnType:   ast.TString,
						Variadic:     true,
						VariadicType: ast.TInt,
						Callback: func([]interface{}) (interface{}, error) {
							return "42", nil
						},
					},
				},
			},
			false,
		},

		{
			"foo #{rand(\"foo\", 42)} ",
			&ast.BasicScope{
				FuncMap: map[string]ast.Function{
					"rand": ast.Function{
						ArgTypes:     []ast.Type{ast.TString},
						ReturnType:   ast.TString,
						Variadic:     true,
						VariadicType: ast.TInt,
						Callback: func([]interface{}) (interface{}, error) {
							return "42", nil
						},
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

		visitor := &IdentifierCheck{Scope: tc.Scope}
		err = visitor.Visit(node)
		if err != nil != tc.Error {
			t.Fatalf("Error: %s\n\nInput: %s", err, tc.Input)
		}
	}
}
