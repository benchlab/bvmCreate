package sample

import (
	"log"

	"github.com/benchlab/bvmCreate"
)

const (
	bvmPath   = "sample.bvm"
	dataPath = "sample.bytes"
)

func newBVM(disasms map[string]bvmCreate.DisasmFunction) *bvmCreate.BVM {
	ops := map[string]bvmCreate.ExecuteFunction{
		"ADD":  add,
		"PUSH": push,
	}
	oxygen := map[string]bvmCreate.FuelFunction{}
	parameters := map[string]int{
		"Size Byte": 1,
	}
	bvm, err := bvmCreate.CreateBVM(bvmPath, parameters, ops, oxygen, disasms)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return bvm
}

func add(bench *bvmCreate.BVM) {

}

func push(bench *bvmCreate.BVM) {
}
