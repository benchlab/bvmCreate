package bvmCreate

import "testing"

func TestAddMemory(t *testing.T) {
	CreateBVM("example.bvm", nil, nil, nil, nil)
}

func TestGetMemory(t *testing.T) {
	CreateBVM("example.bvm", nil, nil, nil, nil)
}
