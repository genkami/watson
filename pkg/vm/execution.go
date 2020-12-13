package vm

import (
	"errors"
	"fmt"
	"math"
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
	case Ineg:
		return vm.feedIneg()
	case Isht:
		return vm.feedIsht()
	case Itof:
		return vm.feedItof()
	case Finf:
		return vm.feedFinf()
	case Fnan:
		return vm.feedFnan()
	case Fneg:
		return vm.feedFneg()
	case Snew:
		return vm.feedSnew()
	case Sadd:
		return vm.feedSadd()
	case Onew:
		return vm.feedOnew()
	case Oadd:
		return vm.feedOadd()
	case Anew:
		return vm.feedAnew()
	case Aadd:
		return vm.feedAadd()
	case Bnew:
		return vm.feedBnew()
	case Bneg:
		return vm.feedBneg()
	case Nnew:
		return vm.feedNnew()
	case Gdup:
		return vm.feedGdup()
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

func (vm *VM) feedIneg() error {
	v, err := vm.popInt()
	if err != nil {
		return err
	}
	return vm.pushInt(-v)
}

func (vm *VM) feedIsht() error {
	b, err := vm.popInt()
	if err != nil {
		return err
	}
	a, err := vm.popInt()
	if err != nil {
		return err
	}
	if b >= 0 {
		return vm.pushInt(a << b)
	} else {
		return vm.pushInt(a >> -b)
	}
}

func (vm *VM) feedItof() error {
	n, err := vm.popInt()
	if err != nil {
		return err
	}
	return vm.pushFloat(math.Float64frombits(uint64(n)))
}

func (vm *VM) feedFinf() error {
	return vm.pushFloat(math.Inf(1))
}

func (vm *VM) feedFnan() error {
	return vm.pushFloat(math.NaN())
}

func (vm *VM) feedFneg() error {
	x, err := vm.popFloat()
	if err != nil {
		return err
	}
	return vm.pushFloat(-x)
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

func (vm *VM) feedAnew() error {
	return vm.pushArray([]*Value{})
}

func (vm *VM) feedAadd() error {
	x, err := vm.pop()
	if err != nil {
		return err
	}
	a, err := vm.popArray()
	if err != nil {
		return err
	}
	a = append(a, x.DeepCopy())
	return vm.pushArray(a)
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

func (vm *VM) feedGdup() error {
	v, err := vm.pop()
	if err != nil {
		return err
	}
	err = vm.push(v)
	if err != nil {
		return err
	}
	return vm.push(v.DeepCopy())
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

func (vm *VM) pushFloat(val float64) error {
	return vm.push(NewFloatValue(val))
}

func (vm *VM) pushString(val []byte) error {
	return vm.push(NewStringValue(val))
}

func (vm *VM) pushObject(val map[string]*Value) error {
	return vm.push(NewObjectValue(val))
}

func (vm *VM) pushArray(val []*Value) error {
	return vm.push(NewArrayValue(val))
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

func (vm *VM) popFloat() (float64, error) {
	v, err := vm.pop()
	if err != nil {
		return 0, err
	}
	if v.Kind != KFloat {
		return 0, ErrTypeMismatch
	}
	return v.Float, nil
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

func (vm *VM) popArray() ([]*Value, error) {
	v, err := vm.pop()
	if err != nil {
		return nil, err
	}
	if v.Kind != KArray {
		return nil, ErrTypeMismatch
	}
	return v.Array, nil
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
