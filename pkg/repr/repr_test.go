package repr

import (
	"testing"

	"github.com/genkami/watson/pkg/vm"
)

func TestOpTableIsSurjectiveWhenModeIsA(t *testing.T) {
	ops := map[vm.Op]bool{}
	for _, op := range vm.AllOps() {
		ops[op] = true
	}
	for _, op := range opTableA {
		delete(ops, op)
	}
	for op := range ops {
		t.Errorf("%#v is not in opTable", op)
	}
}

func TestShowOpIsDefinedForAllOpsWhenModeIsA(t *testing.T) {
	for _, op := range vm.AllOps() {
		ShowOp(A, op)
	}
}
