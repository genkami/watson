package lexer

import (
	"bytes"
	"testing"

	"github.com/genkami/watson/pkg/vm"
)

func TestNextWithInew(t *testing.T) {
	expected := vm.Inew
	actual, err := readOne("y")
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Errorf("expected %#v but got %#v", expected, actual)
	}
}

func readOne(s string) (vm.Op, error) {
	buf := bytes.NewReader([]byte(s))
	l := NewLexer(buf)
	return l.Next()
}
