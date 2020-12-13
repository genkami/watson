// Package repr is responsible for Watson Representation or an ASCII representation of Watson's opcodes.
package repr

import (
	"fmt"

	"github.com/genkami/watson/pkg/vm"
)

var opTable = map[byte]vm.Op{
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
}

var reversedTable map[vm.Op]byte

func init() {
	reversedTable = make(map[vm.Op]byte)
	for k, v := range opTable {
		reversedTable[v] = k
	}
}

// Returns a Op that corresponds to the given byte.
// This returns false if and only if b is not in the byte-to-op map.
func ReadOp(b byte) (op vm.Op, ok bool) {
	op, ok = opTable[b]
	return
}

// Returns a ascii representation of the given Op.
func ShowOp(op vm.Op) byte {
	if b, ok := reversedTable[op]; ok {
		return b
	}
	panic(fmt.Errorf("unknown Op: %#v\n", op))
}

func char(s string) byte {
	return []byte(s)[0]
}
