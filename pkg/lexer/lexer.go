package lexer

import (
	"io"

	"github.com/genkami/watson/pkg/vm"
)

// Lexer is responsible for converting a Watson Representation into a sequence of vm.Ops.
type Lexer struct {
	r   io.Reader
	buf [1]byte
}

// Creates a new Lexer that reads Watson Representation from r.
func NewLexer(r io.Reader) *Lexer {
	return &Lexer{r: r}
}

var opTable = map[byte]vm.Op{
	0x59: vm.Inew, // 'Y'
	0x75: vm.Iinc, // 'u'
	0x6d: vm.Ishl, // 'm'
	0x79: vm.Iadd, // 'y'
	0x3f: vm.Snew, // '?'
	0x21: vm.Sadd, // '!'
	0x7e: vm.Onew, // '~'
	0x4d: vm.Oadd, // 'M'
	0x7a: vm.Bnew, // 'z'
	0x6f: vm.Bneg, // 'o'
	0x2e: vm.Nnew, // '.'
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
		if op, ok := opTable[l.buf[0]]; ok {
			return op, nil
		}
	}
}
