package main

type Abc interface {
	m1() bool
	m2() int
}

type s1 struct {
}

func (s1) m1() bool {
	println("m1")
	return true
}

func main() {
	s := s1{}
	s.m1()
}
