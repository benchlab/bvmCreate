package bvmCreate

import "testing"

func TestGenerateReadme(t *testing.T) {
	bvm, _ := CreateBVM("example.bvm", nil, nil, nil, nil)
	bench.GenerateReadMe("SPEC.md")
}
