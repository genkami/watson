package vm

import (
	"errors"
	"fmt"
)

var (
	ErrStackEmpty               = errors.New("stack is empty")
	ErrMaximumStackSizeExceeded = errors.New("maximum stack size exceeded")
	ErrTypeMismatch             = errors.New("type mismatch")
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
	case Iinc:
		return vm.feedIinc()
	case Nnew:
		return vm.feedNnew()
	default:
		panic(fmt.Errorf("invalid opcode: %d", op))
	}
}

// FeedMulti takes a series of Ops and execute them sequentially.
// If one of them fails, it stops execution and returns an error.
func (vm *VM) FeedMulti(ops []Op) error {
	for _, op := range ops {
		if err := vm.Feed(op); err != nil {
			return err
		}
	}
	return nil
}

func (vm *VM) feedInew() error {
	return vm.pushInt(0)
}

func (vm *VM) feedIinc() error {
	v, err := vm.popInt()
	if err != nil {
		return err
	}
	return vm.pushInt(v + 1)
}

func (vm *VM) feedNnew() error {
	return vm.pushNil()
}

//
// Miscellaneous functions
//

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

func (vm *VM) pushNil() error {
	return vm.push(NewNilValue())
}

func (vm *VM) pop() (*Value, error) {
	if vm.sp < 0 {
		return nil, ErrStackEmpty
	}
	top := vm.stack[vm.sp]
	vm.stack[vm.sp] = nil
	vm.sp--
	return top, nil
}

func (vm *VM) popInt() (int64, error) {
	v, err := vm.pop()
	if err != nil {
		return 0, err
	}
	if v.Kind != KInt {
		return 0, ErrTypeMismatch
	}
	return v.Int, nil
}
