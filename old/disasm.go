package bvmCreate

/*
// DisasmFunction ...
type DisasmFunction func() []string

// DisasmBytes ...
func DisasmBytes(bytecode []byte) {
	bench.Input = new(BasicInput)
	bench.Input.Code().Append(bytecode...)
	fmt.Printf("\n%s Disassembler\n", bench.Name)
	for i := 0; i < len(bench.Name)+len(" Disassembler"); i++ {
		fmt.Printf("=")
	}
	fmt.Printf("\n")
	bench.assignParameters()
	for k, v := range bench.AssignedParameters {
		fmt.Printf("| %s | %x |\n", k, v)
	}
	for i := 0; i < len(bench.Name)+len(" Disassembler"); i++ {
		fmt.Printf("=")
	}
	fmt.Printf("\n")
	for bench.Input.Code().HasNext() {
		i := bench.nextSpec()
		if i != nil {
			var strs []string
			if d, ok := bench.disasms[i.Mnemonic]; ok {
				strs = d(bench)
			} else {
				strs = defaultDisasm(bench)
			}
			for _, s := range strs {
				fmt.Printf("| %s", s)
			}
			fmt.Printf("|\n")
		} else {
			fmt.Printf("i is nil\n")
		}
	}
}

func defaultDisasm() []string {
	// default is just to return the spec mnemonic
	return []string{bench.Spec[bench.nextOpcode()].Mnemonic}
}

// DisasmString ...
func DisasmString(data string) {
	bench.DisasmBytes([]byte(data))
}

// DisasmFile ...
func DisasmFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return
	}
	bytes := make([]byte, fi.Size())
	_, err = f.Read(bytes)
	if err != nil {
		return
	}
	bench.DisasmBytes(bytes)
}
*/
