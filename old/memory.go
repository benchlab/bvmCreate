package bvmCreate

// Memory interface
type Memory interface {
	pushViewer()
}

// AddMemory ...
func (bvm *BVM) AddMemory(key string, m Memory) {
	if bench.Memory == nil {
		bench.Memory = make(map[string]Memory)
	}
	bench.Memory[key] = m
}

// GetMemory ...
func (bvm *BVM) GetMemory(key string) Memory {
	return bench.Memory[key]
}
