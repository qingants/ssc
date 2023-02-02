package upstream

import (
	"net"
	"sync/atomic"
)

type ResolveRule struct {
	Prefix  string
	Suffix  string
	Port    string
	Pattern string
	// rePattern *regexp.Regexp
}

func (r *ResolveRule) Normalize() error {
	return nil
}

func (r *ResolveRule) Validate(name string) bool {
	return false
}

func (r *ResolveRule) Fullname(name string) string {
	return ""
}

type Option struct {
	Net     string
	Resolve *ResolveRule
}

type Host struct {
	Name        string
	Addr        string
	Weight      int
	MaxFails    int
	FailTimeout int
	Backup      bool
}

type upstreams struct {
	option      atomic.Value
	allHosts    atomic.Value
	byNameHosts atomic.Value
}

func (u *upstreams) SetOption(option *Option) {
	u.option.Store(option)
}

func lookupTCPAddr(hostport string) ([]*net.TCPAddr, error) {
	// host, service, err := net.SplitHostPort(hostport)
	return nil, nil
}
