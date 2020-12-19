// Package vm contains a stack-based virtual machine that interprets Watson.
package vm

import (
	"fmt"

	"github.com/genkami/watson/pkg/types"
)

const (
	DefaultStackSize = 1024 // the default size of the stack
)

// VM is a virtual machine that consists of a stack of values and a pointer to the top of the stack.
type VM struct {
	stack []*types.Value
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
			v.stack = make([]*types.Value, size)
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
		vm.stack = make([]*types.Value, DefaultStackSize)
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
	Itou           // n: int = pop(); push(n interpreted as 64-bit unsigned integer)

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
	case Itou:
		return "Itou"
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
