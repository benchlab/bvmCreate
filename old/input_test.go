package bvmCreate

import (
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestIterableNext(t *testing.T) {
	i := NewByterable([]byte("abcde"))
	bvmUtils.Assert(t, i.Offset == 0, "wrong starting offset")
	i.Next(1)
	bvmUtils.Assert(t, i.Offset == 1, "wrong offset increment one")
	i.Next(2)
	bvmUtils.Assert(t, i.Offset == 3, "wrong offset increment two")
}

func TestIterableNextValues(t *testing.T) {
	i := NewByterable([]byte("abcde"))
	bvmUtils.Assert(t, i.Offset == 0, "wrong starting offset")
	b1 := i.Next(1)
	bvmUtils.Assert(t, i.Offset == 1, "wrong offset increment one")
	bvmUtils.Assert(t, len(b1) == 1, "wrong b1 length")
	b2 := i.Next(2)
	bvmUtils.Assert(t, i.Offset == 3, "wrong offset increment two")
	bvmUtils.Assert(t, len(b2) == 2, "wrong b2 length")
}

func TestHasNext(t *testing.T) {
	i := NewByterable([]byte("abcde"))
	i.Next(1)
	bvmUtils.Assert(t, i.HasNext(), "should have 4 left")
	i.Next(2)
	bvmUtils.Assert(t, i.HasNext(), "should have 2 left")
	i.Next(2)
	bvmUtils.Assert(t, !i.HasNext(), "should have 0 left")
}
