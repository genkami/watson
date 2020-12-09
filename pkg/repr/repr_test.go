package repr

import (
	"testing"

	"github.com/genkami/watson/pkg/vm"
)

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

func TestShowOpIsDefinedForAllOps(t *testing.T) {
	for _, op := range vm.AllOps() {
		ShowOp(op)
	}
}
