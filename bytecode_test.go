package bvmCreate

import (
	"fmt"
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestBytecodeAddCommand(t *testing.T) {
	b := new(Bcode)
	c := Command{
		mnemonic: "ADD",
	}
	b.AddCommand(&c)
	bvmUtil.Assert(t, b.Length() == 1, "BVM Error! Wrong length")
}

func TestBytecodeAdd(t *testing.T) {
	b := new(Bcode)
	b.Add("ADD")
	bvmUtil.Assert(t, b.Length() == 1, "BVM Error! Wrong length.")
}

func TestBytecodeConcat(t *testing.T) {
	b := new(Bcode)
	b.Add("ADD")
	bvmUtil.Assert(t, b.Length() == 1, "BVM Error! Length of b is wrong.")

	o := Bcode{}
	o.Add("ADD")
	bvmUtil.Assert(t, o.Length() == 1, "BVM Error! Length of o is wrong.")

	b.Concat(o)
	bvmUtil.Assert(t, b.Length() == 2, "BVM Error! Length is wrong.")
}

func TestBytecodeFinalise(t *testing.T) {
	b := new(Bcode)
	b.Add("ADD")
	b.AddMarker("PUSH", 10)
	b.Finalise()
	bvmUtil.Assert(t, b.Length() == 2, "BVM Error! Length is wrong.")
	bvmUtil.Assert(t, b.commands[1].offset == 10, fmt.Sprintf("wrong offset: %d", b.commands[1].offset))
}
