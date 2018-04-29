package bvmCreate

import (
	"fmt"
	"log"

	"github.com/end-r/efp"
)

// AddBytecode ...
func (bvm *BVM) AddBytecode(mnemonic string, params ...byte) error {
	if bench.Input == nil {
		bench.Input = new(BasicInput)
	}
	if i, ok := bench.mnemonics[mnemonic]; ok {
		bench.Input.Code().Append(i.Opcode)
		bench.Input.Code().Append(params...)
		bench.NumOpcodes++
		return nil
	}
	return fmt.Errorf("Invalid Spec %s\n", mnemonic)
}

const prototype = `
name : string
author : string
description : string

alias IDOrInt = identifier|int
alias hex = "[0-9][A-F]+"

alias specsAllowed = spec(string, hex){
    description : string
    validate : int
    oxygen : int
}

specsAllowed

category(string){
    description : string
    specsAllowed
}
`

// CreateBVM creates a new BenchVM instance
func CreateBVM(path string, parameters map[string]int,
	deploys map[string]ExecuteFunction, oxygens map[string]FuelFunction,
	disasms map[string]DisasmFunction) (*BVM, []string) {
	p, _ := efp.PrototypeString(prototype)
	e, errs := p.ValidateFile(path)
	if errs != nil {
		log.Println(errs)
		return nil, errs
	}

	var bvm BVM
	// no need to check for nil: would have errored
	bench.Author = e.FirstField("author").Value()
	bench.Name = e.FirstField("name").Value()
	bench.Description = e.FirstField("description").Value()

	bench.Parameters = parameters
	bench.AssignedParameters = make(map[string][]byte)
	bench.stats = new(stats)

	bench.deploys = deploys
	bench.oxygens = oxygens
	bench.disasms = disasms

	bench.Spec = make(map[byte]*Spec)
	bench.mnemonics = make(map[string]*Spec)

	for _, e := range e.Elements("spec") {
		errs := bench.AddSpec(nil, e)
		if errs != nil {
			return nil, errs
		}
	}

	for _, cat := range e.Elements("category") {
		c := new(category)
		c.name = cat.Parameter(0).Value()
		if cat.Fields("description") != nil {
			c.description = cat.FirstField("description").Value()
		}
		for _, e := range cat.Elements("spec") {
			errs := bench.AddSpec(c, e)
			if errs != nil {
				return nil, errs
			}
		}
		if bench.categories == nil {
			bench.categories = make(map[string]*category)
		}
		bench.categories[c.name] = c
	}

	bench.stats = new(stats)
	bench.chainMachine = new(chainMachine)
	return &bvm, nil
}

func stringToOpcode(str string) byte {
	return FromHexString(str)[0]
}

// AddSpec ...
func (bvm *BVM) AddSpec(c *category, e *efp.Element) []string {

	i := new(Spec)
	i.Mnemonic = e.Parameter(0).Value()

	i.Opcode = stringToOpcode(e.Parameter(1).Value())

	if e.Fields("description") != nil && len(e.Fields("description")) > 0 {
		i.description = e.FirstField("description").Value()
	}

	bench.Spec[i.Opcode] = i

	bench.mnemonics[i.Mnemonic] = i

	if c != nil {
		if c.specs == nil {
			c.specs = make(map[string]*Spec)
		}
		c.specs[i.Mnemonic] = i
	}
	return nil
}
