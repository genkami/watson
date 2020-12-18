// Package vm contains a stack-based virtual machine that interprets Watson.
package vm

import (
	"fmt"
	"math"
)

const (
	DefaultStackSize = 1024 // the default size of the stack
)

// VM is a virtual machine that consists of a stack of values and a pointer to the top of the stack.
type VM struct {
	stack []*Value
	sp    int
}

// VMOption provides the way to build VMs with custom configurations.
type VMOption interface {
	apply(*VM)
}

type vmOption func(*VM)

func (opt vmOption) apply(vm *VM) {
	opt(vm)
}

// WithStackSize sets the stack size of a VM to the given value.
// If given size is less than or equal to zero, DefaultStackSize will be used.
func WithStackSize(size int) VMOption {
	return vmOption(func(v *VM) {
		if size > 0 {
			v.stack = make([]*Value, size)
		}
	})
}

// Returns a new VM with its stack allocated.
// For more details see VMOption.
func NewVM(opts ...VMOption) *VM {
	vm := &VM{sp: -1}
	for _, opt := range opts {
		opt.apply(vm)
	}
	if len(vm.stack) == 0 {
		vm.stack = make([]*Value, DefaultStackSize)
	}
	return vm
}

// Op is an instruction executed by VM. Each op just manipulates the stack.
type Op int

const (
	// Integer Operations
	Inew Op = iota // push(0);
	Iinc           // v: int = pop(); push(v + 1);
	Ishl           // v: int = pop(); push(v << 1);
	Iadd           // b: int = pop(); a: int = pop(); push(a + b);
	Ineg           // v: int = pop(); push(-n);
	Isht           // b: int = pop(); a: int = pop(); push(a << b);
	Itof           // n: int = pop(); push(n interpreted as an IEEE-754 64-bit floating-point number)

	// Float Operations
	Finf // push(Inf);
	Fnan // push(NaN);
	Fneg // x: float = pop(); push(-x);

	// String Operations
	Snew // push("");
	Sadd // n: int = pop(); s: str = pop(); c = n && 0xff; push(s + c);

	// Object Operations
	Onew // push({});
	Oadd // v: any = pop(); k: str = pop(); o: obj = pop(); o[k] = v; push(o);

	// Array Operations
	Anew // push([]);
	Aadd // x: any = pop(); a: array = pop(); a.append(x); push(a);

	// Bool Operations
	Bnew // push(false);
	Bneg // b: bool = pop(); push(!b);

	// Nil Operations
	Nnew // push(nil);

	// Generic Operations
	Gdup // x: any = pop(); push(x); push(x);
	Gpop // pop();
	Gswp // a: any = pop(); b: any = pop(); push(a); push(b);

	// This can only be used to iterate over all defined Ops.
	numOps
)

// Returns all defined Ops.
func AllOps() []Op {
	ops := make([]Op, 0, numOps)
	for op := Op(0); op < numOps; op++ {
		ops = append(ops, op)
	}
	return ops
}

func (op Op) GoString() string {
	switch op {
	case Inew:
		return "Inew"
	case Iinc:
		return "Iinc"
	case Ishl:
		return "Ishl"
	case Iadd:
		return "Iadd"
	case Ineg:
		return "Ineg"
	case Isht:
		return "Isht"
	case Itof:
		return "Itof"
	case Finf:
		return "Finf"
	case Fnan:
		return "Fnan"
	case Fneg:
		return "Fneg"
	case Snew:
		return "Snew"
	case Sadd:
		return "Sadd"
	case Onew:
		return "Onew"
	case Oadd:
		return "Oadd"
	case Anew:
		return "Anew"
	case Aadd:
		return "Aadd"
	case Bnew:
		return "Bnew"
	case Bneg:
		return "Bneg"
	case Nnew:
		return "Nnew"
	case Gdup:
		return "Gdup"
	case Gpop:
		return "Gpop"
	case Gswp:
		return "Gswp"
	default:
		panic(fmt.Errorf("invalid opcode: %d", op))
	}
}

var _ fmt.GoStringer = Op(0)

// Kind is a type of Value.
type Kind int

