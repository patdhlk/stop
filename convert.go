package stop

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/patdhlk/stop/ast"
)

var STOPMapstructureDecodeHookSlice []interface{}
var STOPMapstructureDecodeHookStringSlice []string
var STOPMapstructureDecodeHookMap map[string]interface{}

// STOPMapstructureWeakDecode behaves in the same way as mapstructure.WeakDecode
// but has a DecodeHook which defeats the backward compatibility mode of mapstructure
// which WeakDecodes []interface{}{} into an empty map[string]interface{}. This
// allows us to use WeakDecode (desirable), but not fail on empty lists.
func STOPMapstructureWeakDecode(m interface{}, rawVal interface{}) error {
	config := &mapstructure.DecoderConfig{
		DecodeHook: func(source reflect.Type, target reflect.Type, val interface{}) (interface{}, error) {
			sliceType := reflect.TypeOf(STOPMapstructureDecodeHookSlice)
			stringSliceType := reflect.TypeOf(STOPMapstructureDecodeHookStringSlice)
			mapType := reflect.TypeOf(STOPMapstructureDecodeHookMap)

			if (source == sliceType || source == stringSliceType) && target == mapType {
				return nil, fmt.Errorf("Cannot convert %s into a %s", source, target)
			}

			return val, nil
		},
		WeaklyTypedInput: true,
		Result:           rawVal,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(m)
}

func InterfaceToVariable(input interface{}) (ast.Variable, error) {
	if inputVariable, ok := input.(ast.Variable); ok {
		return inputVariable, nil
	}

	var stringVal string
	if err := STOPMapstructureWeakDecode(input, &stringVal); err == nil {
		return ast.Variable{
			Type:  ast.TString,
			Value: stringVal,
		}, nil
	}

	var mapVal map[string]interface{}
	if err := STOPMapstructureWeakDecode(input, &mapVal); err == nil {
		elements := make(map[string]ast.Variable)
		for i, element := range mapVal {
			varElement, err := InterfaceToVariable(element)
			if err != nil {
				return ast.Variable{}, err
			}
			elements[i] = varElement
		}

		return ast.Variable{
			Type:  ast.TMap,
			Value: elements,
		}, nil
	}

	var sliceVal []interface{}
	if err := STOPMapstructureWeakDecode(input, &sliceVal); err == nil {
		elements := make([]ast.Variable, len(sliceVal))
		for i, element := range sliceVal {
			varElement, err := InterfaceToVariable(element)
			if err != nil {
				return ast.Variable{}, err
			}
			elements[i] = varElement
		}

		return ast.Variable{
			Type:  ast.TList,
			Value: elements,
		}, nil
	}

	return ast.Variable{}, fmt.Errorf("value for conversion must be a string, interface{} or map[string]interface: got %T", input)
}

func VariableToInterface(input ast.Variable) (interface{}, error) {
	if input.Type == ast.TString {
		if inputStr, ok := input.Value.(string); ok {
			return inputStr, nil
		} else {
			return nil, fmt.Errorf("ast.Variable with type string has value which is not a string")
		}
	}

	if input.Type == ast.TList {
		inputList, ok := input.Value.([]ast.Variable)
		if !ok {
			return nil, fmt.Errorf("ast.Variable with type list has value which is not a []ast.Variable")
		}

		result := make([]interface{}, 0)
		if len(inputList) == 0 {
			return result, nil
		}

		for _, element := range inputList {
			if convertedElement, err := VariableToInterface(element); err == nil {
				result = append(result, convertedElement)
			} else {
				return nil, err
			}
		}

		return result, nil
	}

	if input.Type == ast.TMap {
		inputMap, ok := input.Value.(map[string]ast.Variable)
		if !ok {
			return nil, fmt.Errorf("ast.Variable with type map has value which is not a map[string]ast.Variable")
		}

		result := make(map[string]interface{}, 0)
		if len(inputMap) == 0 {
			return result, nil
		}

		for key, value := range inputMap {
			if convertedValue, err := VariableToInterface(value); err == nil {
				result[key] = convertedValue
			} else {
				return nil, err
			}
		}

		return result, nil
	}

	return nil, fmt.Errorf("unknown input type: %s", input.Type)
}
