package main

import (
	"io"
	"os"

	"github.com/genkami/watson/pkg/decoder/yaml"
	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/vm"
)

func main() {
	m := vm.NewVM()
	lexer := lexer.NewLexer(os.Stdin)
	for {
		op, err := lexer.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		err = m.Feed(op)
		if err != nil {
			panic(err)
		}
	}
	v, err := m.Top()
	if err != nil {
		panic(err)
	}
	err = yaml.Decode(os.Stdout, v)
	if err != nil {
		panic(err)
	}
}
