package main

import (
	"fmt"
	"io"
	"os"

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
	obj := toDumpable(v)
	enc := yaml.NewEncoder(os.Stdout)
	defer enc.Close()
	err = enc.Encode(obj)
	if err != nil {
		panic(err)
	}
}

func toDumpable(val *vm.Value) interface{} {
	switch val.Kind {
	case vm.KInt:
		return val.Int
	case vm.KString:
		return val.String
	case vm.KObject:
		obj := map[string]interface{}{}
		for k, v := range val.Object {
			obj[k] = toDumpable(v)
		}
		return obj
	case vm.KBool:
		return val.Bool
	case vm.KNil:
		return nil
	default:
		panic(fmt.Errorf("invalid kind: %d", val.Kind))
	}
}
