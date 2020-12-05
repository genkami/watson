// Package vm contains a stack-based virtual machine that interprets Watson.
package vm

const (
	DefaultStackSize = 1024 // the default size of the stack
)

// VM is a virtual machine that consists of a stack of values and a pointer to the top of the stack.
//
// In this document, we call the N-th element of the stack s[N] (counting from the top).
type VM struct {
	stack []*Value
	sp    int
}

// Returns a new VM with its stack allocated.
func NewVM() *VM {
	return &VM{
		stack: make([]*Value, DefaultStackSize),
		sp:    -1,
	}
}

// Op is an instruction executed by VM. Each op just manipulates the stack.
type Op int

const (
	// Every op
	Inew Op = iota // push(0);
	Iinc           // v = pop(); push(v+1);
	Iadd           // b = pop(); a = pop(); push(a + b);
)

// Value is an element of the stack.
type Value struct {
	Kind Kind
	Int  int64
}

// Kind is a type of Value.
type Kind int

const (
	KInt Kind = iota // 64-bit signed integer
)
