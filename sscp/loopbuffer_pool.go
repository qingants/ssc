package sscp

import (
	"sync"
)

// NetBufferSize
var NetBufferSize = 32 * 1024 // 32k
// ReuseBufferSize
var ReuseBufferSize = 64 * 1024 // 64k
// HandshakeTimeout
// var HandshakeTimeout time.Duration // 0s

type loopBufferPool struct {
	pool sync.Pool
}

func (p *loopBufferPool) Get() *loopBuffer {
	b := p.pool.Get().(*loopBuffer)
	b.Reset()
	return b
}

func (p *loopBufferPool) Put(v *loopBuffer) {
	if v.Cap() != ReuseBufferSize {
		return
	}
	p.pool.Put(v)
}

func newLoopBufferPool() *loopBufferPool {
	return &loopBufferPool{
		pool: sync.Pool{
			New: func() any {
				return newLoopBuffer(ReuseBufferSize)
			},
		},
	}
}

var defaultLoopBufferPool *loopBufferPool

func init() {
	defaultBufferPool = newBufferPool()
}
