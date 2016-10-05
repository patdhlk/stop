package stop_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/patdhlk/stop"
	"github.com/patdhlk/stop/ast"
)

func Example_functions() {
	input := "#{lower(var.test)} - #{6 + 2}"

	tree, err := stop.Parse(input)
	if err != nil {
		log.Fatal(err)
	}

	lowerCase := ast.Function{
		ArgTypes:   []ast.Type{ast.TString},
		ReturnType: ast.TString,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			input := inputs[0].(string)
			return strings.ToLower(input), nil
		},
	}

	config := &stop.EvalConfig{
		GlobalScope: &ast.BasicScope{
			VarMap: map[string]ast.Variable{
				"var.test": ast.Variable{
					Type:  ast.TString,
					Value: "TEST STRING",
				},
			},
			FuncMap: map[string]ast.Function{
				"lower": lowerCase,
			},
		},
	}

	result, err := stop.Eval(tree, config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Type: %s\n", result.Type)
	fmt.Printf("Value: %s\n", result.Value)
	// Output:
	// Type: TString
	// Value: test string - 8
}
