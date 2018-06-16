package main

import (
	"github.com/synw/terr"
)

func f0() *terr.Trace {
	tr := terr.New("Error zero")
	return tr
}

func f1() *terr.Trace {
	tr := f0()
	tr = tr.Add("First error", "debug")
	return tr
}

func f2() *terr.Trace {
	tr := f1()
	tr = tr.Add("Second error", "info")
	return tr
}

func f3() *terr.Trace {
	tr := f2()
	tr = tr.Add("Third error", "warning")
	return tr
}

func f4() *terr.Trace {
	tr := f3()
	tr = tr.Add("Third error", "error")
	return tr
}

func f5() *terr.Trace {
	tr := f4()
	tr = tr.Add("Fourth error", "fatal")
	return tr
}

func main() {
	tr := f5()
	tr.Check()
}
