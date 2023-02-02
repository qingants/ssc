package sscp

import (
	"fmt"
	"testing"
)

func TestIDAllocator(t *testing.T) {
	t.Logf("Test")
	start := 1
	a := NewIDAllocator(start)
	ID := a.AcquireID()
	if ID != start {
		t.Errorf("error")
	}
	a.Release(ID)

	nextID := start
	for i := 0; i < 100; i++ {
		id := a.AcquireID()
		if id != nextID {
			t.Errorf("error acquire ID")
		}
		nextID++
	}

	for i := start; i < nextID; i++ {
		a.Release(i)
	}

	if a.AcquireID() != start {
		t.Errorf("start ID error")
	}
}

func BenchmarkIDAllocator(b *testing.B) {
	b.Logf("benchmark")
}

func FuzzIDAllactor(f *testing.F) {
	f.Logf("Fuzz")
}

func ExampleIDAllocator() {
	fmt.Println("Example")
}
