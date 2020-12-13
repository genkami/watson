// Package lexer provides a way to convert a byte sequence into a sequence of Watson's instructions and vice versa.
package lexer

import (
	"fmt"
	"io"

	"github.com/genkami/watson/pkg/vm"
)

// Mode is an important concept that is unique to Watson.
// It determines the correspondence between Vm's instructions and their representation.
type Mode int

const (
	A Mode = iota // A is the initial mode of the lexer. See the specification for more details.
)

var opTableA = map[byte]vm.Op{
	char("B"): vm.Inew,
	char("u"): vm.Iinc,
	char("b"): vm.Ishl,
	char("a"): vm.Iadd,
	char("A"): vm.Ineg,
	char("e"): vm.Isht,
	char("i"): vm.Itof,
	char("q"): vm.Finf,
	char("t"): vm.Fnan,
	char("p"): vm.Fneg,
	char("?"): vm.Snew,
	char("!"): vm.Sadd,
	char("~"): vm.Onew,
	char("M"): vm.Oadd,
	char("@"): vm.Anew,
	char("s"): vm.Aadd,
	char("z"): vm.Bnew,
	char("o"): vm.Bneg,
	char("."): vm.Nnew,
	char("*"): vm.Gdup,
	char("#"): vm.Gpop,
	char("%"): vm.Gswp,
}

var reversedTableA map[vm.Op]byte

func init() {
	reversedTableA = make(map[vm.Op]byte)
	for k, v := range opTableA {
		reversedTableA[v] = k
	}
}

// Returns a Op that corresponds to the given byte.
// This returns false if and only if b is not in the byte-to-op map.
func ReadOp(m Mode, b byte) (op vm.Op, ok bool) {
	op, ok = opTableA[b]
	return
}

// Returns a ascii representation of the given Op.
func ShowOp(m Mode, op vm.Op) byte {
	if b, ok := reversedTableA[op]; ok {
		return b
	}
	panic(fmt.Errorf("unknown Op: %#v\n", op))
}

func char(s string) byte {
	return []byte(s)[0]
}

// Lexer is responsible for converting a Watson Representation into a sequence of vm.Ops.
type Lexer struct {
	r   io.Reader
	buf [1]byte
}

// Creates a new Lexer that reads Watson Representation from r.
func NewLexer(r io.Reader) *Lexer {
	return &Lexer{r: r}
}

// Returns the next Op.
// This returns io.EOF if it hits on the end of the input.
func (l *Lexer) Next() (vm.Op, error) {
	for {
		_, err := l.r.Read(l.buf[:])
		if err != nil {
			// Note that it returns io.EOF if the underlying Reader returns io.EOF.
			return 0, err
		}
		if op, ok := ReadOp(A, l.buf[0]); ok {
			return op, nil
		}
	}
}
