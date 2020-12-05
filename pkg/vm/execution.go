package vm

import (
	"errors"
	"fmt"
)

var (
	ErrStackEmpty               = errors.New("stack is empty")
	ErrMaximumStackSizeExceeded = errors.New("maximum stack size exceeded")
)

// Top returns a value in the top of the stack.
// This returns ErrStackEmpty if the stack is empty.
func (vm *VM) Top() (*Value, error) {
	if vm.sp < 0 {
		return nil, ErrStackEmpty
	}
	return vm.stack[vm.sp], nil
}

// Feed takes a op and executed corresponding operation.
// This can fail in various ways; e.g. type mismatch, stack overflow, etc.
func (vm *VM) Feed(op Op) error {
	switch op {
	case Inew:
		return vm.feedInew()
	default:
		panic(fmt.Errorf("invalid opcode: %d", op))
	}
}

func (vm *VM) feedInew() error {
	return vm.pushInt(0)
}

// Miscellaneous functions

func (vm *VM) pop() (*Value, error) {
	if vm.sp < 0 {
		return nil, ErrStackEmpty
	}
	top := vm.stack[vm.sp]
	vm.stack[vm.sp] = nil
	vm.sp--
	return top, nil
}

func (vm *VM) push(v *Value) error {
	if len(vm.stack)-1 <= vm.sp {
		return ErrMaximumStackSizeExceeded
	}
	vm.sp++
	vm.stack[vm.sp] = v
	return nil
}

func (vm *VM) pushInt(val int64) error {
	return vm.push(NewIntValue(val))
}
