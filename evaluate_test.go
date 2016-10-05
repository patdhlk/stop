package stop

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/patdhlk/stop/ast"
)

func TestEval(t *testing.T) {
	cases := []struct {
		Input      string
		Scope      *ast.BasicScope
		Error      bool
		Result     interface{}
		ResultType EvalType
	}{
		{
			Input:      "Hello World",
			Scope:      nil,
			Result:     "Hello World",
			ResultType: TString,
		},
		{
			Input:      `#{"foo\\bar"}`,
			Scope:      nil,
			Result:     `foo\bar`,
			ResultType: TString,
		},
		{
			Input:      `#{"foo\\\\bar"}`,
			Scope:      nil,
			Result:     `foo\\bar`,
			ResultType: TString,
		},
		{
			"#{var.alist}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"var.alist": ast.Variable{
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
			[]interface{}{"Hello", "World"},
			TList,
		},
		{
			"#{var.alist[1]}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"var.alist": ast.Variable{
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
			"World",
			TString,
		},
		{
			"#{var.alist[1]} #{var.alist[0]}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"var.alist": ast.Variable{
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
			"World Hello",
			TString,
		},
		{
			"#{var.alist} #{var.alist}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"var.alist": ast.Variable{
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
			true,
			nil,
			TUnsupported,
		},
		{
			`#{foo}`,
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
				},
			},
			false,
			map[string]interface{}{
				"foo": "hello",
				"bar": "world",
			},
			TMap,
		},
		{
			`#{foo["bar"]}`,
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
				},
			},
			false,
			"world",
			TString,
		},
		{
			`#{foo["bar"]} #{foo["foo"]}`,
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
				},
			},
			false,
			"world hello",
			TString,
		},
		{
			`#{foo} #{foo}`,
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
				},
			},
			true,
			nil,
			TUnsupported,
		},
		{
			`#{foo} #{bar}`,
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"foo": ast.Variable{
						Type:  ast.TString,
						Value: "Hello",
					},
					"bar": ast.Variable{
						Type:  ast.TString,
						Value: "World",
					},
				},
			},
			false,
			"Hello World",
			TString,
		},
		{
			`#{foo} #{bar}`,
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"foo": ast.Variable{
						Type:  ast.TString,
						Value: "Hello",
					},
					"bar": ast.Variable{
						Type:  ast.TInt,
						Value: 4,
					},
				},
			},
			false,
			"Hello 4",
			TString,
		},
		{
			`#{foo}`,
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
				},
			},
			false,
			map[string]interface{}{
				"foo": "hello",
				"bar": "world",
			},
			TMap,
		},
		{
			"#{var.alist}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"var.alist": ast.Variable{
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
			[]interface{}{
				"Hello",
				"World",
			},
			TList,
		},
	}

	for _, tc := range cases {
		node, err := Parse(tc.Input)
		if err != nil {
			t.Fatalf("Error: %s\n\nInput: %s", err, tc.Input)
		}

		result, err := Eval(node, &EvalConfig{GlobalScope: tc.Scope})
		if err != nil != tc.Error {
			t.Fatalf("Error: %s\n\nInput: %s", err, tc.Input)
		}
		if tc.ResultType != TUnsupported && result.Type != tc.ResultType {
			t.Fatalf("Bad: %s\n\nInput: %s", result.Type, tc.Input)
		}
		if !reflect.DeepEqual(result.Value, tc.Result) {
			t.Fatalf("\n     Bad: %#v\nExpected: %#v\n\nInput: %s", result.Value, tc.Result, tc.Input)
		}
	}
}

