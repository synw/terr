package main

import (
	"github.com/synw/terr"
)

func f0() *terr.Trace {
	tr := terr.New("f0", "Error zero")
	return tr
}

func f1() *terr.Trace {
	tr := f0()
	tr = tr.Add("f1", "First error", "warn")
	return tr
}

func f2() *terr.Trace {
	tr := f1()
	tr = tr.Add("f2", "Second error", "debug")
	return tr
}

func f3() *terr.Trace {
	tr := f2()
	tr = tr.Add("f3", "Third error", "important")
	return tr
}

func main() {
	tr := f3()
	tr.Check()
}
