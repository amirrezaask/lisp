package eval

import (
	"fmt"

	"github.com/amirrezaask/lisp/parser"
)

type SymbolTable struct {
	Data map[string]*Value
}

func (s *SymbolTable) Get(key string) (*Value, error) {
	val, exists := s.Data[key]
	if !exists {
		return nil, fmt.Errorf("not exists")
	}
	return val, nil
}

type ValueType string

const (
	ValueType_String   = "string"
	ValueType_Number   = "number"
	ValueType_Atom     = "atom"
	ValueType_List     = "list"
	ValueType_Function = "function"
)

type Value struct {
	Type  ValueType
	Value interface{}
}

type Function struct {
	Callable func(map[string]interface{}) *Value
	Args     []string
}

func returnNodeValue(node *parser.Node) (*Value, error) {
	if node.Type == parser.NodeTypeAtom {
		return &Value{Type: ValueType_Atom, Value: node.Value}, nil
	} else if node.Type == parser.NodeTypeNumber {
		return &Value{Type: ValueType_Number, Value: node.Value}, nil
	} else if node.Type == parser.NodeTypeString {
		return &Value{Type: ValueType_String, Value: node.Value}, nil
	} else {
		return nil, fmt.Errorf("not supported type for evaluation: %s", node.Type)
	}

}

func Eval(node *parser.Node, st *SymbolTable) (*Value, error) {
	if node.Type != parser.NodeTypeList {
		return returnNodeValue(node)
	}
	list := node.Value.(*parser.List).Value
	fn := list[0]
	args := list[1:]
	argsEvaluated := []*Value{}
	for _, arg := range args {
		val, err := Eval(arg, st)
		if err != nil {
			return nil, err
		}
		argsEvaluated = append(argsEvaluated, val)
	}
	//fn should be atom
	if fn.Type != parser.NodeTypeAtom {
		return nil, fmt.Errorf("function name should be of type atom")
	}
	fnVal, err := st.Get(fn.Value.(string))
	if err != nil {
		return nil, err
	}
	fnFn := fnVal.Value.(*Function)
	argsMap := map[string]interface{}{}
	for idx, arg := range fnFn.Args {
		argsMap[arg] = argsEvaluated[idx].Value
	}
	return fnFn.Callable(argsMap), nil
}
