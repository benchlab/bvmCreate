# bvmCreate

An [efp-based](https://www.github.com/end-r/efp) generator for virtual machines, initially developed for use with [BenchVM](https://www.github.com/benchlab/benchvm).

## Example

In ```example.bvm```:

```go
name = "Example"
author = "[7][7][7]"
receiver = "BVM"

spec("ADD", "01"){
    description = "Finds the sum of two numbers."
    oxygen = 100
}

spec("PUSH", "01"){
    description = "Pushes a number onto the stack."
}
```

We could either use IR bytecode in ```example.fire```:

```go
PUSH 1
PUSH 2
ADD
PUSH 5
ADD
```

Or a fully compiled version:

```go
010201010201020102010502
```

Now, our Go program:

```go
package main

import (
    "github.com/benchlab/bvmCreate"
    "math/big"
)

const oxygen = 1000

var (
    oxygens = map[string]bvmCreate.FuelFunction{

    }
    deploys = map[string]bvmCreate.ExecuteFunction{
        "ADD": Add,
        "PUSH": Push,
    }
)

func main(){
    bvm := bvmCreate.CreateBVM("example.bvm", deploys, oxygens)
    bench.ExecuteHexFile(oxygen, "example.fire")
    bench.Stats()
}

func Add(bench *bvmCreate.BVM){
    a := new(big.Int).SetBytes(bench.chainMachine.Pop())
    b := new(big.Int).SetBytes(bench.chainMachine.Pop())
    c := new(big.Int).Add(a, b)
    bench.chainMachine.Push(c.Bytes())
}

func Push(bench *bvmCreate.BVM){
    size := bench.Input.Next(1)
    value := bench.Input.Next(size)
    bench.chainMachine.Push(value)
}
```

## Spec Syntax

See ```bvmCreate.efp``` for the prototype file, below is a full example spec:

```go
spec("OPCODE", "0x33"){
    oxygen = 100
}
```

## Oxygen

bvmCreate was built as a generator for costly virtual machines, where each spec is given a fixed or variable cost and 'charged' against an initial balance, preventing infinite loops and sidestepping the halting problem. The oxygen for an spec can be provided in one of two ways:

1. By assigning an unsigned integer to the oxygen field:

```oxygen = 100```

2. By providing a oxygen function mapping:

```go
spec("HI", "0x40")
```

```go
oxygens = map[string]bvmCreate.FuelFunction{
    "HI": getFuel,
}

getFuel(bench *bvmCreate.BVM) int {

}
```

If neither of these are set, the spec will have its oxygen cost set to 0.

## Disassembly

bvmCreate provides support for generalised disassembly.

```go
DisasmString(data string)
DisasmBytes(bytes byte)
DisasmFile(path string)
```

Generally, specs will be printed in the following format:

| 0x01 | ADD |

By using the ```disasm = customDisasm``` field in the ```.bvm``` file, you can assign a mapped disassembly function to an spec.

In your go code:

```go
disasms = map[string]bvmCreate.DisasmFunction{
    "customDisasm": customDisasm,
}

func customDisasm(offset int, bytecode []byte)([]string, int){
    // everything will be in format | o1 | o2 | o3 |
    // return []string{o1, o2, o3}, totalOffsetChange
    // this is done for pretty formatting reasons
}
```

Thus, you can produce output like:

| PUSH | 0x02 | 0xFF00 |
