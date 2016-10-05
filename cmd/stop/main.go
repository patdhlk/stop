package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"reflect"
	"strings"

	"github.com/patdhlk/stop"
	"github.com/patdhlk/stop/ast"
)

var (
	inputPtr string
)

var (
	LowerCase = ast.Function{
		ArgTypes:   []ast.Type{ast.TString},
		ReturnType: ast.TString,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			input := inputs[0].(string)
			return strings.ToLower(input), nil
		},
	}

	Pow = ast.Function{
		ArgTypes:   []ast.Type{ast.TFloat, ast.TFloat},
		ReturnType: ast.TFloat,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			basis := inputs[0].(float64)
			exponent := inputs[1].(float64)
			return math.Pow(basis, exponent), nil
		},
	}

	Equal = ast.Function{
		ArgTypes: []ast.Type{
			ast.TAny,
			ast.TAny,
		},
		ReturnType: ast.TBool,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			first := inputs[0]
			second := inputs[1]

			return reflect.DeepEqual(first, second), nil
		},
	}
)

func init() {
	flag.StringVar(&inputPtr, "input", `#{lower(var.test)} - #{6 + 2} + #{pow(var2.test,2)}`, "the stop string which should be parsed")
}

func main() {
	flag.Parse()
	var five float64 = 5
	fmt.Printf("Input: %s\n", inputPtr)
	tree, err := stop.Parse(inputPtr)
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
				"var2.test": ast.Variable{
					Type:  ast.TFloat,
					Value: five,
				},
			},
			FuncMap: map[string]ast.Function{
				"lower": LowerCase,
				"pow":   Pow,
			},
		},
	}

	result, err := stop.Eval(tree, config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Type: %s\n", result.Type)
	fmt.Printf("Value: %s\n", result.Value)

	input := "#{equal(1, var.test)}"
	fmt.Printf("Input: %s\n", input)

	tree, err = stop.Parse(input)
	if err != nil {
		log.Fatal(err)
	}

	config = &stop.EvalConfig{
		GlobalScope: &ast.BasicScope{
			VarMap: map[string]ast.Variable{
				"var.test": ast.Variable{
					Type: ast.TInt,
					// change to 2 to see that it outputs false
					Value: 2,
				},
			},
			FuncMap: map[string]ast.Function{
				"equal": Equal,
			},
		},
	}

	result, err = stop.Eval(tree, config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Type: %s\n", result.Type)
	fmt.Printf("Value: %s\n", result.Value)
}
