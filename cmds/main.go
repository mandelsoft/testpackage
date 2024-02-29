package main

import (
	"ocm.test/testpackage/generics"
)

type Elem string

func (e Elem) String() string {
	return string(e)
}

func main() {
	elems := []Elem{"a", "b"}
	generics.Join(elems, ",")
}
