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

// Feed takes a op and executes corresponding operation.
// This can fail in various ways; e.g. type mismatch, stack overflow, etc.
func (vm *VM) Feed(op Op) error {
	switch op {
	case Inew:
		return vm.feedInew()
	case Iinc:
		return vm.feedIinc()
	case Ishl:
		return vm.feedIshl()
	case Iadd:
		return vm.feedIadd()
	case Snew:
		return vm.feedSnew()
	case Sadd:
		return vm.feedSadd()
	case Onew:
		return vm.feedOnew()
	case Oadd:
		return vm.feedOadd()
	case Bnew:
		return vm.feedBnew()
	case Bneg:
		return vm.feedBneg()
	case Nnew:
		return vm.feedNnew()
	default:
		panic(fmt.Errorf("invalid opcode: %d", op))
	}
}

// FeedMulti takes a series of Ops and executes them sequentially.
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

func (vm *VM) feedIshl() error {
	v, err := vm.popInt()
	if err != nil {
		return err
	}
	return vm.pushInt(v << 1)
}

func (vm *VM) feedIadd() error {
	b, err := vm.popInt()
	if err != nil {
		return err
	}
	a, err := vm.popInt()
	if err != nil {
		return err
	}
	return vm.pushInt(a + b)
}

func (vm *VM) feedSnew() error {
	return vm.pushString([]byte{})
}

func (vm *VM) feedSadd() error {
	n, err := vm.popInt()
	if err != nil {
		return err
	}
	s, err := vm.popString()
	if err != nil {
		return err
	}
	t := append(s, byte(n))
	return vm.pushString(t)
}

func (vm *VM) feedOnew() error {
	return vm.pushObject(map[string]*Value{})
}

func (vm *VM) feedOadd() error {
	v, err := vm.pop()
	if err != nil {
		return err
	}
	k, err := vm.popString()
	if err != nil {
		return err
	}
	o, err := vm.popObject()
	if err != nil {
		return err
	}
	o[string(k)] = v.DeepCopy()
	return vm.pushObject(o)
}

func (vm *VM) feedBnew() error {
	return vm.pushBool(false)
}

func (vm *VM) feedBneg() error {
	v, err := vm.popBool()
	if err != nil {
		return err
	}
	return vm.pushBool(!v)
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

func (vm *VM) pushString(val []byte) error {
	return vm.push(NewStringValue(val))
}

func (vm *VM) pushObject(val map[string]*Value) error {
	return vm.push(NewObjectValue(val))
}

func (vm *VM) pushBool(val bool) error {
	return vm.push(NewBoolValue(val))
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

func (vm *VM) popString() ([]byte, error) {
	v, err := vm.pop()
	if err != nil {
		return nil, err
	}
	if v.Kind != KString {
		return nil, ErrTypeMismatch
	}
	return v.String, nil
}

func (vm *VM) popObject() (map[string]*Value, error) {
	v, err := vm.pop()
	if err != nil {
		return nil, err
	}
	if v.Kind != KObject {
		return nil, ErrTypeMismatch
	}
	return v.Object, nil
}

func (vm *VM) popBool() (bool, error) {
	v, err := vm.pop()
	if err != nil {
		return false, err
	}
	if v.Kind != KBool {
		return false, ErrTypeMismatch
	}
	return v.Bool, nil
}
