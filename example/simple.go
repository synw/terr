package main

import (
	"errors"
	"github.com/synw/terr"
)

func f1() *terr.Trace {
	// from an error
	err := errors.New("First error")
	tr := terr.New("f1", err)
	return tr
}

func f2() *terr.Trace {
	tr := f1()
	// from an error message string
	tr = tr.Add("f2", "Second error")
	return tr
}

func main() {
	tr := f2()
	tr.Check()
}
