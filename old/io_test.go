package bvmCreate

import (
	"fmt"
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestGetFileBytes(t *testing.T) {
	bytes := GetFileBytes("example.bytes")
	bvmUtils.Assert(t, bytes != nil, "bytes should not be nil")
	bvmUtils.Assert(t, len(bytes) == 4, fmt.Sprintf("wrong bytes length %d, expected %d", len(bytes), 4))
}

func TestGetFileBytesErrors(t *testing.T) {
	bytes := GetFileBytes("not_example.bytes")
	bvmUtils.Assert(t, bytes == nil, "bytes should not be nil")
}

func TestNonExistentBVM(t *testing.T) {
	bvm, errs := CreateBVM("notRealVM.bvm", nil, nil, nil, nil)
	bvmUtils.Assert(t, bvm == nil, "bvm should be nil")
	bvmUtils.Assert(t, errs != nil, "errs shouldn't be nil")
}
