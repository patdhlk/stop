package stop

import (
	"errors"
	"strconv"

	"github.com/patdhlk/stop/ast"
)

// NOTE: All builtins are tested in engine_test.go

func registerBuiltins(scope *ast.BasicScope) *ast.BasicScope {
	if scope == nil {
		scope = new(ast.BasicScope)
	}
	if scope.FuncMap == nil {
		scope.FuncMap = make(map[string]ast.Function)
	}

	// Implicit conversions
	scope.FuncMap["__builtin_BoolToString"] = builtinBoolToString()
	scope.FuncMap["__builtin_FloatToInt"] = builtinFloatToInt()
	scope.FuncMap["__builtin_FloatToString"] = builtinFloatToString()
	scope.FuncMap["__builtin_IntToFloat"] = builtinIntToFloat()
	scope.FuncMap["__builtin_IntToString"] = builtinIntToString()
	scope.FuncMap["__builtin_StringToInt"] = builtinStringToInt()
	scope.FuncMap["__builtin_StringToFloat"] = builtinStringToFloat()
	scope.FuncMap["__builtin_StringToBool"] = builtinStringToBool()

	// Math operations
	scope.FuncMap["__builtin_IntMath"] = builtinIntMath()
	scope.FuncMap["__builtin_FloatMath"] = builtinFloatMath()
	return scope
}

func builtinFloatMath() ast.Function {
	return ast.Function{
		ArgTypes:     []ast.Type{ast.TInt},
		Variadic:     true,
		VariadicType: ast.TFloat,
		ReturnType:   ast.TFloat,
		Callback: func(args []interface{}) (interface{}, error) {
			op := args[0].(ast.ArithmeticOp)
			result := args[1].(float64)
			for _, raw := range args[2:] {
				arg := raw.(float64)
				switch op {
				case ast.ArithmeticOpAdd:
					result += arg
				case ast.ArithmeticOpSub:
					result -= arg
				case ast.ArithmeticOpMul:
					result *= arg
				case ast.ArithmeticOpDiv:
					result /= arg
				}
			}

			return result, nil
		},
	}
}

func builtinIntMath() ast.Function {
	return ast.Function{
		ArgTypes:     []ast.Type{ast.TInt},
		Variadic:     true,
		VariadicType: ast.TInt,
		ReturnType:   ast.TInt,
		Callback: func(args []interface{}) (interface{}, error) {
			op := args[0].(ast.ArithmeticOp)
			result := args[1].(int)
			for _, raw := range args[2:] {
				arg := raw.(int)
				switch op {
				case ast.ArithmeticOpAdd:
					result += arg
				case ast.ArithmeticOpSub:
					result -= arg
				case ast.ArithmeticOpMul:
					result *= arg
				case ast.ArithmeticOpDiv:
					if arg == 0 {
						return nil, errors.New("divide by zero")
					}

					result /= arg
				case ast.ArithmeticOpMod:
					if arg == 0 {
						return nil, errors.New("divide by zero")
					}

					result = result % arg
				}
			}

			return result, nil
		},
	}
}

func builtinFloatToInt() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TFloat},
		ReturnType: ast.TInt,
		Callback: func(args []interface{}) (interface{}, error) {
			return int(args[0].(float64)), nil
		},
	}
}

func builtinFloatToString() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TFloat},
		ReturnType: ast.TString,
		Callback: func(args []interface{}) (interface{}, error) {
			return strconv.FormatFloat(
				args[0].(float64), 'g', -1, 64), nil
		},
	}
}

func builtinBoolToString() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TBool},
		ReturnType: ast.TString,
		Callback: func(args []interface{}) (interface{}, error) {
			return strconv.FormatBool(args[0].(bool)), nil
		},
	}
}

func builtinIntToFloat() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TInt},
		ReturnType: ast.TFloat,
		Callback: func(args []interface{}) (interface{}, error) {
			return float64(args[0].(int)), nil
		},
	}
}

func builtinIntToString() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TInt},
		ReturnType: ast.TString,
		Callback: func(args []interface{}) (interface{}, error) {
			return strconv.FormatInt(int64(args[0].(int)), 10), nil
		},
	}
}

func builtinStringToInt() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TInt},
		ReturnType: ast.TString,
		Callback: func(args []interface{}) (interface{}, error) {
			v, err := strconv.ParseInt(args[0].(string), 0, 0)
			if err != nil {
				return nil, err
			}

			return int(v), nil
		},
	}
}

func builtinStringToFloat() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TString},
		ReturnType: ast.TFloat,
		Callback: func(args []interface{}) (interface{}, error) {
			v, err := strconv.ParseFloat(args[0].(string), 64)
			if err != nil {
				return nil, err
			}

			return v, nil
		},
	}
}

func builtinStringToBool() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TString},
		ReturnType: ast.TBool,
		Callback: func(args []interface{}) (interface{}, error) {
			v, err := strconv.ParseBool(args[0].(string))
			if err != nil {
				return nil, err
			}

			return v, nil
		},
	}
}
