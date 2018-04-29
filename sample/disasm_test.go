package sample

import (
	"math/big"
	"testing"

	"github.com/benchlab/bvmCreate"
)

func TestDisasm(t *testing.T) {
	bvm := newBVM(map[string]bvmCreate.DisasmFunction{
		"PUSH": func(bench *bvmCreate.BVM, offset int, bytecode []byte) ([]string, int) {
			size := bench.AssignedParameters["Size Byte"]
			intSize := new(big.Int).SetBytes(size).Int64()
			return []string{string(size), string(bytecode[offset+1 : int64(offset+1)+intSize])}, int(intSize) + 1
		},
	})
	bench.DisasmString("010101")
}
