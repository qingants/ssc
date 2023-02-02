package sscp

import (
	"sync"

	"go.uber.org/zap"
)

type IDAllocator struct {
	sync.Mutex
	start int
	off   int
	free  []int
}

func (a *IDAllocator) AcquireID() int {
	zap.L().Info("IDAllocator.AcquireID")
	a.Lock()
	defer a.Unlock()
	if len(a.free) > 0 {
		index := len(a.free) - 1
		id := a.free[index]
		a.free = a.free[:index]
		return id
	}
	id := a.off
	a.off++
	return id
}

func (a *IDAllocator) Release(id int) {
	zap.L().Info("IDAllocator.Release: ", zap.Int("id", id))
	a.Lock()
	defer a.Unlock()

	index := len(a.free)
	if index == cap(a.free) {
		free := make([]int, index, index*2+1)
		copy(free, a.free)
		a.free = free
	}
	a.free = append(a.free, id)

	if len(a.free) == a.off-a.start {
		a.off = a.start
		a.free = a.free[:0]
	}
}

func NewIDAllocator(start int) *IDAllocator {
	if start < 0 {
		panic("start < 0")
	}
	return &IDAllocator{
		start: start,
		off:   start,
	}
}
