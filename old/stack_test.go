package bvmCreate

import (
	"fmt"
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestchainMachinePush(t *testing.T) {
	s := new(chainMachine)
	data := []byte("hello")
	s.Push(data)
	bvmUtils.Assert(t, s.Size() == 1, fmt.Sprintf("wrong length %d, expected %d", s.Size(), len(data)))
}

func TestchainMachinePop(t *testing.T) {
	data := []byte("me")
	bvmUtils.Assert(t, len(data) == 2, "wrong data length")
	s := new(chainMachine)
	bvmUtils.Assert(t, s.Size() == 0, "wrong starting length")
	s.Push(data)
	bvmUtils.Assert(t, s.Size() == 1, "wrong length after push")
	bytes := s.Pop()
	bvmUtils.Assert(t, s.Size() == 0, "wrong length")
	bvmUtils.Assert(t, len(bytes) == len(data),
		fmt.Sprintf("wrong popped value %d, expected %d", len(bytes), len(data)))

}

func TestchainMachineValidate(t *testing.T) {
	data := []byte("me")
	s := new(chainMachine)
	bvmUtils.Assert(t, s.validate(0), "wrong validate 0")
	bvmUtils.Assert(t, !s.validate(1), "wrong validate 1")
	s.Push(data)
	bvmUtils.Assert(t, s.validate(0), "wrong validate 0")
	bvmUtils.Assert(t, s.validate(1), "wrong validate 1")
	bvmUtils.Assert(t, !s.validate(2), "wrong validate 2")
	s.Pop()
	bvmUtils.Assert(t, s.validate(0), "wrong validate 0")
	bvmUtils.Assert(t, !s.validate(1), "wrong validate 1")
}

func TestchainMachineBack(t *testing.T) {
	d1 := []byte("me444")
	d2 := []byte("al44")
	d3 := []byte("xx4")
	s := new(chainMachine)
	s.Push(d1)
	s.Push(d2)
	s.Push(d3)
	bvmUtils.Assert(t, s.Size() == 3, "wrong stack size")
	bvmUtils.Assert(t, len(s.back(2)) == 5, "wrong back value")
}

func TestchainMachineSwap(t *testing.T) {
	d1 := []byte("me444")
	d2 := []byte("al44")
	d3 := []byte("xx4")
	s := new(chainMachine)
	s.Push(d1)
	s.Push(d2)
	s.Push(d3)
	bvmUtils.Assert(t, s.Size() == 3, "wrong stack size")
	bvmUtils.Assert(t, len(s.back(2)) == 5, "wrong bottom value pre-swap")
	bvmUtils.Assert(t, len(s.back(0)) == 3, "wrong top value pre-swap")
	s.Swap(2)
	bvmUtils.Assert(t, s.Size() == 3, "wrong stack size")
	bvmUtils.Assert(t, len(s.back(2)) == 3, "wrong bottom value post-swap")
	bvmUtils.Assert(t, len(s.back(0)) == 5, "wrong top value post-swap")
	s.Swap(1)
	bvmUtils.Assert(t, s.Size() == 3, "wrong stack size")
	bvmUtils.Assert(t, len(s.back(1)) == 5, "wrong bottom value post-swap 2")
	bvmUtils.Assert(t, len(s.back(0)) == 4, "wrong top value post-swap 2")
}

func TestchainMachineDup(t *testing.T) {
	d1 := []byte("me444")
	s := new(chainMachine)
	s.Push(d1)
	bvmUtils.Assert(t, s.Size() == 1, "wrong stack size")
	s.Dup(1)
	bvmUtils.Assert(t, s.Size() == 2, "wrong stack size")
	bvmUtils.Assert(t, len(s.Pop()) == 5, "wrong duplicated data")
	s.Dup(1)
	s.Dup(2)
	bvmUtils.Assert(t, s.Size() == 4, "wrong stack size 1 2")
	bvmUtils.Assert(t, len(s.Pop()) == 5, "wrong duplicated data")
}
