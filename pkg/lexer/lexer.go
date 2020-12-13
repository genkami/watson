// Package lexer provides a way to convert a Watson Representation into a sequence of Watson's instructions and vice versa.
// (where Watson Representation is a sequence of character that represents Watson's instructions).
//
// The correspondence between characters and instructions depends on the lexer's mode.
// Each lexer has its own mode. The mode of a lexer is either `A` or `S`. The initial mode of a lexer is A unless otherwise specified.
//
// The complete conversion table between instructions and their Watson Representations are as follows:
//
//   +-----------+--------------+--------------+
//   |Instruction|Watson        |Watson        |
//   |           |Representation|Representation|
//   |           |(mode = A)    |(mode = S)    |
//   +-----------+--------------+--------------+
//   |Inew       |B             |S             |
//   +-----------+--------------+--------------+
//   |Iinc       |u             |h             |
//   +-----------+--------------+--------------+
//   |Ishl       |b             |a             |
//   +-----------+--------------+--------------+
//   |Iadd       |a             |k             |
//   +-----------+--------------+--------------+
//   |Ineg       |A             |r             |
//   +-----------+--------------+--------------+
//   |Isht       |e             |A             |
//   +-----------+--------------+--------------+
//   |Itof       |i             |z             |
//   +-----------+--------------+--------------+
//   |Finf       |q             |m             |
//   +-----------+--------------+--------------+
//   |Fnan       |t             |b             |
//   +-----------+--------------+--------------+
//   |Fneg       |p             |u             |
//   +-----------+--------------+--------------+
//   |Snew       |?             |$             |
//   +-----------+--------------+--------------+
//   |Sadd       |!             |-             |
//   +-----------+--------------+--------------+
//   |Onew       |~             |+             |
//   +-----------+--------------+--------------+
//   |Oadd       |M             |g             |
//   +-----------+--------------+--------------+
//   |Anew       |@             |v             |
//   +-----------+--------------+--------------+
//   |Aadd       |s             |?             |
//   +-----------+--------------+--------------+
//   |Bnew       |z             |^             |
//   +-----------+--------------+--------------+
//   |Bneg       |o             |!             |
//   +-----------+--------------+--------------+
//   |Nnew       |.             |y             |
//   +-----------+--------------+--------------+
//   |Gdup       |*             |/             |
//   +-----------+--------------+--------------+
//   |Gpop       |#             |e             |
//   +-----------+--------------+--------------+
//   |Gswp       |%             |:             |
//   +-----------+--------------+--------------+
package lexer

import (
	"fmt"
	"io"

	"github.com/genkami/watson/pkg/vm"
)

// Mode is an important concept that is unique to Watson.
// It determines the correspondence between Vm's instructions and their ASCII representation.
type Mode int

const (
	A Mode = iota // A, S are the modes of the lexer. See the overview for more details.
	S
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

var opTableS = map[byte]vm.Op{
	char("S"): vm.Inew,
	char("h"): vm.Iinc,
	char("a"): vm.Ishl,
	char("k"): vm.Iadd,
	char("r"): vm.Ineg,
	char("A"): vm.Isht,
	char("z"): vm.Itof,
	char("m"): vm.Finf,
	char("b"): vm.Fnan,
	char("u"): vm.Fneg,
	char("$"): vm.Snew,
	char("-"): vm.Sadd,
	char("+"): vm.Onew,
	char("g"): vm.Oadd,
	char("v"): vm.Anew,
	char("?"): vm.Aadd,
	char("^"): vm.Bnew,
	char("!"): vm.Bneg,
	char("y"): vm.Nnew,
	char("/"): vm.Gdup,
	char("e"): vm.Gpop,
	char(":"): vm.Gswp,
}

var reversedTableS map[vm.Op]byte

func init() {
	reversedTableA = make(map[vm.Op]byte)
	for k, v := range opTableA {
		reversedTableA[v] = k
	}
	reversedTableS = make(map[vm.Op]byte)
	for k, v := range opTableS {
		reversedTableS[v] = k
	}
}

func readOp(m Mode, b byte) (op vm.Op, ok bool) {
	var table map[byte]vm.Op
	switch m {
	case A:
		table = opTableA
	case S:
		table = opTableS
	default:
		panic(fmt.Errorf("unknown mode: %d", m))
	}
	op, ok = table[b]
	return
}

func showOp(m Mode, op vm.Op) byte {
	if b, ok := reversedTableA[op]; ok {
		return b
	}
	panic(fmt.Errorf("unknown Op: %#v\n", op))
}

func char(s string) byte {
	return []byte(s)[0]
}

// Lexer converts a Watson Representation into a sequence of `vm.Op`s.
// Each lexer has its state called mode. Its default mode is A, and whenever it yields the `Snew` instruction, it flips its mode.
//
// Example:
// Consider the situation where the lexer tries to read the following string:
//   b?b$q
// As described above, the lexer's initial mode is A. The lexer first hits 'b' and regards it as `Ishl`.
// Then it hits the character '?', where it changes its mode from A to S. More specifically, the lexer reads a character '?' and yields `Snew` since its current state is A. Then it changes its current state to S.
// After that, it hits 'b' again, but in this time the 'b' is interpreted differently from the previous lexing step. Since the current mode of the lexer is S, it regards 'b' as `Fnan` instead of `Ishl`.
// Then it hits '?', which is now interpreted as `Snew`, yields `Snew`, and changes its current mode to A.
// In the end, it hits 'q' and yields `Finf`, and it stops its lexing procedure.
type Lexer struct {
	r    io.Reader
	mode Mode
	buf  [1]byte
}

// Creates a new Lexer that reads Watson Representation from r.
func NewLexer(r io.Reader) *Lexer {
	return &Lexer{r: r, mode: A}
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
		if op, ok := readOp(l.mode, l.buf[0]); ok {
			l.mode = nextMode(l.mode, op)
			return op, nil
		}
	}
}

func nextMode(mode Mode, op vm.Op) Mode {
	var next Mode
	switch mode {
	case A:
		if op == vm.Snew {
			next = S
		} else {
			next = A
		}
	case S:
		if op == vm.Snew {
			next = A
		} else {
			next = S
		}
	default:
		panic(fmt.Errorf("unknown mode: %d", mode))
	}
	return next
}
