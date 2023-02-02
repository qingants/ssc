package sscp

import "net"

type SSCPServer interface {
	AcquireID() int
	ReleaseID(id int)
	QueryByID(id int) *Conn
}

type Conf struct {
	Flag          int
	TargetServer  string
	ConnForReused *Conn
	ScpServer     SSCPServer
}

var defaultConf = &Conf{}

func (conf *Conf) clone() *Conf {
	return &Conf{
		ScpServer: conf.ScpServer,
	}
}

func Server(conn net.Conn, conf *Conf) *Conn {
	return &Conn{
		conn: conn,
		conf: conf.clone(),
	}
}

func Client(conn net.Conn, conf *Conf) (*Conn, error) {
	if conf == nil {
		conf = defaultConf
	}

	c := &Conn{
		conn: conn,
		conf: conf,
	}

	if conf.ConnForReused != nil {
		if !conf.ConnForReused.spawn(c) {
			return nil, ErrNotAcceptable
		}
		c.handshakes = conf.ConnForReused.handshakes + 1
	}
	return c, nil
}
