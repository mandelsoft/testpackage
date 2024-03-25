package main

import (
	"ocm.software/testpackage/generics"
)

type Elem string

func (e Elem) String() string {
	return string(e)
}

func main() {
	elems := []Elem{"a", "b"}
	generics.Join(elems, ",")
}
