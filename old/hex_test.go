package bvmCreate

import (
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestFromHexString(t *testing.T) {
	hex := "01"
	bytes := FromHexString(hex)
	bvmUtils.Assert(t, len(bytes) == 1, "wrong byte length")
	bvmUtils.Assert(t, bytes[0] == 1, "wrong byte value")
}

func TestFromHexStringError(t *testing.T) {
	hex := "0AA"
	bytes := FromHexString(hex)
	bvmUtils.Assert(t, bytes == nil, "bytes should be nil")
}

func TestFromHexStringTwoBytes(t *testing.T) {
	hex := "0101"
	bytes := FromHexString(hex)
	bvmUtils.Assert(t, len(bytes) == 2, "wrong byte length")
	bvmUtils.Assert(t, bytes[0] == 1, "wrong byte 0 value")
	bvmUtils.Assert(t, bytes[1] == 1, "wrong byte 1 value")
}

func TestFromHexBytes(t *testing.T) {
	bytes := []byte("0101")
	bvmUtils.Assert(t, len(bytes) == 4, "wrong byte length")
	bytes = FromHexBytes(bytes)
	bvmUtils.Assert(t, len(bytes) == 2, "wrong byte length")
}
