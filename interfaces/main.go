package main

import "fmt"

type stucture struct {
	N1, N2 int
}

func (s *stucture) sum() int {
	return s.N1 + s.N2
}

type interface1 interface {
	sum() int
}

func main() {
	//	values := stucture{1, 2}
	//result := values.Sum()
	//fmt.Println(result)

	var a interface1

	sh := stucture{
		N1: 2,
		N2: 1,
	}
	os := otherStruct{
		A: 3,
		B: 4,
	}

	a = &sh
	a = os
	fmt.Println(a.sum())
}

type otherStruct struct {
	A, B int
}

func (a otherStruct) sum() int {
	return a.B + a.A
}