func TestEvalInternal(t *testing.T) {
	cases := []struct {
		Input      string
		Scope      *ast.BasicScope
		Error      bool
		Result     interface{}
		ResultType ast.Type
	}{
		{
			"foo",
			nil,
			false,
			"foo",
			ast.TString,
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
			"foo baz",
			ast.TString,
		},

		{
			"#{var.alist}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"var.alist": ast.Variable{
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
			[]ast.Variable{
				ast.Variable{
					Type:  ast.TString,
					Value: "Hello",
				},
				ast.Variable{
					Type:  ast.TString,
					Value: "World",
				},
			},
			ast.TList,
		},

		{
			"foo #{-29}",
			nil,
			false,
			"foo -29",
			ast.TString,
		},

		{
			"foo #{42+1}",
			nil,
			false,
			"foo 43",
			ast.TString,
		},

		{
			"foo #{42-1}",
			nil,
			false,
			"foo 41",
			ast.TString,
		},

		{
			"foo #{42*2}",
			nil,
			false,
			"foo 84",
			ast.TString,
		},

		{
			"foo #{42/2}",
			nil,
			false,
			"foo 21",
			ast.TString,
		},

		{
			"foo #{42/0}",
			nil,
			true,
			"foo ",
			ast.TUnsupported,
		},

		{
			"foo #{42%4}",
			nil,
			false,
			"foo 2",
			ast.TString,
		},

		{
			"foo #{42%0}",
			nil,
			true,
			"foo ",
			ast.TUnsupported,
		},

		{
			"foo #{42.0+1.0}",
			nil,
			false,
			"foo 43",
			ast.TString,
		},

		{
			"foo #{42.0+1}",
			nil,
			false,
			"foo 43",
			ast.TString,
		},

		{
			"foo #{42+1.0}",
			nil,
			false,
			"foo 43",
			ast.TString,
		},

		{
			"foo #{42+2*2}",
			nil,
			false,
			"foo 88",
			ast.TString,
		},

		{
			"foo #{42+(2*2)}",
			nil,
			false,
			"foo 46",
			ast.TString,
		},

		{
			"foo #{-bar}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"bar": ast.Variable{
						Value: 41,
						Type:  ast.TInt,
					},
				},
			},
			false,
			"foo -41",
			ast.TString,
		},

		{
			"foo #{bar+1}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"bar": ast.Variable{
						Value: 41,
						Type:  ast.TInt,
					},
				},
			},
			false,
			"foo 42",
			ast.TString,
		},

		{
			"foo #{bar+1}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"bar": ast.Variable{
						Value: "41",
						Type:  ast.TString,
					},
				},
			},
			false,
			"foo 42",
			ast.TString,
		},

		{
			"foo #{bar+baz}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"bar": ast.Variable{
						Value: "41",
						Type:  ast.TString,
					},
					"baz": ast.Variable{
						Value: "1",
						Type:  ast.TString,
					},
				},
			},
			false,
			"foo 42",
			ast.TString,
		},

		{
			"foo #{bar+baz}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"bar": ast.Variable{
						Value: 0.001,
						Type:  ast.TFloat,
					},
					"baz": ast.Variable{
						Value: "0.002",
						Type:  ast.TString,
					},
				},
			},
			false,
			"foo 0.003",
			ast.TString,
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
			"foo 42",
			ast.TString,
		},

		{
			`foo #{rand("foo", "bar")}`,
			&ast.BasicScope{
				FuncMap: map[string]ast.Function{
					"rand": ast.Function{
						ReturnType:   ast.TString,
						Variadic:     true,
						VariadicType: ast.TString,
						Callback: func(args []interface{}) (interface{}, error) {
							var result string
							for _, a := range args {
								result += a.(string)
							}
							return result, nil
						},
					},
				},
			},
			false,
			"foo foobar",
			ast.TString,
		},

		{
			`#{foo["bar"]}`,
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"foo": ast.Variable{
						Type: ast.TList,
						Value: []ast.Variable{
							ast.Variable{
								Type:  ast.TString,
								Value: "hello",
							},
							ast.Variable{
								Type:  ast.TString,
								Value: "world",
							},
						},
					},
				},
			},
			true,
			nil,
			ast.TUnsupported,
		},

		{
			`#{foo["bar"]}`,
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
				},
			},
			false,
			"world",
			ast.TString,
		},

		{
			`#{foo[var.key]}`,
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
					"var.key": ast.Variable{
						Type:  ast.TString,
						Value: "bar",
					},
				},
			},
			false,
			"world",
			ast.TString,
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
					"var.keyint": ast.Variable{
						Type:  ast.TInt,
						Value: 1,
					},
				},
			},
			false,
			"world",
			ast.TString,
		},

		{
			`#{foo["bar"]} #{bar[1]}`,
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
								Type:  ast.TInt,
								Value: 10,
							},
							ast.Variable{
								Type:  ast.TInt,
								Value: 20,
							},
						},
					},
				},
			},
			false,
			"world 20",
			ast.TString,
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
								Value: "hello",
							},
							ast.Variable{
								Type:  ast.TString,
								Value: "world",
							},
						},
					},
				},
			},
			false,
			"hello",
			ast.TString,
		},

		{
			"#{foo[bar]}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"foo": ast.Variable{
						Type: ast.TList,
						Value: []ast.Variable{
							ast.Variable{
								Type:  ast.TString,
								Value: "hello",
							},
							ast.Variable{
								Type:  ast.TString,
								Value: "world",
							},
						},
					},
					"bar": ast.Variable{
						Type:  ast.TInt,
						Value: 1,
					},
				},
			},
			false,
			"world",
			ast.TString,
		},

		{
			"#{foo[bar[1]]}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"foo": ast.Variable{
						Type: ast.TList,
						Value: []ast.Variable{
							ast.Variable{
								Type:  ast.TString,
								Value: "hello",
							},
							ast.Variable{
								Type:  ast.TString,
								Value: "world",
							},
						},
					},
					"bar": ast.Variable{
						Type: ast.TList,
						Value: []ast.Variable{
							ast.Variable{
								Type:  ast.TInt,
								Value: 1,
							},
							ast.Variable{
								Type:  ast.TInt,
								Value: 0,
							},
						},
					},
				},
			},
			false,
			"hello",
			ast.TString,
		},

		{
			"aaa #{foo} aaa",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"foo": ast.Variable{
						Type:  ast.TInt,
						Value: 42,
					},
				},
			},
			false,
			"aaa 42 aaa",
			ast.TString,
		},

		{
			"aaa #{foo[1]} aaa",
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
								Value: 24,
							},
						},
					},
				},
			},
			false,
			"aaa 24 aaa",
			ast.TString,
		},

		{
			"aaa #{foo[1]} - #{foo[0]}",
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
								Value: 24,
							},
						},
					},
				},
			},
			false,
			"aaa 24 - 42",
			ast.TString,
		},

		{
			"#{var.foo} #{var.foo[0]}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"var.foo": ast.Variable{
						Type: ast.TList,
						Value: []ast.Variable{
							ast.Variable{
								Type:  ast.TString,
								Value: "hello",
							},
							ast.Variable{
								Type:  ast.TString,
								Value: "world",
							},
						},
					},
				},
			},
			true,
			nil,
			ast.TUnsupported,
		},

		{
			"#{var.foo[0]} #{var.foo[1]}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"var.foo": ast.Variable{
						Type: ast.TList,
						Value: []ast.Variable{
							ast.Variable{
								Type:  ast.TString,
								Value: "hello",
							},
							ast.Variable{
								Type:  ast.TString,
								Value: "world",
							},
						},
					},
				},
			},
			false,
			"hello world",
			ast.TString,
		},

		{
			"#{foo[1]} #{foo[0]}",
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
								Value: 24,
							},
						},
					},
				},
			},
			false,
			"24 42",
			ast.TString,
		},

		{
			"#{foo[1-3]}",
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
								Value: 24,
							},
						},
					},
				},
			},
			true,
			nil,
			ast.TUnsupported,
		},

		{
			"#{foo[2]}",
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
								Value: 24,
							},
						},
					},
				},
			},
			true,
			nil,
			ast.TUnsupported,
		},

		// Testing implicit type conversions

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
			"foo 42",
			ast.TString,
		},

		{
			`foo #{foo("42")}`,
			&ast.BasicScope{
				FuncMap: map[string]ast.Function{
					"foo": ast.Function{
						ArgTypes:   []ast.Type{ast.TInt},
						ReturnType: ast.TString,
						Callback: func(args []interface{}) (interface{}, error) {
							return strconv.FormatInt(int64(args[0].(int)), 10), nil
						},
					},
				},
			},
			false,
			"foo 42",
			ast.TString,
		},

		// Multiline
		{
			"foo #{42+\n1.0}",
			nil,
			false,
			"foo 43",
			ast.TString,
		},

		// String vars should be able to implictly convert to floats
		{
			"#{1.5 * var.foo}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"var.foo": ast.Variable{
						Value: "42",
						Type:  ast.TString,
					},
				},
			},
			false,
			"63",
			ast.TString,
		},

		// Unary
		{
			"foo #{-46}",
			nil,
			false,
			"foo -46",
			ast.TString,
		},

		{
			"foo #{-46 + 5}",
			nil,
			false,
			"foo -41",
			ast.TString,
		},

		{
			"foo #{46 + -5}",
			nil,
			false,
			"foo 41",
			ast.TString,
		},

		{
			"foo #{-bar}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"bar": ast.Variable{
						Value: 41,
						Type:  ast.TInt,
					},
				},
			},
			false,
			"foo -41",
			ast.TString,
		},

		{
			"foo #{5 + -bar}",
			&ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"bar": ast.Variable{
						Value: 41,
						Type:  ast.TInt,
					},
				},
			},
			false,
			"foo -36",
			ast.TString,
		},
	}

	for _, tc := range cases {
		node, err := Parse(tc.Input)
		if err != nil {
			t.Fatalf("Error: %s\n\nInput: %s", err, tc.Input)
		}

		out, outType, err := internalEval(node, &EvalConfig{GlobalScope: tc.Scope})
		if err != nil != tc.Error {
			t.Fatalf("Error: %s\n\nInput: %s", err, tc.Input)
		}
		if tc.ResultType != ast.TUnsupported && outType != tc.ResultType {
			t.Fatalf("Bad: %s\n\nInput: %s", outType, tc.Input)
		}
		if !reflect.DeepEqual(out, tc.Result) {
			t.Fatalf("Bad: %#v\n\nInput: %s", out, tc.Input)
		}
	}
}
