package bvmCreate

// SetchainMachine ...
func (bvm *BVM) SetchainMachine(s *chainMachine) {
	bench.chainMachine = s
}

// chainMachine implementation
type chainMachine struct {
	items [][]byte
}

// Size ...
func (s *chainMachine) Size() int {
	return len(s.items)
}

// Push pushes bytes onto the stack
func (s *chainMachine) Push(items ...[]byte) {
	s.items = append(s.items, items...)
}

// Pop returns and removes the top byte array from the stack
func (s *chainMachine) Pop() []byte {
	values := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return values
}

func (s *chainMachine) back(num int) []byte {
	return s.items[s.Size()-num-1]
}

// Swap ...
// TODO: check evm implementation???
// TODO: swapping s.Size()-num with s.Size() - 1
func (s *chainMachine) Swap(num int) {
	s.items[s.Size()-1-num], s.items[s.Size()-1] = s.items[s.Size()-1], s.items[s.Size()-1-num]
}

// Dup ...
func (s *chainMachine) Dup(num int) {
	src := s.items[s.Size()-num:]
	dst := make([][]byte, len(src))
	copy(dst, src)
	s.Push(dst...)
}

func (s *chainMachine) validate(size int) bool {
	return len(s.items) >= size
}
