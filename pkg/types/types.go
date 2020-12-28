// Package types proviedes types that can be represented as Watson.
package types

import (
	"fmt"
	"math"
)

// Value is an arbitrary value that can be represented as Watson.
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
	return &Value{Kind: Int, Int: val}
}

// NewUintValue creates a new Value that contains an unsigned integer.
func NewUintValue(val uint64) *Value {
	return &Value{Kind: Uint, Uint: val}
}

// NewFloatValue creates a new Value that contains a floating point number.
func NewFloatValue(val float64) *Value {
	return &Value{Kind: Float, Float: val}
}

// NewStringValue creates a new Value that contains a string.
func NewStringValue(val []byte) *Value {
	return &Value{Kind: String, String: val}
}

// NewObjectValue creates a new Value that contains an object.
func NewObjectValue(val map[string]*Value) *Value {
	return &Value{Kind: Object, Object: val}
}

// NewArrayValue creates a new value that contains an array.
func NewArrayValue(val []*Value) *Value {
	return &Value{Kind: Array, Array: val}
}

// NewBoolValue creates a new Value that contains a bool.
func NewBoolValue(val bool) *Value {
	return &Value{Kind: Bool, Bool: val}
}

// NewNilValue creates a new Value that contains nil.
func NewNilValue() *Value {
	return &Value{Kind: Nil}
}

// IsNaN returns true if v is a NaN; otherwise it returns false.
func (v *Value) IsNaN() bool {
	return v.Kind == Float && math.IsNaN(v.Float)
}

// DeepCopy returns a deep copy of v.
func (v *Value) DeepCopy() *Value {
	clone := &Value{Kind: v.Kind}
	switch v.Kind {
	case Int:
		clone.Int = v.Int
	case Uint:
		clone.Uint = v.Uint
	case Float:
		clone.Float = v.Float
	case String:
		clone.String = make([]byte, len(v.String))
		copy(clone.String, v.String)
	case Object:
		clone.Object = map[string]*Value{}
		for k, v := range v.Object {
			clone.Object[k] = v.DeepCopy()
		}
	case Array:
		clone.Array = make([]*Value, 0, len(v.Array))
		for _, v := range v.Array {
			clone.Array = append(clone.Array, v.DeepCopy())
		}
	case Bool:
		clone.Bool = v.Bool
	case Nil:
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
	case Int:
		return fmt.Sprintf("%d", v.Int)
	case Uint:
		return fmt.Sprintf("%d", v.Uint)
	case Float:
		return fmt.Sprintf("%f", v.Float)
	case String:
		return fmt.Sprintf("%#v", v.String)
	case Object:
		return fmt.Sprintf("%#v", v.Object)
	case Array:
		return fmt.Sprintf("%#v", v.Array)
	case Bool:
		return fmt.Sprintf("%t", v.Bool)
	case Nil:
		return "nil"
	default:
		panic(fmt.Errorf("invalid kind: %d", v.Kind))
	}
}

var _ fmt.GoStringer = &Value{}

// Kind is a type of Value.
type Kind int

const (
	Int    Kind = iota // 64-bit signed integer
	Uint               // 64-bit unsigned integer
	Float              // IEEE-754 64-bit floating-point number
	String             // string (represented as a byte array)
	Object             // object (set of key-value pairs)
	Array              // array
	Bool               // bool
	Nil                // nil
)

func (k Kind) GoString() string {
	switch k {
	case Int:
		return "Int"
	case Uint:
		return "Uint"
	case Float:
		return "Float"
	case String:
		return "String"
	case Object:
		return "Object"
	case Array:
		return "Array"
	case Bool:
		return "Bool"
	case Nil:
		return "Nil"
	default:
		panic(fmt.Errorf("invalid kind: %d", k))
	}
}

var _ fmt.GoStringer = Kind(0)

// By implementing Marshaler you can configure converting go objects into Values.
type Marshaler interface {
	MarshalWatson() (*Value, error)
}

// By implementing Unmarshaler you can configure converting Values into go objects.
type Unmarshaler interface {
	UnmarshalWatson(*Value) error
}
