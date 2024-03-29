package stop

import (
	"reflect"
	"testing"

	"github.com/patdhlk/stop/ast"
)

func TestInterfaceToVariable_variableInput(t *testing.T) {
	_, err := InterfaceToVariable(ast.Variable{
		Type:  ast.TString,
		Value: "Hello world",
	})

	if err != nil {
		t.Fatalf("Bad: %s", err)
	}
}

func TestInterfaceToVariable(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected ast.Variable
	}{
		{
			name:  "string",
			input: "Hello world",
			expected: ast.Variable{
				Type:  ast.TString,
				Value: "Hello world",
			},
		},
		{
			name:  "empty list",
			input: []interface{}{},
			expected: ast.Variable{
				Type:  ast.TList,
				Value: []ast.Variable{},
			},
		},
		{
			name:  "empty list of strings",
			input: []string{},
			expected: ast.Variable{
				Type:  ast.TList,
				Value: []ast.Variable{},
			},
		},
		{
			name:  "int",
			input: 1,
			expected: ast.Variable{
				Type:  ast.TString,
				Value: "1",
			},
		},
		{
			name:  "list of strings",
			input: []string{"Hello", "World"},
			expected: ast.Variable{
				Type: ast.TList,
				Value: []ast.Variable{
					{
						Type:  ast.TString,
						Value: "Hello",
					},
					{
						Type:  ast.TString,
						Value: "World",
					},
				},
			},
		},
		{
			name:  "list of lists of strings",
			input: [][]interface{}{[]interface{}{"Hello", "World"}, []interface{}{"Goodbye", "World"}},
			expected: ast.Variable{
				Type: ast.TList,
				Value: []ast.Variable{
					{
						Type: ast.TList,
						Value: []ast.Variable{
							{
								Type:  ast.TString,
								Value: "Hello",
							},
							{
								Type:  ast.TString,
								Value: "World",
							},
						},
					},
					{
						Type: ast.TList,
						Value: []ast.Variable{
							{
								Type:  ast.TString,
								Value: "Goodbye",
							},
							{
								Type:  ast.TString,
								Value: "World",
							},
						},
					},
				},
			},
		},
		{
			name:  "map of string->string",
			input: map[string]string{"Hello": "World", "Foo": "Bar"},
			expected: ast.Variable{
				Type: ast.TMap,
				Value: map[string]ast.Variable{
					"Hello": {
						Type:  ast.TString,
						Value: "World",
					},
					"Foo": {
						Type:  ast.TString,
						Value: "Bar",
					},
				},
			},
		},
		{
			name: "map of lists of strings",
			input: map[string][]string{
				"Hello":   []string{"Hello", "World"},
				"Goodbye": []string{"Goodbye", "World"},
			},
			expected: ast.Variable{
				Type: ast.TMap,
				Value: map[string]ast.Variable{
					"Hello": {
						Type: ast.TList,
						Value: []ast.Variable{
							{
								Type:  ast.TString,
								Value: "Hello",
							},
							{
								Type:  ast.TString,
								Value: "World",
							},
						},
					},
					"Goodbye": {
						Type: ast.TList,
						Value: []ast.Variable{
							{
								Type:  ast.TString,
								Value: "Goodbye",
							},
							{
								Type:  ast.TString,
								Value: "World",
							},
						},
					},
				},
			},
		},
		{
			name:  "empty map",
			input: map[string]string{},
			expected: ast.Variable{
				Type:  ast.TMap,
				Value: map[string]ast.Variable{},
			},
		},
		{
			name: "three-element map",
			input: map[string]string{
				"us-west-1": "ami-123456",
				"us-west-2": "ami-456789",
				"eu-west-1": "ami-012345",
			},
			expected: ast.Variable{
				Type: ast.TMap,
				Value: map[string]ast.Variable{
					"us-west-1": {
						Type:  ast.TString,
						Value: "ami-123456",
					},
					"us-west-2": {
						Type:  ast.TString,
						Value: "ami-456789",
					},
					"eu-west-1": {
						Type:  ast.TString,
						Value: "ami-012345",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		output, err := InterfaceToVariable(tc.input)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("%s:\nExpected: %s\n     Got: %s\n", tc.name, tc.expected, output)
		}
	}
}

func TestVariableToInterface(t *testing.T) {
	testCases := []struct {
		name     string
		expected interface{}
		input    ast.Variable
	}{
		{
			name:     "string",
			expected: "Hello world",
			input: ast.Variable{
				Type:  ast.TString,
				Value: "Hello world",
			},
		},
		{
			name:     "empty list",
			expected: []interface{}{},
			input: ast.Variable{
				Type:  ast.TList,
				Value: []ast.Variable{},
			},
		},
		{
			name:     "int",
			expected: "1",
			input: ast.Variable{
				Type:  ast.TString,
				Value: "1",
			},
		},
		{
			name:     "list of strings",
			expected: []interface{}{"Hello", "World"},
			input: ast.Variable{
				Type: ast.TList,
				Value: []ast.Variable{
					{
						Type:  ast.TString,
						Value: "Hello",
					},
					{
						Type:  ast.TString,
						Value: "World",
					},
				},
			},
		},
		{
			name:     "list of lists of strings",
			expected: []interface{}{[]interface{}{"Hello", "World"}, []interface{}{"Goodbye", "World"}},
			input: ast.Variable{
				Type: ast.TList,
				Value: []ast.Variable{
					{
						Type: ast.TList,
						Value: []ast.Variable{
							{
								Type:  ast.TString,
								Value: "Hello",
							},
							{
								Type:  ast.TString,
								Value: "World",
							},
						},
					},
					{
						Type: ast.TList,
						Value: []ast.Variable{
							{
								Type:  ast.TString,
								Value: "Goodbye",
							},
							{
								Type:  ast.TString,
								Value: "World",
							},
						},
					},
				},
			},
		},
		{
			name:     "map of string->string",
			expected: map[string]interface{}{"Hello": "World", "Foo": "Bar"},
			input: ast.Variable{
				Type: ast.TMap,
				Value: map[string]ast.Variable{
					"Hello": {
						Type:  ast.TString,
						Value: "World",
					},
					"Foo": {
						Type:  ast.TString,
						Value: "Bar",
					},
				},
			},
		},
		{
			name: "map of lists of strings",
			expected: map[string]interface{}{
				"Hello":   []interface{}{"Hello", "World"},
				"Goodbye": []interface{}{"Goodbye", "World"},
			},
			input: ast.Variable{
				Type: ast.TMap,
				Value: map[string]ast.Variable{
					"Hello": {
						Type: ast.TList,
						Value: []ast.Variable{
							{
								Type:  ast.TString,
								Value: "Hello",
							},
							{
								Type:  ast.TString,
								Value: "World",
							},
						},
					},
					"Goodbye": {
						Type: ast.TList,
						Value: []ast.Variable{
							{
								Type:  ast.TString,
								Value: "Goodbye",
							},
							{
								Type:  ast.TString,
								Value: "World",
							},
						},
					},
				},
			},
		},
		{
			name:     "empty map",
			expected: map[string]interface{}{},
			input: ast.Variable{
				Type:  ast.TMap,
				Value: map[string]ast.Variable{},
			},
		},
		{
			name: "three-element map",
			expected: map[string]interface{}{
				"us-west-1": "ami-123456",
				"us-west-2": "ami-456789",
				"eu-west-1": "ami-012345",
			},
			input: ast.Variable{
				Type: ast.TMap,
				Value: map[string]ast.Variable{
					"us-west-1": {
						Type:  ast.TString,
						Value: "ami-123456",
					},
					"us-west-2": {
						Type:  ast.TString,
						Value: "ami-456789",
					},
					"eu-west-1": {
						Type:  ast.TString,
						Value: "ami-012345",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		output, err := VariableToInterface(tc.input)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("%s:\nExpected: %s\n     Got: %s\n", tc.name,
				tc.expected, output)
		}
	}
}
