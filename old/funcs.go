package bvmCreate

var (
	deploys = map[string]ExecuteFunction{
		"ADD":  add,
		"PUSH": push,
		"TEST": test,
	}
)

func add(bench *BVM) {

}

func test(bench *BVM) {
	bench.chainMachine.Push([]byte("a"))
	bench.chainMachine.Push([]byte("b"))
	bench.chainMachine.Push([]byte("c"))
}

func push(bench *BVM) {
	bench.chainMachine.Push([]byte("a"))
}
