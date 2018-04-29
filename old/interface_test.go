package bvmCreate

import (
	"fmt"
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestStringToOpcode(t *testing.T) {
	str := "01"
	bvmUtils.Assert(t, stringToOpcode(str) == 1, fmt.Sprintf("wrong opcode: %b", stringToOpcode(str)))
}

func TestAddBytecode(t *testing.T) {
	bvm, _ := CreateBVM("example.bvm", nil, nil, nil, nil)
	bvmUtils.AssertNow(t, bvm != nil, "bvm is nil")
	err := bench.AddBytecode("PUSH", byte(1), byte(1))
	bvmUtils.AssertNow(t, err == nil, "PUSH error")
	bvmUtils.Assert(t, bench.Input.Code().Size() == 3,
		fmt.Sprintf("wrong bytecode length %d not %d", bench.Input.Code().Size(), 3))
	bvmUtils.Assert(t, bench.NumOpcodes == 1, "wrong opcodes length")
	err = bench.AddBytecode("TEST", byte(1))
	bvmUtils.AssertNow(t, err == nil, "TEST error")
	bvmUtils.Assert(t, bench.Input.Code().Size() == 5,
		fmt.Sprintf("wrong bytecode length %d not %d", bench.Input.Code().Size(), 5))
	bvmUtils.Assert(t, bench.NumOpcodes == 2, "wrong opcodes length")
	err = bench.AddBytecode("ADD")
	bvmUtils.AssertNow(t, err == nil, "ADD error")
	bvmUtils.Assert(t, bench.Input.Code().Size() == 6,
		fmt.Sprintf("wrong bytecode length %d not %d", bench.Input.Code().Size(), 6))
	bvmUtils.Assert(t, bench.NumOpcodes == 3, "wrong opcodes length")
}
