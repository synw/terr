package main

import (
	"github.com/synw/terr"
)

func f1() *terr.Trace {
	tr := terr.New("f1", "First error")
	return tr
}

func f2() *terr.Trace {
	tr := f1()
	tr = tr.Stack("f2", "Second error")
	return tr
}

func main() {
	tr := f2()
	tr.Check()
}
