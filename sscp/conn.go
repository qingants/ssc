package sscp

import (
	"bufio"
	"crypto/rc4"
	"encoding/binary"
	"io"
	"net"
	"sync"
	"time"

	"github.com/qingants/ssc/dh64"
	"go.uber.org/zap"
)

var zeroTime time.Time

type cipherConnReader struct {
	sync.Mutex
	reader io.Reader
	cipher *rc4.Cipher
	count  int
}

func (r *cipherConnReader) SetReader(reader io.Reader) {
	r.Lock()
	defer r.Unlock()
	r.reader = reader
}

func (r *cipherConnReader) GetBytesReaded() int {
	return r.count
}

func (r *cipherConnReader) Read(b []byte) (int, error) {
	r.Lock()
	defer r.Unlock()
	n, err := r.reader.Read(b)
	if err != nil {
		return n, err
	}
	r.cipher.XORKeyStream(b[:n], b[:n])
	r.count += n
	return n, err
}

type cipherConnWriter struct {
	sync.Mutex
	writer io.Writer
	cipher *rc4.Cipher
	count  int
}

func (w *cipherConnWriter) SetWriter(writer io.Writer) {
	zap.L().Debug("SetWriter")
	w.Lock()
	defer w.Unlock()
	w.writer = writer
}

func (w *cipherConnWriter) GetBytesWritten() int {
	return w.count
}

func (w *cipherConnWriter) Write(b []byte) (int, error) {
	w.Lock()
	defer w.Unlock()
	sz := len(b)
	buf := defaultBufferPool.Get(sz)
	defer defaultBufferPool.Put(buf)

	space := buf.Bytes()[:sz]
	w.cipher.XORKeyStream(space, b)
	w.count += sz
	_, err := w.writer.Write(b)
	return sz, err
}

func genRC4Key(x, y leu64, key []byte) {
	h := hmac(x, y)
	copy(key, h[:])
}

func newCipherConnReader(secret leu64) *cipherConnReader {
	b := defaultBufferPool.Get(32)
	defer defaultBufferPool.Put(b)

	key := b.Bytes()[:32]
	genRC4Key(secret, toLeu64(0), key[0:8])
	genRC4Key(secret, toLeu64(1), key[8:16])
	genRC4Key(secret, toLeu64(2), key[16:24])
	genRC4Key(secret, toLeu64(3), key[24:32])

	c, _ := rc4.NewCipher(key)
	return &cipherConnReader{
		cipher: c,
	}
}

func newCipherConnWriter(secret leu64) *cipherConnWriter {
	b := defaultBufferPool.Get(32)
	defer defaultBufferPool.Put(b)

	key := b.Bytes()[:32]
	genRC4Key(secret, toLeu64(0), key[0:8])
	genRC4Key(secret, toLeu64(1), key[8:16])
	genRC4Key(secret, toLeu64(2), key[16:24])
	genRC4Key(secret, toLeu64(3), key[24:32])

	c, _ := rc4.NewCipher(key)
	return &cipherConnWriter{
		cipher: c,
	}
}

type Conn struct {
	conn net.Conn
	conf *Config

	connMutex  sync.Mutex
	connErr    error
	handshaked bool
	// HandshakeComplete is true if the handshake has concluded.
	HandshakeComplete bool
	frozen            bool

	in  *cipherConnReader
	out *cipherConnWriter

	id         int
	handshakes int
	secret     [8]byte

	reuseBuffer *loopBuffer

	reused bool // reused conn.
	resend int  // resend data length.

	logger *zap.Logger
}

func (c *Conn) init(id int, secret [8]byte) {
	zap.L().Debug("init conn", zap.Int("id", id), zap.ByteString("secret", secret[:]))

	c.id = id
	c.secret = secret
	c.reuseBuffer = defaultLoopBufferPool.Get()

	c.in = newCipherConnReader(c.secret)
	c.in.SetReader(c.conn)

	c.out = newCipherConnWriter(secret)
	c.out.SetWriter(io.MultiWriter(c.reuseBuffer, c.conn))

	c.reused = false
	c.logger = zap.L().With(zap.Int("id", c.id))

}

func (c *Conn) setConnErr(err error) {
	if c.connErr == nil {
		c.connErr = err
	} else {
		zap.L().Error("errors", zap.Error(c.connErr))
		zap.L().Error("error", zap.Error(err))
	}
}

