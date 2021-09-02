package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var input io.ReadCloser
	var err error
	if len(os.Args) > 1 {
		input, err = os.OpenFile(os.Args[1], os.O_RDONLY, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		input = os.Stdin
	}

	bs, err := io.ReadAll(input)
	if err != nil {
		panic(err)
	}

	p := &Parser{}

	l, err := p.Parse(bs)
	if err != nil {
		panic(err)
	}
	e := &Evaluator{}
	value, err := e.Eval(l)
	if err != nil {
		panic(err)
	}

	fmt.Println(value)

}
