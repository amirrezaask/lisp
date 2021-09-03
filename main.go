package main

import (
	"github.com/amirrezaask/lisp/eval"
	"github.com/amirrezaask/lisp/parser"
)

func main() {
	p := &parser.Parser{Code: `(print 1 2 3)`}
	node, err := p.Parse()
	if err != nil {
		panic(err)
	}

	eval.Eval(node)
}
