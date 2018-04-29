package bvmCreate

/*
func TestEmptyDisasm(t *testing.T) {
	bvm, _ := CreateBVM("example.bvm", nil, nil, nil, nil)
	bench.DisasmString("")
}

func TestDisasmFile(t *testing.T) {
	bvm, _ := CreateBVM("example.bvm", nil, nil, nil, nil)
	bench.DisasmFile("example.bytes")
}

func TestDisasmHexSingleParameter(t *testing.T) {
	bvm, errs := CreateBVM("example.bvm",
		map[string]int{
			"size": 1,
		}, deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bvmUtils.Assert(t, errs == nil, "errs should be nil")
	bench.DisasmString("01")
}

func TestDisasmHexMultipleParameter(t *testing.T) {
	bvm, errs := CreateBVM("example.bvm",
		map[string]int{
			"size":  1,
			"width": 2,
		}, deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bvmUtils.Assert(t, errs == nil, "errs should be nil")
	bench.DisasmString("01AAAA")
}

func TestDisasmHexSingleParameterSingleSpec(t *testing.T) {
	bvm, errs := CreateBVM("example.bvm",
		map[string]int{
			"size": 1,
		}, deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bvmUtils.Assert(t, errs == nil, "errs should be nil")
	bench.DisasmString("0101")
}

func TestDisasmHexMultipleParameterSingleSpec(t *testing.T) {
	bvm, errs := CreateBVM("example.bvm",
		map[string]int{
			"size":  1,
			"width": 2,
		}, deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bvmUtils.Assert(t, errs == nil, "errs should be nil")
	bench.DisasmString("01AAAA01")
}

func TestDisasmHexSingleParameterMultipleSpec(t *testing.T) {
	bvm, errs := CreateBVM("example.bvm",
		map[string]int{
			"size": 1,
		}, deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bvmUtils.Assert(t, errs == nil, "errs should be nil")
	bench.DisasmString("010101")
}

func TestDiasmHexMultipleParameterMultipleSpec(t *testing.T) {
	bvm, errs := CreateBVM("example.bvm",
		map[string]int{
			"size":  1,
			"width": 2,
		}, deploys, nil, nil)
	bvmUtils.Assert(t, bvm != nil, "bvm shouldn't be nil")
	bvmUtils.Assert(t, errs == nil, "errs should be nil")
	bench.DisasmString("01AAAA0101")
}

*/
