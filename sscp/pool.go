package sscp

import (
	"bytes"
	"sync"
)

type bufferPool struct {
	pool sync.Pool
}

// Get gets a buffer, and guarantees cap for n bytes
func (p *bufferPool) Get(n int) *bytes.Buffer {
	buff := p.pool.Get().(*bytes.Buffer)
	buff.Reset()
	buff.Grow(n)
	return buff
}

func (p *bufferPool) Put(b *bytes.Buffer) {
	// You might be tempted to simplify this by just passing &outBuf to Put,
	// but that would make the local copy of the outBuf slice header escape
	// to the heap, causing an allocation. Instead, we keep around the
	// pointer to the slice header returned by Get, which is already on the
	// heap, and overwrite and return that.
	p.pool.Put(b)
}

func newBufferPool() *bufferPool {
	return &bufferPool{
		pool: sync.Pool{
			New: func() any {
				return new(bytes.Buffer)
			},
		},
	}
}

var defaultBufferPool *bufferPool

func init() {
	defaultBufferPool = newBufferPool()
}
