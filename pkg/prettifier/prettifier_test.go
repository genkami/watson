package prettifier

import (
	"bytes"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/types"
	"github.com/genkami/watson/pkg/vm"
)

func TestPrettifierDoesNotChangeSemantics(t *testing.T) {
	test := func(src string, expected string) {
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
		actual, err := unlex(prettified)
		if err != nil {
			t.Fatal(err)
		}
		if expected != actual {
			t.Errorf("expected %#v but got %#v", expected, actual)
		}
	}
	test("B", "B")
	test("?SShak", "?SSharrk")
	test("?SShaShaAk", "?SShaShaArrk")
	test("?+", "?Samee+")
	test("~?$#zM", "~?$#zooM")
	test("~?$#BM", "~?$#BAAME#")
	test("~?$#BuM", "~?$#BuAAME#")
	test("~?$#BBaM", "~?$#BBaAAME#")
	test("~?$#BAM", "~?$#BAAAME#")
	test("~?$#BBeM", "~?$#BBeAAME#")
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

func execute(ops []vm.Op) (*types.Value, error) {
	v := vm.NewVM()
	for _, op := range ops {
		err := v.Feed(op)
		if err != nil {
			return nil, err
		}
	}
	return v.Top()
}

func unlex(ops []vm.Op) (string, error) {
	buf := bytes.NewBuffer(nil)
	ul := lexer.NewUnlexer(buf)
	for _, op := range ops {
		err := ul.Write(op)
		if err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}
