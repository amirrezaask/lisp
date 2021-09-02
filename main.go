package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {
	//var input io.Reader
	//var err error
	//if len(os.Args) > 1 {
	//	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0644)
	//	defer f.Close()
	//	if err != nil {
	//		if os.IsNotExist(err) {
	//			input = strings.NewReader(os.Args[1])
	//		} else {
	//			panic(err)
	//		}
	//	}
	//	input = f
	//} else {
	//	input = os.Stdin
	//}
	//
	//bs, err := io.ReadAll(input)
	//if err != nil {
	//	panic(err)
	//}

	p := &Parser{code: string(`Hello`)}

	n, err := p.Parse()
	if err != nil {
		panic(err)
	}
	// spew.Dump(n)
	_ = n
	spew.Dump(n.Value.(*List))
	// e := &Evaluator{}
	// value, err := e.Eval(n)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(value)

}
