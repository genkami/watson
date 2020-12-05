// Package vm contains a stack-based virtual machine that interprets Watson.
package vm

import (
	"fmt"
)

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
	// Integer Operations
	Inew Op = iota // push(0);
	Iinc           // v = pop(); push(v + 1);
	Iadd           // b = pop(); a = pop(); push(a + b);
	Ishl           // v = pop(); push(v << 1);

	// String Operations
	Snew // push("");

	// Nil Operations
	Nnew // push(nil);
)

// Kind is a type of Value.
type Kind int

const (
	KInt    Kind = iota // 64-bit signed integer
	KString             // string (represented as a byte array)
	KNil                // nil
)

// Value is an element of the stack.
type Value struct {
	Kind   Kind
	Int    int64
	String []byte
}

// NewIntValue creates a new Value that contains an integer.
func NewIntValue(val int64) *Value {
	return &Value{Kind: KInt, Int: val}
}

// NewStringValue creates a new Value that contains a string.
func NewStringValue(val []byte) *Value {
	return &Value{Kind: KString, String: val}
}

// NewNilValue creates a new Value that contains nil.
func NewNilValue() *Value {
	return &Value{Kind: KNil}
}

func (v *Value) DeepCopy() *Value {
	clone := &Value{Kind: v.Kind}
	switch v.Kind {
	case KInt:
		clone.Int = v.Int
	case KString:
		clone.String = make([]byte, len(v.String))
		copy(clone.String, v.String)
	case KNil:
		// nop
	default:
		panic(fmt.Errorf("unknown kind: %d", v.Kind))
	}
	return clone
}
