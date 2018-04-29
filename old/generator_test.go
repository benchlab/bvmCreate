package bvmCreate

import (
	"fmt"
	"log"
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestNextSpec(t *testing.T) {
	bvm, _ := CreateBVM("example.bvm", nil, deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bench.Input = new(BasicInput)
	bytes := []byte("0101")
	bench.Input.Code().Append(FromHexBytes(bytes)...)
	i := bench.nextSpec()
	bvmUtils.AssertNow(t, i != nil, "next spec shouldn't be nil")
	bvmUtils.AssertNow(t, i.Mnemonic == "ADD", "wrong mnemonic")
}

func TestNextOpcode(t *testing.T) {
	bvm, _ := CreateBVM("example.bvm", nil, deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bench.Input = new(BasicInput)
	bytes := []byte("0101")
	bench.Input.Code().Append(FromHexBytes(bytes)...)
	bvmUtils.Assert(t, bench.Input.Code().Size() == 2, "wrong code size")
	log.Println(bench.Input.Code().Data)
	n := bench.nextOpcode()
	bvmUtils.Assert(t, n == 1, fmt.Sprintf("Next opcode was: %d\n", n))
}

func TestVersion(t *testing.T) {
	bvmUtils.Assert(t, version() != "", "version shouldn't be empty")
}
