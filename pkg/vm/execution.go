package vm

import (
	"fmt"
)

// Feed takes a op and executed corresponding operation.
// This can fail in various ways; e.g. type mismatch, stack overflow, etc.
func (vm *VM) Feed(op Op) error {
	switch op {
	default:
		panic(fmt.Errorf("invalid opcode: %d", op))
	}
}
