package bvmCreate

import (
	"fmt"
)

// BVM ...
type BVM struct {
	Name               string
	Author             string
	Description        string
	categories         map[string]*category
	Spec       map[byte]*Spec
	mnemonics          map[string]*Spec
	Parameters         map[string]int
	AssignedParameters map[string][]byte
	PC                 int
	Current            *Spec
	stats              *stats
	chainMachine              *chainMachine
	Memory             map[string]Memory
	Environment        Environment
	State              State
	Input              Input
	NumOpcodes         int

	deploys map[string]ExecuteFunction
	oxygens    map[string]FuelFunction
	disasms  map[string]DisasmFunction
}

// Environment ...
type Environment map[string][]byte

// FuelFunction ...
type FuelFunction func(*BVM) int

// ExecuteFunction ...
type ExecuteFunction func(*BVM)

// Spec for the current BenchVM instance
type Spec struct {
	Mnemonic    string
	Opcode      byte
	description string
	oxygen        int
	count       int
}

type category struct {
	name         string
	description  string
	specs map[string]*Spec
}

func (bvm *BVM) nextSpec() *Spec {
	return bench.Spec[bench.nextOpcode()]
}

func (bvm *BVM) nextOpcode() byte {
	return bench.Input.Code().Next(1)[0]
}

func (bvm *BVM) assignParameters() {
	for k, v := range bench.Parameters {
		bench.AssignedParameters[k] = bench.Input.Code().Next(v)
	}
}

func version() string {
	return fmt.Sprintf("%d.%d.%d", 0, 0, 1)
}
