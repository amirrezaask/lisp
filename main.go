package main

import (
	"fmt"
	"github.com/amirrezaask/lisp/eval"
	"github.com/amirrezaask/lisp/parser"
)

func main() {
	p := &parser.Parser{Code: `(print 1)`}
	node, err := p.Parse()
	if err != nil {
		panic(err)
	}

	eval.Eval(node, &eval.SymbolTable{
		Data: map[string]*eval.Value{
			"print": &eval.Value{
				Type:  eval.ValueType_Function,
				Value: &eval.Function{
					Callable: func(m map[string]interface{}) *eval.Value {
						fmt.Println(m["arg"])
						return nil
					},
					Args:     []string{
						"arg",
					},
				},
			},
		},
	})
}
