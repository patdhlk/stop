package stop_test

import (
	"fmt"
	"log"

	"github.com/patdhlk/stop"
)

func Example_basic() {
	input := "#{6 + 2}"

	tree, err := stop.Parse(input)
	if err != nil {
		log.Fatal(err)
	}

	result, err := stop.Eval(tree, &stop.EvalConfig{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Type: %s\n", result.Type)
	fmt.Printf("Value: %s\n", result.Value)
	// Output:
	// Type: TString
	// Value: 8
}
