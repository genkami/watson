package vm

import (
	"testing"
)

func TestNewVMWithStackSize(t *testing.T) {
	size := 123
	vm := NewVM(WithStackSize(size))
	if len(vm.stack) != size {
		t.Fatalf("expected stack size to be %d, but got %d", size, len(vm.stack))
	}
}

func TestNewVMWithZeroStackSize(t *testing.T) {
	vm := NewVM(WithStackSize(0))
	if len(vm.stack) != DefaultStackSize {
		t.Fatalf("expected stack size to be %d, but got %d", DefaultStackSize, len(vm.stack))
	}
}

func TestNewVMWithNegativeStackSize(t *testing.T) {
	vm := NewVM(WithStackSize(-1))
	if len(vm.stack) != DefaultStackSize {
		t.Fatalf("expected stack size to be %d, but got %d", DefaultStackSize, len(vm.stack))
	}
}

func TestGoStringIsDefinedForAllOps(t *testing.T) {
	for _, op := range AllOps() {
		op.GoString()
	}
}
