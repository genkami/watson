package prettifier

import (
	"bytes"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/vm"
)

func TestPrettifierDoesNotChangeSemantics(t *testing.T) {
	test := func(src string) {
		orig, err := lex(src)
		if err != nil {
			t.Fatal(err)
		}
		prettified, err := prettify(orig)
		if err != nil {
			t.Fatal(err)
		}
		origResult, err := execute(orig)
		if err != nil {
			t.Fatal(err)
		}
		prettifiedResult, err := execute(prettified)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(origResult, prettifiedResult); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
	test("B")
	test("?SShak")
	test("?SShaShaAk")
}

func lex(src string) ([]vm.Op, error) {
	ops := make([]vm.Op, 0, len(src))
	l := lexer.NewLexer(bytes.NewReader([]byte(src)))
	for {
		tok, err := l.Next()
		if err == io.EOF {
			return ops, nil
		} else if err != nil {
			return nil, err
		}
		ops = append(ops, tok.Op)
	}
}

func prettify(orig []vm.Op) ([]vm.Op, error) {
	sw := lexer.NewSliceWriter()
	p := NewPrettifier(sw)
	for _, op := range orig {
		err := p.Write(op)
		if err != nil {
			return nil, err
		}
	}
	return sw.Ops(), nil
}

func execute(ops []vm.Op) (*vm.Value, error) {
	v := vm.NewVM()
	for _, op := range ops {
		err := v.Feed(op)
		if err != nil {
			return nil, err
		}
	}
	return v.Top()
}
