package lexer

import (
	"bytes"
	"testing"

	"github.com/genkami/watson/pkg/vm"
)

func TestOpTable(t *testing.T) {
	table := map[string]vm.Op{
		"Y": vm.Inew,
		"u": vm.Iinc,
		"m": vm.Ishl,
		"y": vm.Iadd,
		"?": vm.Snew,
		"!": vm.Sadd,
		"~": vm.Onew,
		"M": vm.Oadd,
		"z": vm.Bnew,
		"o": vm.Bneg,
		".": vm.Nnew,
	}
	for s, expected := range table {
		actual, err := readOne(s)
		if err != nil {
			t.Fatalf("error on testing %s (%#v): %#v", s, expected, err)
		}
		if expected != actual {
			t.Errorf("expected %#v but got %#v", expected, actual)
		}
	}
}

func readOne(s string) (vm.Op, error) {
	buf := bytes.NewReader([]byte(s))
	l := NewLexer(buf)
	return l.Next()
}
