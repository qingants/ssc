package sscp

import "io"

type loopBuffer struct {
	buf    []byte
	off    int
	looped bool
}

func (b *loopBuffer) Len() int {
	if b.looped {
		return cap(b.buf)
	}
	return b.off
}

func (b *loopBuffer) Cap() int {
	return cap(b.buf)
}

func (b *loopBuffer) Write(p []byte) (int, error) {
	n := len(p)
	capacity := cap(b.buf)

	// can not save p
	if n >= capacity {
		copy(b.buf, p[n-capacity:])
		b.looped = true
		b.off = 0
		return n, nil
	}

	// can save
	right := capacity - b.off
	if n < right {
		copy(b.buf[b.off:], p)
		b.off += n
		return n, nil
	}

	// fill right
	copy(b.buf[b.off:], p[:right])
	copy(b.buf[0:], p[right:])
	b.looped = true
	b.off = n - right

	return n, nil
}

func (b *loopBuffer) Read(n int) ([]byte, error) {
	if n > b.Len() {
		return nil, io.ErrShortBuffer
	}

	buf := make([]byte, n)
	if n < b.off {
		copy(buf, b.buf[b.off-n:b.off])
		return buf, nil
	}

	wrapped := n - b.off
	copy(buf, b.buf[cap(b.buf)-wrapped:])
	copy(buf[wrapped:], b.buf[:b.off])

	return buf, nil
}

func (b *loopBuffer) Reset() {
	b.off = 0
	b.looped = false
}

func (b *loopBuffer) CopyTo(dst *loopBuffer) {
	if cap(dst.buf) != cap(b.buf) {
		dst.buf = make([]byte, cap(b.buf))
	}
	copy(dst.buf, b.buf)
	dst.off = b.off
	dst.looped = b.looped
}

func newLoopBuffer(cap int) *loopBuffer {
	return &loopBuffer{
		buf:    make([]byte, 0),
		off:    0,
		looped: false,
	}
}
