package bvmCreate

import (
	"fmt"
	"sort"
)

type stats struct {
	operations      int
	oxygenConsumption int
}

// Stats ...
func (bvm *BVM) Stats() {
	fmt.Printf("BVM %s\n", version())
	fmt.Printf("%s - %s\n", bench.Name, bench.Author)
	fmt.Printf("Operations Executed: %d\n", bench.stats.operations)
	fmt.Printf("Oxygen Used: %d\n", bench.stats.oxygenConsumption)
	fmt.Printf("Oxygen/operation: %f\n", float64(bench.stats.operations)/float64(bench.stats.oxygenConsumption))
}

func (bvm *BVM) sortByFuelConsumption() []*Spec {
	il := make(specList, len(bench.Spec))
	i := 0
	for _, v := range bench.Spec {
		il[i] = v
		i++
	}
	sort.Sort(sort.Reverse(il))
	return il
}

type specList []*Spec

func (l specList) Len() int { return len(l) }
func (l specList) Less(i, j int) bool {
	return (l[i].oxygen * l[i].count) < (l[j].oxygen * l[j].count)
}
func (l specList) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

// DetailedStats ...
func (bvm *BVM) DetailedStats() {
	si := bench.sortByFuelConsumption()
	for i, op := range si {
		fmt.Printf("| %d ", i+1)
		fmt.Printf("| %b ", op.Opcode)
		fmt.Printf("| %d", op.count)
		fmt.Printf("| %d", op.oxygen*op.count)
		fmt.Printf("|\n")
	}
}