func (c *Conn) serverNewHandshake(nq *newConnReq) error {
	priKey := dh64.PrivateKey()
	pubKey := dh64.PublicKey(priKey)

	id:=c.conf
}

func (c *Conn)serverHandshake() error {
	c.logger.Debug("start server handshake")
	var sq serverReq
	if err := c.readRecord(&sq); err != nil {
		return err
	}

	switch q := sq.msg.(type) {
	case *newConnReq:
		return c.serverHandshake(q)
	}
}

func (c *Conn) Handshake() error {
	zap.L().Debug("Handshake")
	if HandshakeTimeout > 0 {
		c.SetDeadline(time.Now().Add(HandshakeTimeout))
		defer c.SetDeadline(zeroTime)
	}

	c.connMutex.Lock()
	defer c.connMutex.Unlock()

	if c.handshaked {
		return c.connErr
	}

	var err error
	if

	c.setConnErr(err)
	c.handshaked = true

	return nil
}

func (c *Conn) serverHandleshake() error {
	return nil
}

func (c *Conn) writeRecord(msg handshakeMessage) error {
	zap.L().Debug("writeRecord")
	data := msg.marshal()
	sz := uint16(len(data))

	w := bufio.NewWriter(c.conn)

	err := binary.Write(w, binary.BigEndian, sz)
	if err != nil {
		return err
	}

	if _, err := w.Write(data); err != nil {
		return err
	}

	return w.Flush()
}

func (c *Conn) readRecord(msg handshakeMessage) error {
	var sz uint16
	err := binary.Read(c, binary.BigEndian, &sz)
	if err != nil {
		return err
	}

	buf := defaultBufferPool.Get(int(sz))
	defer defaultBufferPool.Put(buf)

	b := buf.Bytes()[:sz]
	if _, err := io.ReadFull(c.conn, b); err != nil {
		return err
	}

	if err := msg.unmarshal(b); err != nil {
		return err
	}
	return nil
}

// Read reads data from the connection.
// Read can be made to time out and return an error with Timeout == true
// after a fixed time limit; see SetDeadline and SetReadDeadline.
func (c *Conn) Read(b []byte) (int, error) {
	if err := c.Handshake(); err != nil {
		return 0, err
	}

	if len(b) == 0 {
		return 0, nil
	}

	return c.conn.Read(b)
}

// Write writes data to the connection and cache in reuseBuffer
// even failed to write to the connection, the data should still be cached.
//
// Write can be made to time out and return an error after a fixed
// time limit; see SetDeadline and SetWriteDeadline.
func (c *Conn) Write(b []byte) (int, error) {
	if err := c.Handshake(); err != nil {
		return 0, err
	}

	return c.conn.Write(b)
}

// SetDeadline sets the read and write deadlines associated
// with the connection. It is equivalent to calling both
// SetReadDeadline and SetWriteDeadline.
//
// A deadline is an absolute time after which I/O operations
// fail instead of blocking. The deadline applies to all future
// and pending I/O, not just the immediately following call to
// Read or Write. After a deadline has been exceeded, the
// connection can be refreshed by setting a deadline in the future.
//
// If the deadline is exceeded a call to Read or Write or to other
// I/O methods will return an error that wraps os.ErrDeadlineExceeded.
// This can be tested using errors.Is(err, os.ErrDeadlineExceeded).
// The error's Timeout method will return true, but note that there
// are other possible errors for which the Timeout method will
// return true even if the deadline has not been exceeded.
//
// An idle timeout can be implemented by repeatedly extending
// the deadline after successful Read or Write calls.
//
// A zero value for t means I/O operations will not time out.
//
// After a Write has timed out, the TLS state is corrupt and all future
// writes will return the same error.
func (c *Conn) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

// SetReadDeadline sets the deadline for future Read calls on the
// underlying connection.
// and any currently-blocked Read call.
// A zero value for t means Read will not time out.
func (c *Conn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

// SetWriteDeadline sets the deadline for future Write calls
// and any currently-blocked Write call on the underlying connection.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (c *Conn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

func (c *Conn) ID() int {
	return c.id
}

func (c *Conn) TargetServer() string {
	return c.conf.TargetServer
}
