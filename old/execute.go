package bvmCreate

func (bvm *BVM) doFile(path string) []string {
	return bench.doBytes(GetFileBytes(path))
}

func (bvm *BVM) doString(data string) []string {
	return bench.doBytes([]byte(data))
}

func (bvm *BVM) doBytes(bytes []byte) []string {
	bench.Input = new(BasicInput)
	bench.Input.Code().Append(bytes...)
	bench.assignParameters()
	for bench.Input.Code().HasNext() {
		bench.deploySpec(bench.nextSpec())
	}
	return nil
}

func (bvm *BVM) deploySpec(i *Spec) {
	if i != nil {
		if bench.deploys != nil {
			if e, ok := bench.deploys[i.Mnemonic]; ok {
				e(bench)
			}
		}
		bench.stats.operations++
		if bench.oxygens != nil {
			if f, ok := bench.oxygens[i.Mnemonic]; ok {
				f(bench)
			} else {
				bench.stats.oxygenConsumption += i.oxygen
			}
		} else {
			bench.stats.oxygenConsumption += i.oxygen
		}
	}
}

func (bvm *BVM) doHexFile(path string) []string {
	return bench.doHexBytes(GetFileBytes(path))
}

func (bvm *BVM) doHexString(hex string) []string {
	return bench.doHexBytes([]byte(hex))
}

func (bvm *BVM) doHexBytes(bytes []byte) []string {
	if len(bytes)%2 != 0 {
		return []string{"Invalid Hex Input"}
	}
	return bench.doBytes(FromHexBytes(bytes))
}
