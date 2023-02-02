package sscp

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

const (
	SCPFlagForbidForwardIP = 0x1
)

func b64decodeleu64(src string) (v leu64, err error) {
	n := base64.StdEncoding.DecodedLen((len(src)))
	if n < 8 {
		err = ErrIllegalMsg
		return
	}

	dst := make([]byte, n)
	if _, err = base64.StdEncoding.Decode(dst, []byte(src)); err != nil {
		return
	}

	v.Write(dst)
	return
}

func b64encodeLeu64(v leu64) string {
	return base64.StdEncoding.EncodeToString(v[:])
}

type handshakeMessage interface {
	marshal() []byte
	unmarshal([]byte) error
}

type newConnReq struct {
	id           int
	key          leu64
	targetServer string
	flag         int
}

func (r *newConnReq) marshal() []byte {
	s := fmt.Sprintf("%d\n%s\n%s\n%d", r.id, r.key, r.targetServer, r.flag)
	return []byte(s)
}

func (r *newConnReq) unmarshal(s []byte) (err error) {
	lines := strings.Split(string(s), "\n")
	if len(lines) < 2 {
		err = ErrIllegalMsg
		return
	}

	if r.id, err = strconv.Atoi(lines[0]); err != nil {
		return
	}

	if r.key, err = b64decodeleu64(lines[1]); err != nil {
		return
	}

	if len(lines) >= 3 {
		r.targetServer = lines[2]
	}

	if len(lines) >= 4 {
		if r.flag, err = strconv.Atoi(lines[3]); err != nil {
			return
		}
	}

	return
}

type newConnResp struct {
	id  int
	key leu64
}

func (r *newConnResp) marshal() []byte {
	return []byte(fmt.Sprintf("%d\n%s", r.id, r.key))
}

func (r *newConnResp) unmarshal(s []byte) (err error) {
	lines := strings.Split(string(s), "\n")
	if len(lines) < 2 {
		err = ErrIllegalMsg
		return
	}

	if r.id, err = strconv.Atoi(lines[0]); err != nil {
		return
	}

	if r.key, err = b64decodeleu64(lines[1]); err != nil {
		return
	}
	return
}

type serverReq struct {
	msg handshakeMessage
}

func (r *serverReq) marshal() []byte {
	panic("serverReq marshal")
}

func (r *serverReq) unmarshal(s []byte) error {
	if strings.HasPrefix(string(s), "0\n") {
		var nq newConnReq
		if err := nq.unmarshal(s); err != nil {
			return err
		}
		r.msg = &nq
	} /* else {
		var rq reuseReq
		if err := rq.unmarshal(s); err != nil {
			return err
		}
		r.msg = &rq
	}*/
	return nil
}
