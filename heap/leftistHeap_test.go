package heap

import (
	"testing"
)

func Test_leftistHeap(t *testing.T) {
	h := NewLtHeapArray()
	TestHeap(t, h)
}

func Test_leftistHeapUnion(t *testing.T) {
	h,h2 := NewLtHeapArray(),NewLtHeapArray()
	TestHeapUnion(t, h,h2)
}

func Benchmark_leftistHeap(b *testing.B) {
	h := NewLtHeapArray()
	BenchmarkHeap(b,h)
}
