// Diffie–Hellman key exchange

package dh64

import "math/rand"

const (
	p uint64 = 0xffffffffffffffc5
	g uint64 = 5
)

func mulModP(a, b uint64) uint64 {
	var m uint64
	for b > 0 {
		if b&1 > 0 {
			t := p - a
			if m >= t {
				m -= t
			} else {
				m += a
			}
		}
		if a >= p-a {
			a = a*2 - p
		} else {
			a = a * 2
		}
		b >>= 1
	}
	return m
}

func powModP(a, b uint64) uint64 {
	if b == 1 {
		return a
	}
	t := powModP(a, b>>1)
	t = mulModP(t, t)
	if b%2 > 0 {
		t = mulModP(t, a)
	}

	return t
}

func powmodp(a, b uint64) uint64 {
	if a == 0 {
		panic("dh64 zero public")
	}
	if b == 0 {
		panic("dh64 zero private key")
	}

	// if a > p {
	// 	a %= p
	// }
	a %= p

	return powModP(a, b)
}

func PrivateKey() uint64 {
	for {
		v := rand.Uint64()
		if v != 0 {
			return v
		}
	}
}

func PublicKey(privateKey uint64) uint64 {
	return powmodp(g, privateKey)
}

func Secret(privateKey, anotherKey uint64) uint64 {
	return powmodp(anotherKey, privateKey)
}
