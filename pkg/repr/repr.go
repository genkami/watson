// Package repr is responsible for Watson Representation or an ASCII representation of Watson's opcodes.
package repr

import (
	"fmt"

	"github.com/genkami/watson/pkg/vm"
)

var opTable = map[byte]vm.Op{
	0x42: vm.Inew, // 'B'
	0x75: vm.Iinc, // 'u'
	0x62: vm.Ishl, // 'b'
	0x61: vm.Iadd, // 'a'
	0x41: vm.Ineg, // 'A'
	0x3f: vm.Snew, // '?'
	0x21: vm.Sadd, // '!'
	0x7e: vm.Onew, // '~'
	0x4d: vm.Oadd, // 'M'
	0x7a: vm.Bnew, // 'z'
	0x6f: vm.Bneg, // 'o'
	0x2e: vm.Nnew, // '.'
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
