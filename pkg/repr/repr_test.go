package repr

import (
	"testing"

	"github.com/genkami/watson/pkg/vm"
)

var table = map[string]vm.Op{
	"B": vm.Inew,
	"u": vm.Iinc,
	"b": vm.Ishl,
	"a": vm.Iadd,
	"?": vm.Snew,
	"!": vm.Sadd,
	"~": vm.Onew,
	"M": vm.Oadd,
	"z": vm.Bnew,
	"o": vm.Bneg,
	".": vm.Nnew,
}

func TestReadOp(t *testing.T) {
	for s, expected := range table {
		actual, ok := ReadOp([]byte(s)[0])
		if !ok {
			t.Fatalf("op not found for %s", s)
		}
		if expected != actual {
			t.Errorf("expected %#v but got %#v", expected, actual)
		}
	}
}

func TestShowOp(t *testing.T) {
	for s, op := range table {
		expected := []byte(s)[0]
		actual := ShowOp(op)
		if expected != actual {
			t.Errorf("expected %#v but got %#v", expected, actual)
		}
	}
}
