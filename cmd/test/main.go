package main

import (
	"io"
	"os"

	"github.com/genkami/watson/pkg/decoder/util"
	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/vm"

	"gopkg.in/yaml.v2"
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
	obj := util.ToObject(v)
	enc := yaml.NewEncoder(os.Stdout)
	defer enc.Close()
	err = enc.Encode(obj)
	if err != nil {
		panic(err)
	}
}
