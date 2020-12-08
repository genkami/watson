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

func TestOpTableIsSurjective(t *testing.T) {
	ops := map[vm.Op]bool{}
	for _, op := range vm.AllOps() {
		ops[op] = true
	}
	for _, op := range opTable {
		delete(ops, op)
	}
	for op := range ops {
		t.Errorf("%#v is not in opTable", op)
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

func TestShowOpIsDefinedForAllOps(t *testing.T) {
	for _, op := range vm.AllOps() {
		ShowOp(op)
	}
}
