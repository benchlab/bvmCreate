package bvmCreate

import (
	"encoding/binary"
	"fmt"
)

type Bcode struct {
	commands []*Command
}

type Command struct {
	mnemonic   string
	parameters []byte
	isMarker   bool
	offset     int
}

type CostFunc func(interface{}) int

type Spec struct {
	Opcode uint
	Cost   CostFunc
}

type SpecMap map[string]Spec

func (im SpecMap) AddAll(m SpecMap) {
	if im == nil {
		im = make(SpecMap)
	}
	for k, v := range m {
		im[k] = v
	}
}

func (b *Bcode) Length() int {
	return len(b.commands)
}

func (b *Bcode) AddCommand(c *Command) {
	if b.commands == nil {
		b.commands = make([]*Command, 0)
	}
	b.commands = append(b.commands, c)
}

func (b *Bcode) Add(mnemonic string, parameters ...byte) {
	c := Command{
		mnemonic:   mnemonic,
		parameters: parameters,
		isMarker:   false,
	}
	b.AddCommand(&c)
}

func (b *Bcode) AddMarker(mnemonic string, offset int) {
	c := Command{
		mnemonic: mnemonic,
		offset:   offset,
		isMarker: true,
	}
	b.AddCommand(&c)
}

func (b *Bcode) Concat(other Bcode) {
	if other.commands == nil {
		return
	}
	b.commands = append(b.commands, other.commands...)
}

func (b *Bcode) Compare(other Bcode) bool {
	if b.commands == nil && other.commands == nil {
		return true
	}
	if other.commands == nil || b.commands == nil {
		return false
	}
	if len(b.commands) != len(other.commands) {
		return false
	}
	for i, c := range b.commands {
		o := other.commands[i]
		if c.mnemonic != o.mnemonic {
			return false
		}
		for j, p := range c.parameters {
			if p != o.parameters[j] {
				return false
			}
		}
	}
	return true
}

func (b *Bcode) Format() string {
	if b.commands == nil {
		return "BVM Error! There are no commands."
	}
	s := fmt.Sprintf("%d commands\n", len(b.commands))
	for i, c := range b.commands {
		// have to do this manually
		s += fmt.Sprintf("%d | %s %v\n", i+1, c.mnemonic, c.parameters)
	}
	return s
}

func (b *Bcode) Finalise() {
	for i, c := range b.commands {
		if c.isMarker {
			c.parameters = make([]byte, 8)
			binary.LittleEndian.PutUint64(c.parameters, uint64(c.offset+i))
		}
	}
}

func (b *Bcode) compareMnemonics(test []string) bool {
	if len(test) != len(b.commands) {
		return false
	}
	for i, t := range test {
		if t != b.commands[i].mnemonic {
			return false
		}
	}
	return true
}

type Generator interface {
	Spec() SpecMap
}

func (b *Bcode) Generate(g Generator) []byte {
	bytes := make([]byte, 0)
	is := g.Spec()
	for _, c := range b.commands {
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, uint32(is[c.mnemonic].Opcode))
		bytes = append(bytes, bs...)
		bytes = append(bytes, c.parameters...)
	}
	return bytes
}
