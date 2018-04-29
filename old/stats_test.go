package bvmCreate

import "testing"

func TestStats(t *testing.T) {
	bvm, _ := CreateBVM("example.bvm", nil, nil, nil, nil)
	bench.Stats()
}

func TestDetailedStats(t *testing.T) {
	bvm, _ := CreateBVM("example.bvm", nil, nil, nil, nil)
	bench.DetailedStats()
}