const (
	KInt    Kind = iota // 64-bit signed integer
	KUint               // 64-bit unsigned integer
	KFloat              // IEEE-754 64-bit floating-point number
	KString             // string (represented as a byte array)
	KObject             // object (set of key-value pairs)
	KArray              // array
	KBool               // bool
	KNil                // nil

	// This can only be used to iterate over all defined Kinds.
	numKinds
)

func (k Kind) GoString() string {
	switch k {
	case KInt:
		return "Int"
	case KUint:
		return "Uint"
	case KFloat:
		return "Float"
	case KString:
		return "String"
	case KObject:
		return "Object"
	case KArray:
		return "Array"
	case KBool:
		return "Bool"
	case KNil:
		return "Nil"
	default:
		panic(fmt.Errorf("invalid kind: %d", k))
	}
}

var _ fmt.GoStringer = Kind(0)

// Value is an element of the stack.
type Value struct {
	Kind   Kind
	Int    int64
	Uint   uint64
	Float  float64
	String []byte
	Object map[string]*Value
	Array  []*Value
	Bool   bool
}

// NewIntValue creates a new Value that contains an integer.
func NewIntValue(val int64) *Value {
	return &Value{Kind: KInt, Int: val}
}

// NewUintValue creates a new Value that contains an unsigned integer.
func NewUintValue(val uint64) *Value {
	return &Value{Kind: KUint, Uint: val}
}

// NewFloatValue creates a new Value that contains a floating point number.
func NewFloatValue(val float64) *Value {
	return &Value{Kind: KFloat, Float: val}
}

// NewStringValue creates a new Value that contains a string.
func NewStringValue(val []byte) *Value {
	return &Value{Kind: KString, String: val}
}

// NewObjectValue creates a new Value that contains an object.
func NewObjectValue(val map[string]*Value) *Value {
	return &Value{Kind: KObject, Object: val}
}

// NewArrayValue creates a new value that contains an array.
func NewArrayValue(val []*Value) *Value {
	return &Value{Kind: KArray, Array: val}
}

// NewBoolValue creates a new Value that contains a bool.
func NewBoolValue(val bool) *Value {
	return &Value{Kind: KBool, Bool: val}
}

// NewNilValue creates a new Value that contains nil.
func NewNilValue() *Value {
	return &Value{Kind: KNil}
}

// IsNaN returns true if v is a NaN; otherwise it returns false.
func (v *Value) IsNaN() bool {
	return v.Kind == KFloat && math.IsNaN(v.Float)
}

func (v *Value) DeepCopy() *Value {
	clone := &Value{Kind: v.Kind}
	switch v.Kind {
	case KInt:
		clone.Int = v.Int
	case KUint:
		clone.Uint = v.Uint
	case KFloat:
		clone.Float = v.Float
	case KString:
		clone.String = make([]byte, len(v.String))
		copy(clone.String, v.String)
	case KObject:
		clone.Object = map[string]*Value{}
		for k, v := range v.Object {
			clone.Object[k] = v.DeepCopy()
		}
	case KArray:
		clone.Array = make([]*Value, 0, len(v.Array))
		for _, v := range v.Array {
			clone.Array = append(clone.Array, v.DeepCopy())
		}
	case KBool:
		clone.Bool = v.Bool
	case KNil:
		// nop
	default:
		panic(fmt.Errorf("unknown kind: %d", v.Kind))
	}
	return clone
}

func (v *Value) GoString() string {
	return fmt.Sprintf("{Kind: %#v, Value: %s}", v.Kind, v.goStringValue())
}

func (v *Value) goStringValue() string {
	switch v.Kind {
	case KInt:
		return fmt.Sprintf("%d", v.Int)
	case KUint:
		return fmt.Sprintf("%d", v.Uint)
	case KFloat:
		return fmt.Sprintf("%f", v.Float)
	case KString:
		return fmt.Sprintf("%#v", v.String)
	case KObject:
		return fmt.Sprintf("%#v", v.Object)
	case KArray:
		return fmt.Sprintf("%#v", v.Array)
	case KBool:
		return fmt.Sprintf("%t", v.Bool)
	case KNil:
		return "nil"
	default:
		panic(fmt.Errorf("invalid kind: %d", v.Kind))
	}
}

var _ fmt.GoStringer = &Value{}
