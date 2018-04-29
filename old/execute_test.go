package bvmCreate

import (
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestExecuteHexFile(t *testing.T) {
	bvm, errs := CreateBVM("example.bvm",
		map[string]int{
			"size": 1,
		},
		deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bvmUtils.Assert(t, errs == nil, "errs should be nil")
	bench.ExecuteHexFile("example.bytes")
}

func TestExecuteFile(t *testing.T) {
	bvm, errs := CreateBVM("example.bvm",
		map[string]int{
			"size": 1,
		},
		deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bvmUtils.Assert(t, errs == nil, "errs should be nil")
	bench.ExecuteFile("example_binary.bytes")
}

func TestExecuteString(t *testing.T) {
	bvm, errs := CreateBVM("example.bvm",
		map[string]int{
			"size": 1,
		},
		deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bvmUtils.Assert(t, errs == nil, "errs should be nil")
	bench.ExecuteString("1010")
}

func TestExecuteHexSingleParameter(t *testing.T) {
	bvm, errs := CreateBVM("example.bvm",
		map[string]int{
			"size": 1,
		}, deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bvmUtils.Assert(t, errs == nil, "errs should be nil")
	err := bench.ExecuteHexString("01")
	bvmUtils.Assert(t, err == nil, "err should be nil")
}

func TestExecuteHexMultipleParameter(t *testing.T) {
	bvm, errs := CreateBVM("example.bvm",
		map[string]int{
			"size":  1,
			"width": 2,
		}, deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bvmUtils.Assert(t, errs == nil, "errs should be nil")
	err := bench.ExecuteHexString("01AAAA")
	bvmUtils.Assert(t, err == nil, "err should be nil")
}

func TestExecuteHexSingleParameterSingleSpec(t *testing.T) {
	bvm, errs := CreateBVM("example.bvm",
		map[string]int{
			"size": 1,
		}, deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bvmUtils.Assert(t, errs == nil, "errs should be nil")
	err := bench.ExecuteHexString("0101")
	bvmUtils.Assert(t, err == nil, "err should be nil")
}

func TestExecuteHexMultipleParameterSingleSpec(t *testing.T) {
	bvm, errs := CreateBVM("example.bvm",
		map[string]int{
			"size":  1,
			"width": 2,
		}, deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bvmUtils.Assert(t, errs == nil, "errs should be nil")
	err := bench.ExecuteHexString("01AAAA01")
	bvmUtils.Assert(t, err == nil, "err should be nil")
}

func TestExecuteHexSingleParameterMultipleSpec(t *testing.T) {
	bvm, errs := CreateBVM("example.bvm",
		map[string]int{
			"size": 1,
		}, deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bvmUtils.Assert(t, errs == nil, "errs should be nil")
	err := bench.ExecuteHexString("010101")
	bvmUtils.Assert(t, err == nil, "err should be nil")
}

func TestExecuteHexMultipleParameterMultipleSpec(t *testing.T) {
	bvm, errs := CreateBVM("example.bvm",
		map[string]int{
			"size":  1,
			"width": 2,
		}, deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bvmUtils.Assert(t, errs == nil, "errs should be nil")
	err := bench.ExecuteHexString("01AAAA0101")
	bvmUtils.Assert(t, err == nil, "err should be nil")
}
