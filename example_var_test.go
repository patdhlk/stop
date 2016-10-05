package stop_test

import (
	"fmt"
	"log"

	"github.com/patdhlk/stop"
	"github.com/patdhlk/stop/ast"
)

func Example_variables() {
	input := "#{var.test} - #{6 + 2}"

	tree, err := stop.Parse(input)
	if err != nil {
		log.Fatal(err)
	}

	config := &stop.EvalConfig{
		GlobalScope: &ast.BasicScope{
			VarMap: map[string]ast.Variable{
				"var.test": ast.Variable{
					Type:  ast.TString,
					Value: "TEST STRING",
				},
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
	// Value: TEST STRING - 8
}
