package main

import (
	"github.com/synw/terr"
)

func f1() *terr.Trace {
	tr := terr.New("f1", "Error from f1")
	return tr
}

func f2() *terr.Trace {
	tr := f1()
	tr = tr.Add("f2", "Error from f2")
	return tr
}

func main() {
	tr := f2()
	tr.Check()
}
