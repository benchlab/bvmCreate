package bvmCreate

import (
	"os"
	"strconv"
)

// GenerateReadMe generates a read me
func (bvm *BVM) GenerateReadMe(name string) {
	f, _ := os.Create(name)
	defer f.Close()
	f.WriteString("# " + bench.Name + " (" + bench.Author + ")\n")
	f.WriteString(bench.Description + "\n")

	for _, c := range bench.categories {
		f.WriteString("## " + c.name + "\n")
		f.WriteString("### " + c.description + "\n")
		f.WriteString("| " + "Opcode")
		f.WriteString("| " + "Description")
		f.WriteString("| " + "Oxygen")
		f.WriteString("|\n")
		for _, v := range c.specs {
			f.WriteString("| " + string(v.Opcode))
			f.WriteString("| " + v.description)
			f.WriteString("| " + strconv.Itoa(v.oxygen))
			f.WriteString("|\n")
		}
	}
	f.Sync()
}
