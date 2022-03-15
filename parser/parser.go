package parser

import (
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	Code string
}

func (p *Parser) tokenize() []string {
	p.Code = strings.Replace(p.Code, "(", " ( ", -1)
	p.Code = strings.Replace(p.Code, ")", " ) ", -1)
	p.Code = strings.Replace(p.Code, "{", " { ", -1)
	p.Code = strings.Replace(p.Code, "}", " } ", -1)
	p.Code = strings.Replace(p.Code, "[", " [ ", -1)
	p.Code = strings.Replace(p.Code, "]", " ] ", -1)
	return strings.Split(p.Code, " ")
}

func (p *Parser) removeWhiteSpaces(tokens []string) []string {
	ts := []string{}
	for _, t := range tokens {
		if t != "" && t != "\n" {
			ts = append(ts, t)
		}
	}
	return ts
}

func pop(list *[]string) string {
	poped := (*list)[0]
	*list = (*list)[1:]
	return poped
}

func isNumber(t string) bool {
	for _, c := range t {
		if !(c >= '0' && c <= '9') {
			return false
		}
	}
	return true
}
func isString(t string) bool {
	return (t[0] == '"' && t[len(t)-1] == '"') || (t[0] == '\'' && t[len(t)-1] == '\'')
}
func isAtom(t string) bool {
	return !(t[0] >= '0' && t[0] <= '9')
}
func (p *Parser) parenthesize(tokens []string, curr *List) (*List, error) {
	if len(tokens) == 0 {
		return curr, nil
	}
	t := pop(&tokens)
	if t == "(" {
		if curr == nil {
			return p.parenthesize(tokens, &List{})
		}
		newL, err := p.parenthesize(tokens, curr)
		if err != nil {
			return nil, err
		}
		curr.Value = append(curr.Value, &Node{Type: NodeType_List, Value: newL})
		return p.parenthesize(tokens, curr)
	} else if t == ")" {
		return p.parenthesize(tokens, curr)
	} else {
		if isNumber(t) {
			num, _ := strconv.Atoi(t)
			curr.Value = append(curr.Value, &Node{Type: NodeType_Number, Value: num})
			return p.parenthesize(tokens, curr)
		} else if isString(t) {
			curr.Value = append(curr.Value, &Node{Type: NodeType_String, Value: t[1 : len(t)-1]})
			return p.parenthesize(tokens, curr)
		} else {
			curr.Value = append(curr.Value, &Node{Type: NodeType_Atom, Value: t})
			return p.parenthesize(tokens, curr)
		}
	}
}

func (p *Parser) Parse() (*Node, error) {
	if p.Code[0] != '(' && p.Code[len(p.Code)-1] != ')' {
		if isNumber(p.Code) {
			num, _ := strconv.Atoi(p.Code)
			return &Node{Type: NodeType_Number, Value: num}, nil
		} else if isString(p.Code) {
			return &Node{Type: NodeType_String, Value: p.Code[1 : len(p.Code)-1]}, nil
		} else if isAtom(p.Code) {
			return &Node{Type: NodeType_Atom, Value: p.Code}, nil
		} else {
			return nil, fmt.Errorf("code is malformed")
		}
	} else {
		tokens := p.tokenize()
		tokens = p.removeWhiteSpaces(tokens)
		list, err := p.parenthesize(tokens, nil)
		return &Node{Type: NodeType_List, Value: list}, err
	}

}
